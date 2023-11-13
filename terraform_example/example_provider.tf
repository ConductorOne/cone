terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "SERVER_URL"
  client_id     = "CONE_CLIENT_ID"
  client_secret = "CONE_CLIENT_SECRET"
}
