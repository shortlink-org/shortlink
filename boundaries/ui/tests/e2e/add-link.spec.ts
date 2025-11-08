import { test, expect, Route, Page } from '@playwright/test'

const ADD_LINK_PATH = '/add-link'
const API_ENDPOINT = '**/api/links'

test.describe('Add link page', () => {
  test.describe.configure({ timeout: 30_000 })

  test('shows success feedback when a link is created', async ({ page, baseURL }: { page: Page; baseURL?: string }) => {
    const createdHash = 'abc123'

    await page.route(API_ENDPOINT, async (route: Route) => {
      await route.fulfill({
        status: 201,
        contentType: 'application/json',
        body: JSON.stringify({ hash: createdHash }),
      })
    })

    await page.goto(`${baseURL}${ADD_LINK_PATH}`)

    await page.getByLabel('Your URL').fill('https://example.com')
    await page.getByLabel('Describe').fill('Example description')

    await page.getByRole('button', { name: 'Add' }).click()

    await expect(page.getByText('Link created successfully.')).toBeVisible()

    await expect(page.getByRole('link', { name: new RegExp(`/s/${createdHash}$`) })).toBeVisible()
  })

  test('shows mapped error message when API returns failure', async ({ page, baseURL }: { page: Page; baseURL?: string }) => {
    const errorMessage = 'URL is not valid for shortening'

    await page.route(API_ENDPOINT, async (route: Route) => {
      await route.fulfill({
        status: 400,
        contentType: 'application/json',
        body: JSON.stringify({
          messages: [
            {
              code: 'UNKNOWN',
              desc: errorMessage,
            },
          ],
        }),
      })
    })

    await page.goto(`${baseURL}${ADD_LINK_PATH}`)

    await page.getByLabel('Your URL').fill('not-a-valid-url')

    await page.getByRole('button', { name: 'Add' }).click()

    await expect(page.getByText(errorMessage)).toBeVisible()

    await expect(page.getByRole('link', { name: /\/s\// })).toHaveCount(0)
  })

  test('shows fallback error when server response lacks hash', async ({ page, baseURL }: { page: Page; baseURL?: string }) => {
    await page.route(API_ENDPOINT, async (route: Route) => {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({}),
      })
    })

    await page.goto(`${baseURL}${ADD_LINK_PATH}`)

    await page.getByLabel('Your URL').fill('https://missing-hash.com')

    await page.getByRole('button', { name: 'Add' }).click()

    await expect(page.getByText('Could not create the link. Please try again later.')).toBeVisible()

    await expect(page.getByRole('link', { name: /\/s\// })).toHaveCount(0)
  })
})
