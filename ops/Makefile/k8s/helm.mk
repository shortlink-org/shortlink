# KUBERNETES TASKS =====================================================================================================
SHORTLINK_HELM_PATH     := ops/Helm
SHORTLINK_HELM_API      := ${SHORTLINK_HELM_PATH}/shortlink-api
SHORTLINK_HELM_LOGGER   := ${SHORTLINK_HELM_PATH}/shortlink-logger
SHORTLINK_HELM_METADATA := ${SHORTLINK_HELM_PATH}/shortlink-metadata
SHORTLINK_HELM_LINK     := ${SHORTLINK_HELM_PATH}/shortlink-link
SHORTLINK_HELM_BOT      := ${SHORTLINK_HELM_PATH}/shortlink-bot
SHORTLINK_HELM_UI       := ${SHORTLINK_HELM_PATH}/shortlink-next
SHORTLINK_HELM_LANDING  := ${SHORTLINK_HELM_PATH}/shortlink-landing
SHORTLINK_HELM_PROXY    := ${SHORTLINK_HELM_PATH}/shortlink-proxy
SHORTLINK_HELM_BOT      := ${SHORTLINK_HELM_PATH}/shortlink-bot
SHORTLINK_HELM_SERVICES := api bot landing link logger metadata next notify proxy
SHORTLINK_HELM_ADDONS := argo cert-manager dashboard flagger flux gateway/istio gateway/istio-ingress gateway/nginx-ingress grafana keda kyverno metallb prometheus-operator mq/rabbitmq rook-ceph store/postgresql store/redis

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/
	@helm repo add stable https://charts.helm.sh/stable
	@helm repo add jetstack https://charts.jetstack.io
	@helm repo add ory https://k8s.ory.sh/helm/charts
	@helm repo add rook-release https://charts.rook.io/release
	@helm repo update

helm-lint: ## Check Helm chart by linter
	helm lint --quiet --with-subcharts ${SHORTLINK_HELM_PATH}/chaos

	for i in $(SHORTLINK_HELM_SERVICES); do \
		helm lint --quiet --with-subcharts ${SHORTLINK_HELM_PATH}/shortlink-$$i; \
  	done; \
  	for i in $(SHORTLINK_HELM_ADDONS); do \
		helm lint --quiet --with-subcharts ${SHORTLINK_HELM_PATH}/addons/$$i; \
  	done

# HELM TASKS ===========================================================================================================
helm-docs: ### Generate HELM docs
	@docker run --rm \
		-v ${PWD}:/helm-docs \
		--workdir="/helm-docs" \
		-u "$(id -u)" \
		jnorwood/helm-docs:v1.11.0
