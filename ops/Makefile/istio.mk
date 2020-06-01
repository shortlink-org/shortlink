# ISTIO TASKS ==========================================================================================================
istio-up: ## Run istio
	@istioctl manifest apply \
		--set profile=demo
	@kubectl label namespace default istio-injection=enabled

# MetalLB TASKS ========================================================================================================
METALLB_SECRET := "ctVLn6wF2Y3dYIMj/UAo+ffNDv2xiHgEA/+vreUbpxHPlkXakoxCQQ=="

metallb-up: ## Run MetalLB
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/namespace.yaml
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/metallb.yaml
	# On first install only
	@kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="${METALLB_SECRET}"
	# Apply configuration
	@envsubst < /home/batazor/myproejct/shortlink/ops/Helm/addons/metallb/metallb.yaml > /tmp/metallb.yaml
	@kubectl apply -f /tmp/metallb.yaml
