# ANSIBLE TASKS ========================================================================================================
export ANSIBLE_CONFIG=ops/ansible/ansible.cfg

ansible-dep: ## Install ansible dep
	@ansible-galaxy install -r ops/ansible/requirements.yml

ansible-up: ## Apply ansible playbook
	@ansible-playbook ops/ansible/playbooks/playbook.yml -i ops/ansible/hosts

ansible-do: ## Test
	@ansible-playbook ops/ansible/playbooks/playbook.yml -v -i ops/ansible/hosts \
 		--tags "test" \
 		--extra-vars "rest_server=http://google.com username=admin password=admin"

