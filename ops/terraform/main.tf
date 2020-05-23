provider "postgresql" {
  alias           = "pg"
  host            = "postgres"
  port            = 5432
  database        = var.psql-database
  username        = var.psql-user
  password        = var.psql-password
  sslmode         = "disable"
  connect_timeout = 15
}

provider "kubernetes" {
  config_context_auth_info = "minikube"
  config_context_cluster   = "minikube"
}

resource "postgresql_database" "my_db" {
  name     = "my_db"
}

# resource "kubernetes_namespace" "example" {
#   metadata {
#     name = "my-first-namespace"
#   }
# }

terraform {
  backend "http" {}
}
