import { chromium } from 'k6/experimental/browser'
import { check } from 'k6'

// eslint-disable-next-line import/no-mutable-exports
export let options = {
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'Browser main page',
    },
  },
}

export default async function () {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'localhost:3001'

  const browser = chromium.launch({ headless: false })
  const page = browser.newPage()

  try {
    await page.goto(TARGET_HOSTNAME)

    page.locator('#full-width-tab-0').click()
    page.locator('#full-width-tab-1').click()
  } finally {
    page.close()
    browser.close()
  }
}
