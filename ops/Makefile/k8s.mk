# KUBERNETES TASKS =====================================================================================================
PATH_TO_COMMON_CHART := ops/Helm/common

SHORTLINK_NAMESPACE := shortlink
SHORTLINK_HELM_API := ops/Helm/shortlink-api
SHORTLINK_HELM_UI := ops/Helm/shortlink-ui
SHORTLINK_HELM_INGRESS := ops/Helm/ingress

helm-init: ## helm init
	@echo "add custom repo for helm"
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/

helm-lint: ## Check Helm chart
	@helm lint ${PATH_TO_COMMON_CHART}
	@helm lint ${SHORTLINK_HELM_API}
	@helm lint ${SHORTLINK_HELM_UI}
	@helm lint ${SHORTLINK_HELM_INGRESS}

helm-common: ## run common service for
	@make helm-init
	@echo helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait

helm-shortlink-up: ## run shortlink in k8s by Helm
	@echo helm install/update ${PROJECT_NAME}

	@helm upgrade api ${SHORTLINK_HELM_API} \
		--install \
		--force \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
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
	@helm -n ${SHORTLINK_NAMESPACE} del api
	@helm -n ${SHORTLINK_NAMESPACE} del ui
