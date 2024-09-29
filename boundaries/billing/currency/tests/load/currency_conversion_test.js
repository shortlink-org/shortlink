import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter } from 'k6/metrics';

// Define custom metrics for success and failure counts
export let successCount = new Counter('success_count');
export let failureCount = new Counter('failure_count');

// Configure the test
export let options = {
    stages: [
        { duration: '30s', target: 10 }, // ramp-up to 10 users over 30 seconds
        { duration: '1m', target: 10 },  // stay at 10 users for 1 minute
        { duration: '30s', target: 0 },  // ramp-down to 0 users over 30 seconds
    ],
};

// Define the base URL of the service
const BASE_URL = 'http://127.0.0.1:3030';

// Define the path for currency conversion
const CURRENCY_CONVERSION_ENDPOINT = `${BASE_URL}/rates/current?base_currency=USD&target_currency=EUR`;

export default function () {
    // Make an HTTP GET request to the currency conversion endpoint
    let response = http.get(CURRENCY_CONVERSION_ENDPOINT);

    // Check the response status
    let checkRes = check(response, {
        'status is 200': (r) => r.status === 200,
        'response body contains exchange_rate': (r) => r.body.includes('exchange_rate'),
    });

    if (checkRes) {
        successCount.add(1);
    } else {
        failureCount.add(1);
    }

    // Add some sleep time to simulate realistic user behavior
    sleep(1);
}
