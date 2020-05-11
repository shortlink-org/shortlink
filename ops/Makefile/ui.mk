# UI TASKS =============================================================================================================
PATH_TO_UI_NUXT := pkg/ui/nuxt

nuxt_generate: ## Deploy nuxt UI
	@npm --prefix ${PATH_TO_UI_NUXT} install
	@npm --prefix ${PATH_TO_UI_NUXT} run generate
