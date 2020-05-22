# Include Makefile
include $(SELF_DIR)/ops/Makefile/k8s.shortlink.mk

# KUBERNETES TASKS =====================================================================================================
PATH_TO_COMMON_CHART := ops/Helm/common

SHORTLINK_NAMESPACE := shortlink
SHORTLINK_HELM_API := ops/Helm/shortlink-api
SHORTLINK_HELM_UI := ops/Helm/shortlink-ui
SHORTLINK_HELM_INGRESS := ops/Helm/ingress

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/

helm-lint: ## Check Helm chart
	@helm lint ${PATH_TO_COMMON_CHART}
	@helm lint ${SHORTLINK_HELM_API}
	@helm lint ${SHORTLINK_HELM_UI}
	@helm lint ${SHORTLINK_HELM_INGRESS}

helm-common: ## run common service for
	@make helm-init
	# helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait

