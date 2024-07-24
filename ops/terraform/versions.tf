terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
    postgresql = {
      source = "terraform-providers/postgresql"
    }
  }
  required_version = ">= 1.9.3"
}
