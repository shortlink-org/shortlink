# MINIKUBE TASKS =======================================================================================================

minikube-up: ## run minikube for dev mode
	@minikube start \
		--cpus 4 \
		--memory "16384" \
		--driver=docker \
		--listen-address=0.0.0.0 \
		--extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
		--extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
		--extra-config=apiserver.service-account-issuer=api \
		--extra-config=apiserver.service-account-api-audiences=api,spire-server \
		--extra-config=apiserver.authorization-mode=Node,RBAC \
		--extra-config=kubelet.authentication-token-webhook=true

	# Addons enable
	@eval $(minikube docker-env) # Set docker env
	@minikube addons enable ingress
	@minikube addons enable ingress-dns

minikube-update: ## update image to last version
	@eval $(minikube docker-env) # Set docker env
	@make docker-build           # Build docker images on remote host (minikube)
	@make helm-shortlink-up      # Deploy shortlink HELM-chart

minikube-clear:  ## minikube clear
	@minikube image rm image

minikube-down: ## minikube delete
	@minikube delete
