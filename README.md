# Cone: The ConductorOne command line tool

Welcome to `cone` – a robust command-line tool from ConductorOne! Written in Go, `cone` is designed to streamline the use-case-specific tasks for developers, security teams, and other end users. `cone` is open source, under the [Aapche 2.0 license](./LICENSE).

# Installation

Cone releases are available from the [official distribution center](https://dist.conductorone.com/ConductorOne/cone). Each version provides signed macOS, Linux, and Windows artifacts with checksums, provenance, and SBOM attestations.

Install with [Homebrew](https://brew.sh/):

```shell
$ brew install conductorone/cone/cone
```

Run Cone in a container with a pinned Public ECR version:

```shell
$ docker pull public.ecr.aws/conductorone/cone:<version>
```

# Authentication

To authenticate to Cone, run `cone login <tenant-name or tenant-url>`, passing in the name (such as `example.conductor.one`) or URL (such as `https://example.conductor.one`) of your ConductorOne instance and follow the prompts. 

# Getting started with Cone

Run `cone help` to see the full list of available Cone commands. 

Check out the ConductorOne documentation for much more on Cone, including:

- [Full installation and authentication guide](https://www.conductorone.com/docs/product/cli/install)
- [Introduction to Cone and tutorial video](https://www.conductorone.com/docs/product/cli/intro)
- [Command reference](https://www.conductorone.com/docs/product/cli/commands)

# Contributing, support, and issues

We value your contributions and ideas, no matter how small. We aim to make `cone` a fantastic tool for everyone. If you encounter any issues, require support, or have suggestions, please open a Github Issue!

Check out our [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more information.

# License

`cone` is licensed under the [Apache 2.0 license](./LICENSE).

# Contact

For further information or assistance, feel free to contact us at [support@conductorone.com](mailto:support@conductorone.com).

## Reporting security issues

At ConductorOne, we prioritize security and take potential issues seriously. If you discover a security issue, please alert us as quickly as possible!

To protect our users, **DO NOT** create a public issue or pull request. Instead, send your report privately to the ConductorOne Security Team at [security@conductorone.com](mailto:security@conductorone.com).

We greatly appreciate security reports, and we will publicly acknowledge your contribution once the issue has been resolved and poses no risk to users. Thank you for helping us ensure the security and integrity of `cone`!
