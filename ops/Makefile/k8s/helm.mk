# KUBERNETES TASKS =====================================================================================================
export HELM_EXPERIMENTAL_OCI=1

helm-lint: ## Check Helm chart by linter
	@ops/Makefile/k8s/scripts/helm_lint.sh

# HELM TASKS ===========================================================================================================
helm-docs: ### Generate HELM docs
	@docker run --rm \
		-v ${PWD}/ops/Helm:/helm-docs \
		-v ${PWD}/ops/Makefile/k8s/conf/Helm/README.md.gotmpl:/helm-docs/README.md.gotmpl \
		--workdir="/helm-docs" \
		-u "$(id -u)" \
		jnorwood/helm-docs:v1.12.0 --template-files=/helm-docs/README.md.gotmpl
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
