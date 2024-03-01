# GIT SEETINGS =========================================================================================================
git-config: ## Set git config
	@git config --global branch.sort -committerdate
	@git maintenance start
	@git config core.fsmonitor true
