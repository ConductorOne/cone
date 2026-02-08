package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func functionsNewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new <name>",
		Short: "Scaffold a new C1 Function project",
		Long: `Creates a new directory with a basic C1 Function template.

The template includes:
- main.ts with typed request/response
- deno.json with C1 Functions configuration
- .gitignore`,
		Args: cobra.ExactArgs(1),
		RunE: functionsNewRun,
	}

	cmd.Flags().Bool("force", false, "Overwrite existing directory")

	return cmd
}

func functionsNewRun(cmd *cobra.Command, args []string) error {
	name := args[0]
	force, _ := cmd.Flags().GetBool("force")

	// Check if directory exists
	if _, err := os.Stat(name); err == nil && !force {
		return fmt.Errorf("directory %q already exists (use --force to overwrite)", name)
	}

	// Create directory
	if err := os.MkdirAll(name, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write main.ts
	mainTS := `import { serve } from "https://deno.land/std@0.224.0/http/server.ts";

interface FunctionRequest {
  // Define your input schema here
  message?: string;
}

interface FunctionResponse {
  // Define your output schema here
  result: string;
}

async function handler(req: Request): Promise<Response> {
  if (req.method !== "POST") {
    return new Response("Method not allowed", { status: 405 });
  }

  const input: FunctionRequest = await req.json();

  // Your function logic here
  const output: FunctionResponse = {
    result: input.message ? ` + "`Received: ${input.message}`" + ` : "Hello from C1 Functions!",
  };

  return new Response(JSON.stringify(output), {
    headers: { "Content-Type": "application/json" },
  });
}

const port = parseInt(Deno.env.get("PORT") || "8000");
console.log(` + "`Function listening on port ${port}`" + `);
serve(handler, { port });
`

	if err := os.WriteFile(filepath.Join(name, "main.ts"), []byte(mainTS), 0600); err != nil {
		return fmt.Errorf("failed to write main.ts: %w", err)
	}

	// Write deno.json
	denoJSON := `{
  "name": "` + name + `",
  "version": "0.1.0",
  "tasks": {
    "dev": "deno run --allow-net --allow-read --allow-env --watch main.ts",
    "start": "deno run --allow-net --allow-read --allow-env main.ts"
  },
  "compilerOptions": {
    "strict": true
  },
  "permissions": {
    "net": []
  }
}
`

	if err := os.WriteFile(filepath.Join(name, "deno.json"), []byte(denoJSON), 0600); err != nil {
		return fmt.Errorf("failed to write deno.json: %w", err)
	}

	// Write .gitignore
	gitignore := `.env
.env.local
*.log
`

	if err := os.WriteFile(filepath.Join(name, ".gitignore"), []byte(gitignore), 0600); err != nil {
		return fmt.Errorf("failed to write .gitignore: %w", err)
	}

	pterm.Success.Printf("Created new function project: %s\n", name)
	pterm.Info.Println("Next steps:")
	pterm.Info.Printf("  cd %s\n", name)
	pterm.Info.Println("  cone functions dev")

	return nil
}
