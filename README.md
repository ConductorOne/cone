# Cone

Cone is an open-source utility that provides common user workflows from ConductorOne on the command line. 

## Installation

### From Source

To install Cone from source run `make build` from the root of the repository. This will create a `cone` binary in the `dist/{your_platform}` directory.

Ex: 
```shell
$ make build
go build -o dist/windows_amd64/cone.exe ./cmd/cone
````

### Homebrew

Cone is also deployed through Homebrew. You can install it by running:

```shell
$ brew install conductorone/cone/cone
```

## Getting Started

To get started with Cone, you need to first create an API key in ConductorOne:

1. Log in to ConductorOne.
2. Once logged in, click the user profile in the bottom left corner of the page. From this menu, select `API Keys`.

![API Keys](./docs/images/api-keys.png)

4. Create an API Credential

![API Credential](./docs/images/api-credential.gif)

Then, copy the values for the `Client ID` and `Client Secret` and put them in a config.yaml. This config.yaml can live in the same directory as cone, or in `$HOME/.conductorone/config.yaml`

Example config.yaml:

```yaml
profiles:
    default:
        client-id: lovely-vampire-20397@c1dev.logan.dev.ductone.com/pcc
        client-secret: secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IkVjX3hZWlJ1V3JjQk0wN05lMXY3dEkyRDNoX3dFWmhFWHNvUzluMjh3djQiLCJkIjoiRjZLWGFNSE9idWNqOUJDdjJlRVJyaTJLNkVxbHNmS19Oa2dzdHNxd3FFUSJ9
```

Alternatively, you can set the `CONE_CLIENT_ID` and `CONE_CLIENT_SECRET` environment variables.


Once authenticated, you can start using Cone to perform various tasks and workflows within ConductorOne.

## Usage

Cone provides the following commands:

```shell
cone is... a cone

Usage:
  cone [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  drop        Create a revoke access ticket for an entitlement by slug
  get         Create an access request for an entitlement by slug
  get-user    Get information about a user
  help        Help about any command
  search      Search for entitlements by name or alias
  whoami      Get information about your current user

Flags:
      --client-id string       Client ID
      --client-secret string   Client secret
      --config-path string     path to config file
      --debug                  Enable debug logging
  -h, --help                   help for cone
  -i, --non-interactive        Disable prompts.
  -o, --output string          Output format. Valid values: table, json, json-pretty. (default "table")
  -p, --profile string         The config profile to use. (default "default")
  -v, --version                version for cone

Use "cone [command] --help" for more information about a command.
```