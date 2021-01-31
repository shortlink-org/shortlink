# NEXT TASKS ===========================================================================================================
PATH_TO_UI_NEXT := ui/next

next_generate: ## Dev-mode Next UI
	@npm --prefix ${PATH_TO_UI_NEXT} start
