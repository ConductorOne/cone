package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"

	"github.com/spf13/cobra"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiry      string `json:"expiry"`
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

	if err := validateArgLenth(0, args, cmd); err != nil {
		return err
	}

	_, tokenHost, err := client.ResolveServerHost(clientId, v)
	if err != nil {
		return err
	}

	tokenSrc, err := client.NewC1TokenSource(ctx,
		clientId, clientSecret,
		tokenHost, v.GetBool("debug"),
	)
	if err != nil {
		return err
	}

	token, err := tokenSrc.Token()
	if err != nil {
		return err
	}

	tokenObj := Token{
		AccessToken: "Bearer " + token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      token.Expiry.String(),
	}

	if v.GetBool(rawTokenFlag) {
		//nolint:forbidigo // We want to raw-print the bearer if this flag is included
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
