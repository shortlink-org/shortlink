# KUBERNETES TASKS =====================================================================================================
PATH_TO_SHORTLINK_CHART := ops/Helm/shortlink-ui
PATH_TO_COMMON_CHART := ops/Helm/common

helm-init: ## helm init
	@add custom repo for helm
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts

helm-lint: ## Check Helm chart
	@helm lint ${PATH_TO_SHORTLINK_CHART}
	@helm lint ${PATH_TO_COMMON_CHART}

helm-deploy: ## Deploy Helm chart to default kube-context and default namespace
	@echo helm install/update ${PROJECT_NAME}
	@helm upgrade ${PROJECT_NAME} ${PATH_TO_SHORTLINK_CHART} \
		--install \
		--force \
		--wait

helm-clean: ## Clean artifact from K8S
	@helm del ${PROJECT_NAME}

helm-common: ## run common service for
	@make helm-init
	@echo helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait
