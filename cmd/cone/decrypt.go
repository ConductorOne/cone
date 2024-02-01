package main

import (
	"encoding/base64"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/crypto/providers"
	"github.com/conductorone/baton-sdk/pkg/crypto/providers/jwk"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/conductorone/cone/pkg/client"
)

func decryptCredentialCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt-credential <encrypted-credential>",
		Short: "Attempts to decrypt a credential",
		RunE:  decryptCredentialRun,
	}

	return cmd
}

func decryptCredentialRun(cmd *cobra.Command, args []string) error {
	ctx, _, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	// so, we store the ciphertext as []byte, but json serialization of this will be base64 encoded, so let's just require that as our input.
	credentialDec, err := base64.StdEncoding.DecodeString(args[0])
	if err != nil {
		return fmt.Errorf("failed to base64 decode credential: %w", err)
	}

	provider, err := providers.GetEncryptionProvider(jwk.EncryptionProviderJwk)
	if err != nil {
		return fmt.Errorf("failed to get encryption provider: %w", err)
	}

	// FIXME(morgabra): This is a hack for testing.
	ciphertext := &v2.EncryptedData{
		Provider:       jwk.EncryptionProviderJwk,
		EncryptedBytes: credentialDec,
	}

	// Get our secret key and dig out the private jwk, this is silly.
	_, clientSecret, err := getCredentials(v)
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}
	privateJWK, err := client.ParseSecret([]byte(clientSecret))
	if err != nil {
		return fmt.Errorf("failed to parse secret: %w", err)
	}
	privateKeyBytes, err := privateJWK.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal private jwk: %w", err)
	}

	plaintext, err := provider.Decrypt(ctx, ciphertext, privateKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to decrypt credential: %w", err)
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
	return nil
}
