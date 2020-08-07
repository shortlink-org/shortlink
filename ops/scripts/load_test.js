import http from 'k6/http';
import { check, sleep } from 'k6';
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.0.0/index.js";

// init code
const BASE_URL = 'http://localhost:7070';

const params = {
  headers: {
    'Content-Type': 'application/json',
  },
};

export let options = {
  ext: {
    loadimpact: {
      projectID: 123456,
      // Test runs with the same name groups test runs together
      name: "TEST BATCH STORE MODE"
    }
  },
  stages: [
    // { duration: "5m", target: 100 },  // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    { duration: "1m", target: 10000 }, // stay at 100 users for 10 minutes
    // { duration: "5m", target: 0 },    // ramp-down to 0 users
  ],
  thresholds: {
    'http_req_duration': ['p(99)<1500'],      // 99% of requests must complete below 1.5s
    'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
};

export default function() {
  // vu code
  const payload = JSON.stringify({
    url: `https://example.com/${uuidv4()}`,
    describe: "example link",
  });

  let res = http.post(`${BASE_URL}/api`, payload, params);
  check(res, { 'status was 201': r => r.status === 201 });
  sleep(1);
}
