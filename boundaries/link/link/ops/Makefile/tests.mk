# E2E ==================================================================================================================
e2e: ### Run end-to-end tests
	@echo "Running end-to-end tests..."
	@k6 run tests/e2e/k6_rpc_link_v1.js
