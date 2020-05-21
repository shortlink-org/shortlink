# ISTIO TASKS ==========================================================================================================
istio-up: ## Run istio
	@istioctl manifest apply --set profile=demo
	@kubectl label namespace default istio-injection=enabled
