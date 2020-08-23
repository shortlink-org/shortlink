terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
    postgresql = {
      source = "terraform-providers/postgresql"
    }
  }
  required_version = ">= 0.13"
}
