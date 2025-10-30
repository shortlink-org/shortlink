import { loadEnvConfig } from '@next/env'
import { defineConfig } from 'cypress'

const { combinedEnv } = loadEnvConfig(process.cwd())

export default defineConfig({
  env: combinedEnv,
  e2e: {
    baseUrl: 'http://localhost:3000/next',
    retries: {
      runMode: 3,
    },
    viewportHeight: 1080,
    viewportWidth: 1920,
    video: false,
    screenshotOnRunFailure: false,
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
})
