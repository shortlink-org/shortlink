# NUXT TASKS ===========================================================================================================
PATH_TO_UI_NUXT := pkg/ui/nuxt

nuxt_generate: ## Deploy nuxt UI
	@npm --prefix ${PATH_TO_UI_NUXT} install
	@npm --prefix ${PATH_TO_UI_NUXT} run generate

# NEXT TASKS ===========================================================================================================
PATH_TO_UI_NEXT := pkg/ui/next

next_run: ## Dev-mode Next UI
	@npm --prefix ${PATH_TO_UI_NUXT} start
