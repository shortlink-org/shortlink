import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:7070/api';
const COOKIES = {
  'ory_kratos_session': 'MTcwNDMwNzMyM3w1Uld1eHNQbXgxRFRjN0I1TUtRNklwSHZ0bkREbXVKekR3S0laQlhFQW10WW51U0pEenBrMlhFbGxWbDBxd0diOWl1Z0dpa1JkZWdRX0RQc3o1c0NCMGpkb0w1Tm1GaGRGVE1rVklSR3Y5T0c5YW9Bc1BwRGwzOE1kUmNyU3BmU3g1cHdkNmY3SjJNQkRyNk5xQzJ5Mi1UWkJiN2FNRWZXVGtXSU5Ba0tKQk1zcXI1YmxVQzFYc3dCMm5ucjREV3BqTWpjUnUxaDBVTzQwQXVnZ2stV18zeUdsaW5UZVp2VVBMeXpNdElUcnJZRVVHaXRoalFuVkxsZC1RZ0tleDVqMUlLR1hYejdoekw5N1d3WldiV3RDdz09fKBq4mA2i_s9wuF0uaRzUh8Bc0-BuVUfUOGc__-xSBzT'
};

function createLink(url, description) {
  const payload = JSON.stringify({ url, describe: description });
  const headers = {
    'Content-Type': 'application/json',
  };
  return http.post(`${BASE_URL}/links`, payload, { headers });
}

function getLink(hash) {
  return http.get(`${BASE_URL}/links/${hash}`);
}

function updateLink(hash, newUrl, newDescription) {
  const payload = JSON.stringify({
    filter: { id: hash },
    link: { url: newUrl, hash: hash, describe: newDescription }
  });
  const headers = { 'Content-Type': 'application/json' };
  return http.put(`${BASE_URL}/links`, payload, { headers });
}

function deleteLink(hash) {
  return http.del(`${BASE_URL}/links/${hash}`);
}

export default function () {
  // Set up the cookie jar for the VU
  const jar = http.cookieJar();
  jar.set(`${BASE_URL}/`, 'ory_kratos_session', COOKIES['ory_kratos_session'], {
    domain: 'localhost',  // Replace with the actual domain
    path: '/',            // Replace with the actual path if more specific
    secure: false,        // Set to true if using https
    httpOnly: true,       // Set to true if the cookie is httpOnly
    // Other cookie attributes as needed
  });

  // Create a new link
  let createResponse = createLink("https://example.com", "Example description");
  check(createResponse, { 'Link created successfully': (r) => r.status === 201 });
  console.warn("TEST", createResponse.json());
  let linkHash = createResponse.json().hash;
  sleep(1);

  // Retrieve the created link
  let getResponse = getLink(linkHash);
  check(getResponse, { 'Link retrieved successfully': (r) => r.status === 200 });
  sleep(1);

  // Update the link
  // TOOD: enable after fixing the update endpoint
  // let updateResponse = updateLink(linkHash, "https://example-updated.com", "Updated description");
  // check(updateResponse, { 'Link updated successfully': (r) => r.status === 200 });
  // sleep(1);

  // Delete the link
  let deleteResponse = deleteLink(linkHash);
  check(deleteResponse, { 'Link deleted successfully': (r) => r.status === 204 });
  sleep(1);

  // Verify the link is deleted
  let verifyDeleteResponse = getLink(linkHash);
  console.warn("TEST", verifyDeleteResponse.json());
  check(verifyDeleteResponse, { 'Link not found after deletion': (r) => r.status === 404 });
}
