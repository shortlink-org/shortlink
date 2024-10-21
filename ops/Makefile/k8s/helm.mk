# KUBERNETES TASKS =====================================================================================================
export HELM_EXPERIMENTAL_OCI=1

helm-deps:
	@helm plugin install https://github.com/losisin/helm-values-schema-json.git

helm-lint: ## Check Helm chart by linter
	@ops/Makefile/k8s/scripts/helm_lint.sh

# HELM TASKS ===========================================================================================================
helm-docs: ### Generate HELM docs
	@docker run --rm \
		-v ${PWD}/ops/Helm:/helm-docs \
		-v ${PWD}/ops/Makefile/k8s/conf/Helm/README.md.gotmpl:/helm-docs/README.md.gotmpl \
		--workdir="/helm-docs" \
		-u "$(id -u)" \
		jnorwood/helm-docs:v1.14.2 --template-files=/helm-docs/README.md.gotmpl
	# TODO: remove artifact after generation docs
	@rm ops/Helm/README.md.gotmpl

.PHONY: helm-upgrade
helm-upgrade: ### Upgrade all helm charts
	@helm repo update
	# Find all files named "Chart.yaml" in the current directory and its subdirectories
	@find ./ops/Helm -name "Chart.yaml" | xargs -I '{}' -P 8 bash -c ' \
            dir=$$(dirname "{}"); \
            cd "$$dir"; \
            rm Chart.lock || true; \
            helm dependencies build --skip-refresh \
        '

	@make helm-docs

.PHONY: helm-values-generate
helm-values-generate: ### Generate or process values schema for all Helm charts
	@find "$(CURDIR)/ops/Helm" -type f -name "Chart.yaml" -print0 | \
	while IFS= read -r -d '' file; do \
		dir="$$(dirname "$$file")"; \
		echo "Processing directory: $$dir"; \
		if [ -f "$$dir/.schema.yaml" ]; then \
			echo "Generating values.schema.json in $$dir from .schema.yaml..."; \
			cd "$$dir" && helm schema; \
		else \
			echo "No .schema.yaml found in $$dir, skipping..."; \
		fi; \
	done
