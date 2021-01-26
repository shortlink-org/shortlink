variable "psql-user" {
  default = "postgres"
  description = "Default PostgreSQL user"
  sensitive = true
}

variable "psql-password" {
  default = "postgres"
  description = "Default PostgreSQL password"
  sensitive = true
}

variable "psql-database" {
  default = "postgres"
  description = "Default PostgreSQL database"
}
