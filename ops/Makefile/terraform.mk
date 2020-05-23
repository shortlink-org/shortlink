# TERRAFORM TASKS ======================================================================================================
TF_DATA_DIR := ops/terraform

terraform-init: ## Terraform init
	@terraform init ${TF_DATA_DIR}
	@terraform plan ${TF_DATA_DIR}
	@terraform validate ${TF_DATA_DIR}

terraform-up: ## Terraform up
	@terraform apply -auto-approve ${TF_DATA_DIR}
