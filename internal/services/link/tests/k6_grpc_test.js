import grpc from 'k6/net/grpc';

const client = new grpc.Client();

client.load(['definitions'], './infrastructure/rpc/link/v1/link.proto');

export default () => {
  client.connect("localhost:50051", { timeout: "5s" });

  const data = { link: 'https://google.com' };
  const response = client.invoke('infrastructure.rpc.link.v1.LinkService/Add', data);

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
  });

  console.log(JSON.stringify(response.message));

  client.close();
  sleep(1);
}
