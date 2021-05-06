import 'dotenv/config';
import * as grpc from 'grpc'
import GrpcMiddleware, { GrpcCall } from 'grpc-ts-middleware'

import proxyHandler from './handlers/proxy'

const GRPC_PORT: string | number = process.env.GRPC_PORT || 50051

type StartServerType = () => void
export const startServer: StartServerType = (): void => {
  // create a new gRPC server
  const server: grpc.Server = new grpc.Server();

  // register all the handler here...
  server.addService(proxyHandler.service, proxyHandler.handler)

  // define the host/port for server
  server.bindAsync(`0.0.0.0:${GRPC_PORT}`, grpc.ServerCredentials.createInsecure(), (err: Error | null, port: number) => {
    if (err != null) {
      return console.error(err)
    }
    console.info(`gRPC listening on ${port}`)
  })

  // start the gRPC server
  server.start()
}

startServer()

