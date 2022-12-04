import { sleep, check } from 'k6';
import tracing, { Http } from 'k6/x/tracing'
import { Counter } from 'k6/metrics';

// A simple counter for http requests
export const requests = new Counter('http_reqs');

// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for
export const options = {
  stages: [
    { target: 20, duration: '1m' },
    { target: 15, duration: '1m' },
    { target: 0, duration: '1m' },
  ],
  thresholds: {
    requests: ['count < 100'],
  },
};

export function setup() {
  console.log(`Running xk6-distributed-tracing v${tracing.version}`, tracing);
}

export default function() {
  const http = new Http({
    propagator: "w3c",
  });

  // our HTTP request, note that we are saving the response to res, which can be accessed later
  const res = http.get('http://test.k6.io');
  console.log(`trace_id=${res.trace_id}`);

  check(res, {
    'status is 200': r => r.status === 200,
    'response body': r => r.body.indexOf('Feel free to browse') !== -1,
  });
  sleep(1);
}
