# CHECK TASKS ==========================================================================================================

# Docs: https://github.com/aquasecurity/kube-bench
kube-bench: ## Run kube-bench
	@docker run \
		--rm -it \
		--pid=host \
		-v /etc:/etc:ro \
		-v /var/lib/kubelet/config.yaml:/var/lib/kubelet/config.yaml:ro \
		-v $(which kubectl):/usr/local/mount-from-host/bin/kubectl \
		-v $$HOME/.kube:/.kube \
		-e KUBECONFIG=/.kube/config \
		aquasec/kube-bench:latest \
		run

# Docs: https://github.com/aquasecurity/kube-hunter
kube-hunter: ## Run kube-hunter
	@docker run \
		-it --rm \
		--network host \
		aquasec/kube-hunter

# Docs: https://github.com/armosec/kubescape
kubescape: ## Run kubescape
	@kubescape scan framework nsa --exclude-namespaces kube-system,kube-public
