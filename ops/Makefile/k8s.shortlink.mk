# SHORTLINK TASKS ======================================================================================================
helm-shortlink-dep:
	-kubectl create namespace ${SHORTLINK_NAMESPACE}
	-kubectl label namespace ${SHORTLINK_NAMESPACE} istio-injection=enabled

helm-shortlink-up: ## run shortlink in k8s by Helm
	@echo helm install/update ${PROJECT_NAME}

	@helm upgrade api ${SHORTLINK_HELM_API} \
		--install \
		--force \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait

	@helm upgrade ui ${SHORTLINK_HELM_UI} \
		--install \
		--force \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait \
		--set serviceAccount.create=false

	@helm upgrade ingress ${SHORTLINK_HELM_INGRESS} \
		--install \
		--force \
		--namespace=${SHORTLINK_NAMESPACE} \
		--wait

helm-shortlink-down: ## Clean artifact from K8S
	-helm -n ${SHORTLINK_NAMESPACE} del api
	-helm -n ${SHORTLINK_NAMESPACE} del ui
	-helm -n ${SHORTLINK_NAMESPACE} del ingress
