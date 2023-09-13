import grpc from 'k6/experimental/grpc';
import { check } from 'k6';

export const options = {
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
client.load(['../../'], 'infrastructure/rpc/link/v1/link.proto');

export default () => {
  // eslint-disable-next-line no-undef
  const TARGET_HOSTNAME = __ENV.TARGET_HOSTNAME || 'localhost:50051'

  // Connect to the gRPC server
  client.connect(TARGET_HOSTNAME, { timeout: "5s" });

  // Test the Add method
  let addRequest = {
    link: {
      url: 'google.com',
      describe: 'yourDescription',
    },
  };
  let addResponse = client.invoke('LinkService/Add', addRequest);
  check(addResponse, {
    'Add call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Extract the hash from the AddResponse
  let hash = addResponse.message.link.hash;

  // Test the Get method
  let getRequest = { hash: hash };
  let getResponse = client.invoke('LinkService/Get', getRequest);
  check(getResponse, {
    'Get call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the List method
  let listRequest = { filter: hash };
  let listResponse = client.invoke('LinkService/List', listRequest);
  check(listResponse, {
    'List call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the Update method
  let updateRequest = {
    link: {
      url: 'google.com',
      hash: hash,
      describe: 'yourUpdatedDescription',
      // Add timestamps as needed
    },
  };
  let updateResponse = client.invoke('LinkService/Update', updateRequest);
  check(updateResponse, {
    'Update call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the Delete method
  let deleteRequest = { hash: hash };
  let deleteResponse = client.invoke('LinkService/Delete', deleteRequest);
  check(deleteResponse, {
    'Delete call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Negative test: Get method should fail after Delete
  getResponse = client.invoke('LinkService/Get', getRequest);
  check(getResponse, {
    'Get call failed after delete': (r) => r && r.status !== grpc.StatusOK,
  });

  // Close the connection at the end of the test
  client.close();
};
