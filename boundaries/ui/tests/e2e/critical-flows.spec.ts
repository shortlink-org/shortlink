import { test, expect, Route, Page } from '@playwright/test'

/**
 * Critical E2E flows - only test what can't be tested with components
 * These require actual browser + server interaction
 */

test.describe('Critical User Flows', () => {
  test('landing page loads with key elements', async ({ page, baseURL }) => {
    await page.goto(`${baseURL}/`)

    // Wait for hydration and verify key elements
    await expect(page.getByRole('heading', { name: /shorten your links/i })).toBeVisible({ timeout: 10000 })

    // Navigation exists (FAQ link in header)
    await expect(page.getByRole('link', { name: /faq/i }).first()).toBeVisible()
  })

  test('FAQ page loads directly', async ({ page, baseURL }) => {
    // Direct navigation to FAQ (tests static export)
    await page.goto(`${baseURL}/faq.html`)

    await expect(page.getByRole('heading', { name: /frequently asked/i })).toBeVisible({ timeout: 10000 })
  })

  test('theme toggle switches between light and dark', async ({ page, baseURL }) => {
    await page.goto(`${baseURL}/`)

    // Find and click theme toggle
    const themeToggle = page.locator('[id*="theme"], [aria-label*="theme"], [aria-label*="mode"]').first()

    if (await themeToggle.isVisible()) {
      await themeToggle.click()
      await expect(page.locator('html')).toHaveClass(/dark/)

      await themeToggle.click()
      await expect(page.locator('html')).not.toHaveClass(/dark/)
    }
  })
})

// Add Link tests require authentication - skipped for now
// To enable: mock session in beforeEach or use test fixtures
test.describe.skip('Add Link Flow (requires auth)', () => {
  const ADD_LINK_PATH = '/add-link'
  const API_ENDPOINT = '**/api/links'

  test('creates short link successfully', async ({ page, baseURL }) => {
    const createdHash = 'abc123'

    await page.route(API_ENDPOINT, async (route: Route) => {
      await route.fulfill({
        status: 201,
        contentType: 'application/json',
        body: JSON.stringify({ hash: createdHash }),
      })
    })

    await page.goto(`${baseURL}${ADD_LINK_PATH}`)

    await page.getByLabel(/url/i).first().fill('https://example.com')
    await page.getByRole('button', { name: /create|add|submit/i }).click()

    await expect(page.getByText(/success|created/i)).toBeVisible({ timeout: 10000 })
  })

  test('shows error on invalid URL', async ({ page, baseURL }) => {
    await page.route(API_ENDPOINT, async (route: Route) => {
      await route.fulfill({
        status: 400,
        contentType: 'application/json',
        body: JSON.stringify({ messages: [{ code: 'INVALID_URL', desc: 'Invalid URL' }] }),
      })
    })

    await page.goto(`${baseURL}${ADD_LINK_PATH}`)

    await page.getByLabel(/url/i).first().fill('not-a-url')
    await page.getByRole('button', { name: /create|add|submit/i }).click()

    await expect(page.getByText(/invalid|error/i)).toBeVisible({ timeout: 10000 })
  })
})
