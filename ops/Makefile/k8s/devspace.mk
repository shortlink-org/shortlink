# DEVSPACE TASKS =======================================================================================================
devspace-init: ## Run local minikube and set default params
	@make minikube-up
	@devspace use namespace shortlink

devspace-up: ## Up dev-workspace
	@devspace dev

devspace-down: ## Down dev-workspace
	@devspace purge --dependencies

