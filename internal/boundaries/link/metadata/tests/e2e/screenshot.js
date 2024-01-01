import grpc from 'k6/net/grpc';
import { check } from 'k6';

export const options = {
  insecureSkipTLSVerify: true,
  ext: {
    loadimpact: {
      // eslint-disable-next-line no-undef
      projectID: __ENV.K6_PROJECT_ID,
      // Test runs with the same name groups test runs together
      name: 'gRPC MetaDataService CRUD',
    },
  },
}

// Load the gRPC client
const client = new grpc.Client();

// Load the proto file
client.load(['../../'], 'infrastructure/rpc/metadata/v1/metadata_rpc.proto');

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

  // Test the Set method
  let addRequest = {
    url: "https://google.com"
  }
  let setResponse = client.invoke('infrastructure.rpc.metadata.v1.MetadataService/Set', addRequest, params);
  check(setResponse, {
    'Add call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  // Test the Get method
  let getRequest = {
    url: "https://google.com"
  }
  let getResponse = client.invoke('infrastructure.rpc.metadata.v1.MetadataService/Get', getRequest, params);
  check(getResponse, {
    'Get call succeeded': (r) => r && r.status === grpc.StatusOK,
  });

  console.warn('getResponse', getResponse.message.meta.imageUrl)
}
