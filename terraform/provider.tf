terraform {
  required_providers {
    conductorone = {
      source  = "registry.terraform.io/ConductorOne/conductorone"
      version = "1.0.0"
    }
  }
}

provider "conductorone" {
  server_url    = "https://c1dev.anthony.dev.ductone.com:2443"
  client_id     = "outstanding-werewolf-26337@c1dev.anthony.dev.ductone.com/pcc"
  client_secret = "secret-token:conductorone.com:v1:eyJrdHkiOiJPS1AiLCJjcnYiOiJFZDI1NTE5IiwieCI6IlFWWjVIejJuTjFlcm1PNXdtaGNsQThpRTVRQjlZOTQ4a1hyU3BuWjR5NkkiLCJkIjoiUDhucUxCZlo0UnZPTHM3MDdhU1JYbXBDY3c0QU1ZejU5bE9Wb0tQQTFLTSJ9"
}
