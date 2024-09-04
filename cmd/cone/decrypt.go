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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/square/go-jose.v2"

	"github.com/conductorone/cone/pkg/client"
)

func decryptCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt-credential <app-id>",
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

	fmt.Printf("Thumbprint: %s\n", thumbprint(privateJWK.Public()))
	return plaintext, nil
}

func decryptCredentialRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	err = validateArgLenth(0, args, cmd)
	if err != nil || validateArgLenth(1, args, cmd) != nil {
		return err
	}

	var apps []shared.App

	if len(args) > 0 {
		apps = []shared.App{{ID: &args[0]}}

	} else {
		apps, err = c.ListApps(ctx)
		if err != nil {
			return err
		}
	}

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
			creds, err := c.ListAppUserCredentials(ctx, *app.ID, *appUser.ID)
			if err != nil {
				return err
			}
			allCreds = append(allCreds, creds...)
		}
	}

	fmt.Printf("Found %d credentials\n", len(allCreds))
	for _, cred := range allCreds {
		plaintext, err := decodeCredential(ctx, v, cred)
		if err != nil {
			return fmt.Errorf("failed to decode credential: %w", err)
		}
		opt := &protojson.MarshalOptions{
			Multiline:       true,
			Indent:          "  ",
			EmitUnpopulated: true,
		}
		pt, err := opt.Marshal(plaintext)
		if err != nil {
			return fmt.Errorf("failed to marshal plaintext: %w", err)
		}
		fmt.Printf("Decrypted credential: %s\n", pt)
		fmt.Printf("Decrypted bytes: %s\n", plaintext.Bytes)
	}

	return nil
}
