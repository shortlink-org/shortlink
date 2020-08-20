# MINIKUBE TASKS =======================================================================================================
MINIKUBE_IP := $(shell minikube ip)

minikube-up: ## run minikube for dev mode
	@minikube start \
		--cpus 4 \
		--memory "16384" \
		--driver=docker \
		--extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
		--extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
		--extra-config=apiserver.service-account-issuer=api \
		--extra-config=apiserver.service-account-api-audiences=api,spire-server \
		--extra-config=apiserver.authorization-mode=Node,RBAC \
		--extra-config=kubelet.authentication-token-webhook=true
	@minikube addons enable ingress
	@eval $(minikube docker-env) # Set docker env

minikube-update: ## update image to last version
	@eval $(minikube docker-env) # Set docker env
	@make docker-build           # Build docker images on remote host (minikube)
	@make helm-shortlink-up      # Deploy shortlink HELM-chart

minikube-down: ## minikube delete
	@minikube delete
