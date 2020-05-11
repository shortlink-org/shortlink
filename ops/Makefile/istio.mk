# ISTIO TASKS ==========================================================================================================
istio-run: ## Run istio
	@istioctl manifest apply --set profile=demo

