import * as grpc from '@grpc/grpc-js'

import statsServer from './proxy/infrastructure/rpc/stats'

const GRPC_PORT: string | number = process.env.GRPC_PORT || 50051

type StartServerType = () => void

export const startServer: StartServerType = (): void => {
  // create a new gRPC server
  const server: grpc.Server = new grpc.Server();

  // register all the handler here...
  server.addService(statsServer.service, statsServer.handler)

  // define the host/port for server
  server.bindAsync(`0.0.0.0:${GRPC_PORT}`, grpc.ServerCredentials.createInsecure(), (err: Error | null, port: number) => {
    if (err != null) {
      return console.error(err)
    }
    console.info(`gRPC listening on ${port}`)

    // start the gRPC server
    server.start()
  })
}

startServer()

