# KUBERNETES TASKS =====================================================================================================
SHORTLINK_HELM_PATH     := ops/Helm
SHORTLINK_HELM_SERVICES := api bot common landing link logger metadata next notify proxy
SHORTLINK_HELM_ADDONS := argocd cert-manager flagger gateway/istio gateway/nginx-ingress grafana keda knative-operator kyverno metallb prometheus-operator mq/rabbitmq rook-ceph store/postgresql store/redis

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/
	@helm repo add stable https://charts.helm.sh/stable
	@helm repo add jetstack https://charts.jetstack.io
	@helm repo add ory https://k8s.ory.sh/helm/charts
	@helm repo add rook-release https://charts.rook.io/release
	@helm repo add bitnami https://charts.bitnami.com/bitnami
	@helm repo add kiali https://kiali.org/helm-charts
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

helm-update: ### Update Helm charts
	@find . -name "Chart.yaml" -exec bash -c 'cd "$(dirname "$1")" && rm Chart.lock || true && helm dependencies build --skip-refresh' _ {} \;
	@make helm-docs
