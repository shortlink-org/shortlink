# GIT SEETINGS =========================================================================================================
git-config: ## Set git config
	@git config branch.sort -committerdate
	@git maintenance start

	# This command is used to enable the file system monitor feature in Git.
	# The `--global` flag applies the configuration to all repositories on your system.
	# `core.fsmonitor` is the configuration option being set. When its value is `true`, Git will use a file system monitor to automatically update the index state as new files are added or existing files are modified.
	@git config core.fsmonitor true

	# This command is used to automatically set up a remote tracking branch when you push a new local branch to a remote repository.
	# The `--global` flag applies the configuration to all repositories on your system.
	# The `--add` flag adds a new line to the option without altering any existing lines.
	# The `--bool` flag is used to specify that the value being added is a boolean.
	# `push.autoSetupRemote` is the configuration option being set. When its value is `true`, Git will automatically create a remote tracking branch when you push a new local branch to a remote repository.
	@git config --global --add --bool push.autoSetupRemote true
