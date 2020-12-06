# Include Makefile
include $(SELF_DIR)/ops/Makefile/k8s.shortlink.mk
include $(SELF_DIR)/ops/Makefile/k8s.velero.mk

# KUBERNETES TASKS =====================================================================================================
PATH_TO_COMMON_CHART := ops/Helm/common

SHORTLINK_HELM_API      := ops/Helm/shortlink-api
SHORTLINK_HELM_LOGGER   := ops/Helm/shortlink-logger
SHORTLINK_HELM_METADATA := ops/Helm/shortlink-metadata
SHORTLINK_HELM_UI       := ops/Helm/shortlink-ui
SHORTLINK_HELM_INGRESS  := ops/Helm/ingress

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
	@helm lint ${SHORTLINK_HELM_INGRESS}

helm-common-up: ## run common service
	@make helm-init
	# helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait

helm-common-down: ## down common service
	-helm del common

# CT TASKS =============================================================================================================
export KIND=v0.9.0

ct-lint: ### Check Helm chart by ct lint
	@docker run -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v ${PWD}:/home \
		quay.io/helmpack/chart-testing bash -c "cd /home && ct lint --all --config ct.yaml"

ct-run: ### Check Helm chart by ct install
	@docker run -it --rm --network host \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v ${PWD}:/home \
		quay.io/helmpack/chart-testing bash -c "\
				cd /home && pwd && ls -la && \
				apk add -U docker && \
				wget -O /usr/local/bin/kind https://github.com/kubernetes-sigs/kind/releases/download/${KIND}/kind-linux-amd64 && \
				chmod +x /usr/local/bin/kind && \
				kind create cluster --wait 2m --config=./ops/Helm/kind-config.yaml && \
				ct install --all --config ct.yaml"
