package proxy

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"io"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/conductorone/dpop/pkg/dpop"
	"github.com/elazarl/goproxy"
	"golang.org/x/oauth2"
)

//go:embed libs/*
var libsFS embed.FS

const (
	// C1FunctionImportHost is the virtual host for @c1/ imports.
	C1FunctionImportHost = "c1-function-import"
)

// Config holds proxy configuration.
type Config struct {
	// Allowlist of hosts that functions can access (in addition to C1 API).
	// Empty means allow all.
	Allowlist []string

	// C1APIHost is the ConductorOne API host (e.g., "tenant.conductor.one").
	C1APIHost string

	// TokenSource provides OAuth2 tokens for C1 API requests.
	TokenSource oauth2.TokenSource

	// Proofer generates DPoP proofs for C1 API requests (optional).
	Proofer *dpop.Proofer

	// Secrets to inject as environment variables.
	Secrets map[string]string

	// SourceDir is the local source directory for serving local imports.
	SourceDir string
}

// Proxy is a MITM proxy that intercepts HTTPS traffic,
// injects auth headers for C1 API, and enforces an egress allowlist.
type Proxy struct {
	proxy    *goproxy.ProxyHttpServer
	ca       *CA
	config   Config
	listener net.Listener
}

// New creates a new proxy with the given CA and configuration.
func New(ca *CA, config Config) (*Proxy, error) {
	proxy := goproxy.NewProxyHttpServer()

	// Create TLS certificate for MITM
	tlsCert, err := tls.X509KeyPair(ca.CertPEM, ca.KeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to create TLS cert: %w", err)
	}
	goproxy.GoproxyCa = tlsCert

	// Configure MITM for all HTTPS
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// Handle @c1/ imports.
	//nolint:bodyclose // goproxy handlers manage response body lifecycle
	proxy.OnRequest(goproxy.DstHostIs(C1FunctionImportHost)).DoFunc(importsHandler(config.SourceDir))

	// Handle C1 API requests - inject auth.
	if config.C1APIHost != "" {
		//nolint:bodyclose // goproxy handlers manage response body lifecycle
		proxy.OnRequest(isDstC1API(config.C1APIHost)).DoFunc(apiHandler(config.TokenSource, config.Proofer, config.C1APIHost))
	}

	// Handle all other requests - check allowlist.
	//nolint:bodyclose // goproxy handlers manage response body lifecycle
	proxy.OnRequest().DoFunc(allowlistHandler(config.Allowlist, config.C1APIHost))

	proxy.Verbose = false

	return &Proxy{
		proxy:  proxy,
		ca:     ca,
		config: config,
	}, nil
}

// Start starts the proxy on a random available port and returns the address.
func (p *Proxy) Start(ctx context.Context) (string, error) {
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", "127.0.0.1:0")
	if err != nil {
		return "", fmt.Errorf("failed to start proxy listener: %w", err)
	}
	p.listener = listener

	server := &http.Server{
		Handler:           p.proxy,
		ReadHeaderTimeout: 30 * time.Second,
	}
	go func() {
		<-ctx.Done()
		_ = server.Close()
	}()
	go func() {
		_ = server.Serve(listener)
	}()

	return listener.Addr().String(), nil
}

// Stop stops the proxy.
func (p *Proxy) Stop() error {
	if p.listener != nil {
		return p.listener.Close()
	}
	return nil
}

// ProxyURL returns the proxy URL for use in HTTPS_PROXY environment variable.
func (p *Proxy) ProxyURL() string {
	if p.listener == nil {
		return ""
	}
	return "http://" + p.listener.Addr().String()
}

// Secrets returns the secrets map for environment injection.
func (p *Proxy) Secrets() map[string]string {
	return p.config.Secrets
}

// isDstC1API returns a condition that matches C1 API hosts.
func isDstC1API(c1Host string) goproxy.ReqConditionFunc {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
		host := req.URL.Host
		if host == "" {
			host = req.Host
		}
		// Strip port
		if h, _, err := net.SplitHostPort(host); err == nil {
			host = h
		}
		return host == c1Host || strings.HasSuffix(host, ".conductor.one")
	}
}

// apiHandler injects OAuth2 + DPoP authentication for C1 API requests.
func apiHandler(tokenSource oauth2.TokenSource, proofer *dpop.Proofer, c1Host string) func(*http.Request, *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if tokenSource == nil {
			return req, nil
		}

		// Get token
		token, err := tokenSource.Token()
		if err != nil {
			return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusUnauthorized, "Failed to get token: "+err.Error())
		}

		// Inject Authorization header
		req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)

		// Rewrite host to actual C1 API
		req.URL.Host = c1Host
		req.URL.Scheme = "https"
		req.Host = c1Host

		// Generate DPoP proof if proofer is available
		if proofer != nil {
			proof, err := proofer.CreateProof(context.Background(), req.Method, req.URL.String(),
				dpop.WithProofNowFunc(time.Now),
				dpop.WithAccessToken(token.AccessToken),
			)
			if err == nil {
				req.Header.Set("DPoP", proof)
			}
		}

		return req, nil
	}
}

// allowlistHandler checks if the request host is in the allowlist.
func allowlistHandler(allowlist []string, c1Host string) func(*http.Request, *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		host := req.URL.Host
		if host == "" {
			host = req.Host
		}
		// Strip port
		if h, _, err := net.SplitHostPort(host); err == nil {
			host = h
		}

		// C1 API and import host always allowed
		if host == c1Host || strings.HasSuffix(host, ".conductor.one") || host == C1FunctionImportHost {
			return req, nil
		}

		// If no allowlist, allow all
		if len(allowlist) == 0 {
			return req, nil
		}

		// Check allowlist
		for _, allowed := range allowlist {
			if host == allowed {
				return req, nil
			}
			// Wildcard support
			if strings.HasPrefix(allowed, "*.") {
				suffix := allowed[1:] // ".example.com"
				if strings.HasSuffix(host, suffix) || host == allowed[2:] {
					return req, nil
				}
			}
		}

		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusForbidden,
			fmt.Sprintf("Host %q not in egress allowlist", host))
	}
}

// importsHandler serves @c1/ imports from embedded libs or local source.
func importsHandler(_ string) func(*http.Request, *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		cleanPath := path.Clean(req.URL.Path)
		if !path.IsAbs(cleanPath) {
			return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusNotFound, "Invalid path")
		}

		ext := path.Ext(cleanPath)
		if ext != ".ts" && ext != ".js" && ext != ".json" {
			return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusNotFound, "Invalid extension")
		}

		// Try embedded libs first
		libPath := path.Join("libs", cleanPath)
		if f, err := libsFS.Open(libPath); err == nil {
			defer f.Close()
			data, _ := io.ReadAll(f)
			return nil, goproxy.NewResponse(req, getContentType(ext), http.StatusOK, string(data))
		}

		return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusNotFound, "Import not found")
	}
}

func getContentType(ext string) string {
	switch ext {
	case ".ts":
		return "application/typescript; charset=utf-8"
	case ".js":
		return "application/javascript; charset=utf-8"
	case ".json":
		return "application/json; charset=utf-8"
	default:
		return "text/plain; charset=utf-8"
	}
}

// ParseAllowlist parses an allowlist from a deno.json-style config.
func ParseAllowlist(allowlist []string) []string {
	var result []string
	for _, entry := range allowlist {
		entry = strings.TrimPrefix(entry, "https://")
		entry = strings.TrimPrefix(entry, "http://")
		// Strip path
		if idx := strings.Index(entry, "/"); idx > 0 {
			entry = entry[:idx]
		}
		// Strip port
		if h, _, err := net.SplitHostPort(entry); err == nil {
			entry = h
		}
		if entry != "" {
			result = append(result, entry)
		}
	}
	return result
}
