import grpc from 'k6/net/grpc';
import { check } from 'k6';

export const options = {
  insecureSkipTLSVerify: true,
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'gRPC LinkService CRUD',
    },
  },
}

// Load the gRPC client
const client = new grpc.Client();

// Load the proto file
client.load(['../..'], 'internal/infrastructure/rpc/link/v1/link_rpc.proto');

export default () => {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || '127.0.0.1:443'

  // Connect to the gRPC server
  client.connect(TARGET_HOSTNAME, { timeout: "5s" });

  // Set the metadata for requests
  let params = {
    metadata: {
      'user-id': '1',
    },
    tags: { k6test: 'yes' },
  }

  // Mock input data
  const url = 'http://google.com/' + Math.random();

  // Test the Add method
  let addRequest = {
    link: {
      url: url,
      describe: 'yourDescription',
    },
  };
  let addResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/Add', addRequest, params);
  check(addResponse, {
    'Add call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Extract the hash from the addResponse
  let hash = addResponse.message.link.hash;

  // Test the Get method
  let getRequest = { hash: hash };
  let getResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/Get', getRequest, params);
  check(getResponse, {
    'Get call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the List method
  let listRequest = {
    filter: JSON.stringify({
      hash: { eq: hash },
    }),
  };
  let listResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/List', listRequest, params);
  check(listResponse, {
    'List call succeeded': (r) => r && r.status === grpc.StatusOK,
    'List call returned one link ore more': (r) => r && r.message.links.link.length === 1,
  });

  // Test the Update method
  let updateRequest = {
    link: {
      url: url,
      hash: hash,
      describe: 'yourUpdatedDescription',
    },
  };
  let updateResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/Update', updateRequest, params);
  check(updateResponse, {
    // TODO: implement the Update method in the service
    // 'Update call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the Delete method
  let deleteRequest = { hash: hash };
  let deleteResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/Delete', deleteRequest, params);
  check(deleteResponse, {
    'Delete call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Negative test: Get method should fail after Delete
  getResponse = client.invoke('infrastructure.rpc.link.v1.LinkService/Get', getRequest, params);
  check(getResponse, {
    'Get call failed after delete': (r) => r && r.status !== grpc.StatusOK,
  });

  // Close the connection at the end of the test
  client.close();
};
