package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/conductorone/baton-sdk/pkg/crypto"
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
	_, _, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}
	credential := args[0]

	fmt.Println(credential)

	crypto.EncryptionManager{}
	return nil
}
