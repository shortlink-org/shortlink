# MINIKUBE TASKS =======================================================================================================
VARIABLE_NAME ?= containerd

minikube-up: ## run minikube for dev mode
	# for enable audit
	@mkdir -p ~/.minikube/files/etc/ssl/certs
	@cp ${PWD}/ops/Makefile/conf/tracing-config-file.yaml ~/.minikube/files/etc/ssl/certs/tracing-config-file.yaml

	@minikube start \
		--nodes 1 \
		--cpus 4  \
		--memory "4192" \
		--driver=docker \
		--container-runtime=${VARIABLE_NAME} \
		--addons=pod-security-policy,ingress \
		--feature-gates="GracefulNodeShutdown=true,EphemeralContainers=true" \
		--extra-config=apiserver.tracing-config-file=/etc/ssl/certs/tracing-config-file.yaml \
		--extra-config=apiserver.authorization-mode=Node,RBAC \
		--extra-config=apiserver.audit-log-path=- \
		--extra-config=kubelet.authentication-token-webhook=true

	# Addons enable
	@eval $(minikube docker-env) # Set docker env

	# Change context (optional)
	-kubectx minikube

minikube-update: ## update image to last version
	@eval $(minikube docker-env) # Set docker env
	@make docker-build           # Build docker images on remote host (minikube)
	@make helm-shortlink-up      # Deploy shortlink HELM-chart

minikube-clear:  ## minikube clear
	@minikube image rm image

minikube-down: ## minikube delete
	@minikube delete

minikube-stop: ## minikube stop
	@minikube stop
