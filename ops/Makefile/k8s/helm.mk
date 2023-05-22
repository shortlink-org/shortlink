# KUBERNETES TASKS =====================================================================================================
export HELM_EXPERIMENTAL_OCI=1

SHORTLINK_HELM_PATH     := ops/Helm
SHORTLINK_HELM_SERVICES := api bot common landing link logger metadata next notify proxy workflows
SHORTLINK_HELM_ADDONS   := argocd cert-manager flagger gateway/istio gateway/nginx-ingress grafana keda knative-operator kyverno metallb prometheus-operator mq/rabbitmq rook-ceph store/postgresql store/redis

helm-init: ## helm init
	# add custom repo for helm
	@helm repo add istio https://storage.googleapis.com/istio-release/releases/1.5.4/charts/
	@helm repo add stable https://charts.helm.sh/stable
	@helm repo add jetstack https://charts.jetstack.io
	@helm repo add ory https://k8s.ory.sh/helm/charts
	@helm repo add rook-release https://charts.rook.io/release
	@helm repo add bitnami oci://registry-1.docker.io/bitnamicharts
	@helm repo add kiali https://kiali.org/helm-charts
	@helm repo update

helm-lint: ## Check Helm chart by linter
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

.PHONY: helm-update-charts
helm-update-charts: ### Update Helm charts
	@helm repo update
	# Find all files named "Chart.yaml" in the current directory and its subdirectories
	@find . -name "Chart.yaml" | xargs -I '{}' -P 8 bash -c ' \
            dir=$$(dirname "{}"); \
            cd "$$dir"; \
            rm Chart.lock || true; \
            helm dependencies build --skip-refresh \
        '

	@make helm-docs
