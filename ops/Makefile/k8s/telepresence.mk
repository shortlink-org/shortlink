# TELEPRESENCE TASKS ===================================================================================================
telepresence-up: ## Starts the local daemon and connects Telepresence to your cluster and installs the Traffic Manager if it is missing.
	@telepresence connect

telepresence-down: ## Quits the local daemon, stopping all intercepts and outbound traffic to the cluster
	@telepresence quit
	@telepresence uninstall --everything
