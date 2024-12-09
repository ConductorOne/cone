package main

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/crypto/providers"
	"github.com/conductorone/baton-sdk/pkg/crypto/providers/jwk"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/square/go-jose.v2"

	"github.com/conductorone/cone/pkg/client"
)

func decryptCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt-credential [app-id]",
		Short: "Attempts to decrypt a credential",
		RunE:  decryptCredentialRun,
	}

	return cmd
}

func thumbprint(jwk jose.JSONWebKey) string {
	tp, err := jwk.Thumbprint(crypto.SHA256)
	if err != nil {
		panic(fmt.Errorf("failed to compute key id: %w", err))
	}
	return hex.EncodeToString(tp)
}

func decodeCredential(ctx context.Context, v *viper.Viper, cred shared.AppUserCredential) (*v2.PlaintextData, error) {
	// so, we store the ciphertext as []byte, but json serialization of this will be base64 encoded, so let's just require that as our input.
	credentialDec, err := base64.StdEncoding.DecodeString(*cred.EncryptedData.EncryptedBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode credential: %w", err)
	}

	provider, err := providers.GetEncryptionProvider(jwk.EncryptionProviderJwk)
	if err != nil {
		return nil, fmt.Errorf("failed to get encryption provider: %w", err)
	}

	// FIXME(morgabra): This is a hack for testing.
	ciphertext := &v2.EncryptedData{
		Provider:       jwk.EncryptionProviderJwk,
		EncryptedBytes: credentialDec,
	}

	// Get our secret key and dig out the private jwk, this is silly.
	_, clientSecret, err := getCredentials(v)
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}
	privateJWK, err := client.ParseSecret([]byte(clientSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to parse secret: %w", err)
	}
	privateKeyBytes, err := privateJWK.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private jwk: %w", err)
	}

	plaintext, err := provider.Decrypt(ctx, ciphertext, privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt credential: %w", err)
	}

	pterm.Printf("Thumbprint: %s\n", thumbprint(privateJWK.Public()))
	return plaintext, nil
}

func decryptCredentialRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if validateArgLenth(0, args, cmd) != nil && validateArgLenth(1, args, cmd) != nil {
		return fmt.Errorf("expected 0 or 1 arguments, got %d\n%s", len(args), cmd.UsageString())
	}

	appMap := make(map[string]shared.App)
	var apps []shared.App
	if len(args) > 0 {
		app, err := c.GetApp(ctx, args[0])
		if err != nil {
			return err
		}
		apps = append(apps, *app)
		appMap[*app.ID] = shared.App{ID: &args[0]}
	} else {
		apps, err = c.ListApps(ctx)
		if err != nil {
			return err
		}
	}

	// Get the c1 user ID
	resp, err := c.AuthIntrospect(ctx)
	if err != nil {
		return err
	}
	userID := client.StringFromPtr(resp.UserID)

	allCreds := make([]shared.AppUserCredential, 0)

	for _, app := range apps {
		if app.ID == nil {
			continue
		}
		appUsers, err := c.ListAppUsers(ctx, *app.ID)
		if err != nil {
			return err
		}
		for _, appUser := range appUsers {
			if *appUser.GetIdentityUserID() != userID {
				continue
			}
			creds, err := c.ListAppUserCredentials(ctx, *app.ID, *appUser.ID)
			if err != nil {
				return err
			}
			allCreds = append(allCreds, creds...)
			appMap[*app.ID] = app
		}
	}

	pterm.Printf("Found %d credentials\n", len(allCreds))
	pterm.Printf("========================================\n")
	for i, cred := range allCreds {
		pterm.Printf("========================================\n")
		pterm.Printf("Credential #%d\n", i+1)
		pterm.Printf("App Display Name: %s\n", *appMap[*cred.AppID].DisplayName)
		pterm.Printf("App ID: %s\n", *cred.AppID)
		pterm.Printf("App User ID: %s\n", *cred.AppUserID)
		plaintext, err := decodeCredential(ctx, v, cred)
		if err != nil {
			pterm.Printf("Failed to decode credential: %s\n", err.Error())
			pterm.Printf("========================================\n")
			continue
		}
		pterm.Printf("Decrypted Credential: %s\n", plaintext.Bytes)
		pterm.Printf("========================================\n")
	}

	return nil
}
