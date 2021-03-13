# DEVSPACE TASKS =======================================================================================================
devspace-init: ## Run local minikube and set default params
	@make minikube-up
	@devspace use namespace shortlink

devspace-up: ## Up dev-workspace
	@devspace dev --config ops/devspace.yaml

devspace-down: ## Down dev-workspace
	@devspace purge --dependencies --config ops/devspace.yaml

