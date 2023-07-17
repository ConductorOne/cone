package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"

	"github.com/spf13/cobra"
)

type Token struct {
	AccessToken string `header:"access_token"`
	TokenType   string `header:"token_type"`
	Expiry      string `header:"expiry"`
}

func tokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "",
		RunE:  tokenRun,
	}
	addRawTokenFlag(cmd)
	return cmd
}

func tokenRun(cmd *cobra.Command, args []string) error {
	ctx, _, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return err
	}

	if len(args) != 0 {
		usageErrorString := cmd.UsageString()
		return fmt.Errorf("expected 0 arguments, got %d\n"+usageErrorString, len(args))
	}

	tokenSrc, _, _, err := client.NewC1TokenSource(ctx, clientId, clientSecret)
	if err != nil {
		return err
	}

	token, err := tokenSrc.Token()
	if err != nil {
		return err
	}

	tokenObj := Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      token.Expiry.String(),
	}

	if v.GetBool(rawTokenFlag) {
		fmt.Println(tokenObj.AccessToken)
		return nil
	}

	outputManager := output.NewManager(ctx, v)
	err = outputManager.Output(ctx, &tokenObj, output.WithTransposeTable())
	if err != nil {
		return err
	}

	return nil
}

func (r *Token) Header() []string {
	return []string{
		"Access Token",
		"Token Type",
		"Expiry",
	}
}

func (r *Token) rows() []string {
	return []string{
		r.AccessToken,
		r.TokenType,
		r.Expiry,
	}
}

func (r *Token) Rows() [][]string {
	return [][]string{r.rows()}
}
