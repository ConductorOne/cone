terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "c1dev.anthony.dev.ductone.com:2443"
  client_id     = "envious-fairy-21060@c1dev.anthony.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IjZSV3MtMGc0bGIxZU5VMGxMeThKOHd3TUY1RnU3SUMxR0hhNmVQVXhEYnMiLCJkIjoidnpMd0hJRXZ0eWhEZjBFdGNWMDQ1T0c2dFhINXNKSElQYlJVbHVNM2pYVSJ9"
}
