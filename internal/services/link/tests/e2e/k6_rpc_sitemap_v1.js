import grpc from 'k6/net/grpc';
import { check } from 'k6';

export const options = {
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'gRPC LinkService Sitemap',
    },
  },
}

// Load the gRPC client
const client = new grpc.Client();

// Load the proto file
client.load(['definitions'], './infrastructure/rpc/sitemap/v1/sitemap.proto');

export default () => {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'localhost:50051'

  // Connect to the gRPC server
  client.connect(TARGET_HOSTNAME, { timeout: "5s" });

  // Test the Parse method
  let parseRequest = {
    url: 'https://www.google.com/sitemap.xml',
  };
  let parseResponse = client.invoke('SitemapService/Parse', parseRequest);
  check(parseResponse, {
    'Parse call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Close the connection at the end of the test
  client.close();
};
