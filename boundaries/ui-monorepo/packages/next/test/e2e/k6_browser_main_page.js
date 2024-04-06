import { chromium } from 'k6/experimental/browser'

// eslint-disable-next-line no-undef
const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'localhost:3000'

export const options = {
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'Browser main page',
    },
  },
  thresholds: {
    webvital_largest_content_paint: ['p(90) < 1000'],
    [`webvital_first_input_delay{url:${TARGET_HOSTNAME}`]: ['p(90) < 80'],
  },
}

export default async () => {
  const browser = chromium.launch({ headless: false })
  const page = browser.newPage()

  try {
    await page.goto(TARGET_HOSTNAME)
  } finally {
    page.close()
    browser.close()
  }
}
