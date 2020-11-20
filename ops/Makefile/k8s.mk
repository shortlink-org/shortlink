# Include Makefile
include $(SELF_DIR)/ops/Makefile/k8s.shortlink.mk
include $(SELF_DIR)/ops/Makefile/k8s.velero.mk

# KUBERNETES TASKS =====================================================================================================
PATH_TO_COMMON_CHART := ops/Helm/common

SHORTLINK_NAMESPACE     := shortlink
SHORTLINK_HELM_API      := ops/Helm/shortlink-api
SHORTLINK_HELM_LOGGER   := ops/Helm/shortlink-logger
SHORTLINK_HELM_METADATA := ops/Helm/shortlink-metadata
SHORTLINK_HELM_UI       := ops/Helm/shortlink-ui
SHORTLINK_HELM_INGRESS  := ops/Helm/ingress

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/
	@helm repo add stable https://kubernetes-charts.storage.googleapis.com
	@helm repo update

helm-lint: ## Check Helm chart by linter
	@helm lint ${PATH_TO_COMMON_CHART}
	@helm lint ${SHORTLINK_HELM_API}
	@helm lint ${SHORTLINK_HELM_LOGGER}
	@helm lint ${SHORTLINK_HELM_METADATA}
	@helm lint ${SHORTLINK_HELM_UI}
	@helm lint ${SHORTLINK_HELM_INGRESS}

helm-test: ### Check Helm chart by run
	@docker run -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(pwd):/home \
		quay.io/helmpack/chart-testing bash -c "cd /home && ct lint --all --config ct.yaml"

helm-common: ## run common service for
	@make helm-init
	# helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait
