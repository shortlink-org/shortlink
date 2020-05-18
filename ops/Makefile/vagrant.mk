# VAGRANT TASKS ========================================================================================================
export VAGRANT_VAGRANTFILE=./ops/vagrant/Vagrantfile

vagrant-dep: ## Install vagrant plugins
	@vagrant plugin install vagrant-env

vagrant-up: ## Run vagrant VM
	@vagrant up --provider=virtualbox

vagrant-reload: ## Reload vagrant VM, loads new Vagrantfile configuration
	@vagrant reload

vagrant-down: ## Down vagrant VM
	@vagrant destroy --force
	@rm -rf .vagrant
	@rm -f *-console.log
