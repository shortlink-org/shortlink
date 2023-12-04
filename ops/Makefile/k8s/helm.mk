# KUBERNETES TASKS =====================================================================================================
export HELM_EXPERIMENTAL_OCI=1

SHORTLINK_HELM_PATH     := ops/Helm
SHORTLINK_HELM_SERVICES := api bot common landing link logger metadata next notify proxy workflows
SHORTLINK_HELM_ADDONS   := argocd cert-manager gateway/istio gateway/nginx-ingress grafana keda knative-operator kyverno prometheus-operator mq/rabbitmq rook-ceph store/postgresql store/redis

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
		-v ${PWD}/ops/Helm:/helm-docs \
		-v ${PWD}/ops/Makefile/k8s/conf/Helm/README.md.gotmpl:/helm-docs/README.md.gotmpl \
		--workdir="/helm-docs" \
		-u "$(id -u)" \
		jnorwood/helm-docs:v1.11.0 --template-files=/helm-docs/README.md.gotmpl
	# TODO: remove artifact after generation docs
	@rm ops/Helm/README.md.gotmpl

.PHONY: helm-upgrade
helm-upgrade: ### Upgrade all helm charts
	@helm repo update
	# Find all files named "Chart.yaml" in the current directory and its subdirectories
	@find . -name "Chart.yaml" | xargs -I '{}' -P 8 bash -c ' \
            dir=$$(dirname "{}"); \
            cd "$$dir"; \
            rm Chart.lock || true; \
            helm dependencies build --skip-refresh \
        '

	@make helm-docs
