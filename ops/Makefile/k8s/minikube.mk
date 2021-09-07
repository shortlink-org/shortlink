# MINIKUBE TASKS =======================================================================================================

minikube-up: ## run minikube for dev mode
	# for enable audit
	@mkdir -p ~/.minikube/files/etc/ssl/certs
	@cp ${PWD}/ops/Makefile/conf/audit-policy.yaml ~/.minikube/files/etc/ssl/certs/audit-policy.yaml
	@cp ${PWD}/ops/Makefile/conf/tracing-config-file.yaml ~/.minikube/files/etc/ssl/certs/tracing-config-file.yaml

	@minikube start \
		--nodes 3 \
		--cpus 4 \
		--memory "6192" \
		--driver=docker \
		--container-runtime=containerd \
		--listen-address=0.0.0.0 \
		--addons=pod-security-policy,ingress \
		--feature-gates="GracefulNodeShutdown=true" \
		--extra-config=apiserver.tracing-config-file=/etc/ssl/certs/tracing-config-file.yaml \
		--extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
		--extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
		--extra-config=apiserver.service-account-issuer=api \
		--extra-config=apiserver.service-account-api-audiences=api,spire-server \
		--extra-config=apiserver.authorization-mode=Node,RBAC \
		--extra-config=apiserver.enable-admission-plugins=PodSecurityPolicy \
		--extra-config=apiserver.audit-policy-file=/etc/ssl/certs/audit-policy.yaml \
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
