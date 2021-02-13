# KUBERNETES TASKS =====================================================================================================
PATH_TO_COMMON_CHART := ops/Helm/common

SHORTLINK_HELM_API      := ops/Helm/shortlink-api
SHORTLINK_HELM_LOGGER   := ops/Helm/shortlink-logger
SHORTLINK_HELM_METADATA := ops/Helm/shortlink-metadata
SHORTLINK_HELM_UI       := ops/Helm/shortlink-ui
SHORTLINK_HELM_LANDING  := ops/Helm/shortlink-landing

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/
	@helm repo add stable https://charts.helm.sh/stable
	@helm repo add jetstack https://charts.jetstack.io
	@helm repo update

helm-lint: ## Check Helm chart by linter
	@helm lint ${PATH_TO_COMMON_CHART}
	@helm lint ${SHORTLINK_HELM_API}
	@helm lint ${SHORTLINK_HELM_LOGGER}
	@helm lint ${SHORTLINK_HELM_METADATA}
	@helm lint ${SHORTLINK_HELM_UI}
	@helm lint ${SHORTLINK_HELM_LANDING}

helm-common-up: ## run common service
	@make helm-init
	# helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait

helm-common-down: ## down common service
	-helm del common

# HELM TASKS ===========================================================================================================
helm-docs: ### Generate HELM docs
	@docker run --rm --volume "$(pwd):/helm-docs" -u "$(id -u)" jnorwood/helm-docs:v1.5.0
