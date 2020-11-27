# SHORTLINK TASKS ======================================================================================================
SHORTLINK_NAMESPACE := shortlink

helm-shortlink-dep:
	-kubectl create namespace ${SHORTLINK_NAMESPACE}
	-kubectl label namespace ${SHORTLINK_NAMESPACE} istio-injection=enabled

helm-shortlink-up: ## run shortlink in k8s by Helm
	@echo helm install/update ${PROJECT_NAME}

	-make helm-shortlink-dep

	@helm upgrade api ${SHORTLINK_HELM_API} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait

	@helm upgrade ui ${SHORTLINK_HELM_UI} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait \
		--set serviceAccount.create=false

	# Add IP in /etc/hosts
	@echo "$(minikube ip) ui-nuxt.local" | sudo tee -a /etc/hosts

	@helm upgrade ingress ${SHORTLINK_HELM_INGRESS} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait

helm-shortlink-down: ## Clean artifact from K8S
	-helm -n ${SHORTLINK_NAMESPACE} del api
	-helm -n ${SHORTLINK_NAMESPACE} del ui
	-helm -n ${SHORTLINK_NAMESPACE} del ingress
