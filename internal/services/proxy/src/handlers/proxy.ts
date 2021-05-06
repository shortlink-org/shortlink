import * as grpc from '@grpc/grpc-js'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';

import {Stats} from "../proto/domain/proxy/v1/proxy_pb";
import { StatsRequest, StatsResponse } from '../proto/infrastructure/rpc/v1/proxy_pb'
import { ProxyService, IProxyServer } from '../proto/infrastructure/rpc/v1/proxy_grpc_pb'

class ProxyHandler implements IProxyServer {
  [name: string]: grpc.UntypedHandleCall;
  /**
   * Return stats by use URL
   * @param url
   */
  stats = (url: grpc.ServerUnaryCall<StatsRequest, StatsResponse>): StatsResponse => {
    console.log(123)
    const resp: StatsResponse = new StatsResponse()

    console.info(`hash: ${url.request.getHash()}`)

    let stats = new Stats()
    stats.setCountRedirect(0)

    const timestamp = new Timestamp();
    timestamp.fromDate(new Date());
    stats.setUpdatedAt(timestamp)

    resp.setStats(stats)
    return resp
  }
}

export default {
  service: ProxyService,       // Service interface
  handler: new ProxyHandler(), // Service interface definitions
}
