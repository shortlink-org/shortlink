import { defineConfig, devices } from '@playwright/test'

export default defineConfig({
  testDir: './tests/e2e',
  timeout: 30_000,
  expect: {
    timeout: 5_000,
  },
  // Optimization: single worker = single browser instance
  fullyParallel: false,
  workers: 1,

  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,

  reporter: [
    ['list'],
    ['html', { open: 'never' }],
    ...(process.env.CI
      ? [['allure-playwright', { detail: true, outputFolder: 'playwright-report/allure-results', suiteTitle: false }] as const]
      : []),
  ],

  use: {
    baseURL: process.env.PLAYWRIGHT_BASE_URL ?? 'http://localhost:3000',
    trace: 'retain-on-failure',
    screenshot: 'only-on-failure',
    video: 'off',
    contextOptions: {
      reducedMotion: 'reduce', // Faster animations
    },
  },

  projects: [
    {
      name: 'chromium',
      use: {
        ...devices['Desktop Chrome'],
        launchOptions: {
          args: ['--disable-gpu', '--no-sandbox'],
        },
      },
    },
  ],

  webServer: process.env.PLAYWRIGHT_WEB_SERVER
    ? undefined
    : {
        command: 'npx serve out -l 3000',
        port: 3000,
        reuseExistingServer: true,
        timeout: 60_000,
      },
})
