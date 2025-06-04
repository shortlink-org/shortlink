import { check } from 'k6'
import http from 'k6/http'
import tracing from 'k6/experimental/tracing'

export const options = {
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'HTTP main page',
    },
  },
}

tracing.instrumentHTTP({
  // possible values: "w3c", "jaeger"
  propagator: 'w3c',
})

export default () => {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'http://localhost:3000'

  const res = http.get(`${TARGET_HOSTNAME}/`, {
    headers: {
      trace_id: 'instrumented/get',
    },
  })
  check(res, {
    'is status 200': (r) => r.status === 200,
  })
}
