package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/conductorone/cone/pkg/scaffold"
	"github.com/spf13/cobra"
)

// connectorInitCmd returns the command for initializing new connector projects.
func connectorInitCmd() *cobra.Command {
	var modulePath string
	var description string

	cmd := &cobra.Command{
		Use:   "init <name>",
		Short: "Create a new connector project",
		Long: `Create a new ConductorOne connector project from the standard template.

The project will be created in a directory named "baton-<name>" in the current
working directory.

Examples:
  cone connector init my-app
  cone connector init my-app --module github.com/myorg/baton-my-app
  cone connector init my-app --description "Connector for My App"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			// Normalize name (remove baton- prefix if present)
			name = strings.TrimPrefix(name, "baton-")

			// Determine output directory
			outputDir := fmt.Sprintf("baton-%s", name)

			// Check if directory exists
			if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
				return fmt.Errorf("directory already exists: %s", outputDir)
			}

			cfg := &scaffold.Config{
				Name:        name,
				ModulePath:  modulePath,
				OutputDir:   outputDir,
				Description: description,
			}

			fmt.Printf("Creating connector project: %s\n", outputDir)

			if err := scaffold.Generate(cfg); err != nil {
				return fmt.Errorf("failed to generate project: %w", err)
			}

			// Verify Go installation
			fmt.Println("\nProject created successfully!")
			fmt.Println("\nNext steps:")
			fmt.Printf("  cd %s\n", outputDir)
			fmt.Println("  go mod tidy")
			fmt.Println("  # Edit pkg/client/client.go to implement API calls")
			fmt.Println("  # Edit pkg/connector/*.go to implement resource syncers")
			fmt.Println("  go build")
			fmt.Println("  cone connector dev")

			return nil
		},
	}

	cmd.Flags().StringVarP(&modulePath, "module", "m", "", "Go module path (default: github.com/conductorone/baton-<name>)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Connector description")

	return cmd
}
