# VAGRANT TASKS ========================================================================================================
export VAGRANT_VAGRANTFILE=./ops/vagrant/Vagrantfile

vagrant-dep: ## Install vagrant plugins
	@vagrant plugin install vagrant-env

vagrant-up: ## Run vagrant VM
	@vagrant up --provider=virtualbox

