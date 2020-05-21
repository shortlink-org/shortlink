# MINIKUBE TASKS =======================================================================================================
minikube-up: ## run minikube for dev mode
	@minikube start \
		--cpus 4 \
		--memory "16384" \
		--driver=docker
	@minikube addons enable ingress
	@eval $(minikube docker-env) # Set docker env

minikube-update: ## update image to last version
	@eval $(minikube docker-env) # Set docker env
	@make docker-build           # Build docker images on remote host (minikube)
	@make helm-deploy            # Deploy shortlink HELM-chart

minikube-down: ## minikube delete
	@minikube delete
