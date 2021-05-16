import * as grpc from '@grpc/grpc-js'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';

import {Stats} from "../../../proto/domain/proxy/v1/proxy_pb";
import { StatsRequest, StatsResponse } from '../../../proto/infrastructure/rpc/v1/proxy_pb'
import { StatsService, IStatsServer } from '../../../proto/infrastructure/rpc/v1/proxy_grpc_pb'
import {injectable} from "inversify";

@injectable()
class StatsServer implements IStatsServer {
  [name: string]: grpc.UntypedHandleCall;
  /**
   * Return stats by use URL
   * @param url
   */
  stats = (url: grpc.ServerUnaryCall<StatsRequest, StatsResponse>): StatsResponse => {
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
  service: StatsService,       // Service interface
  handler: new StatsServer(), // Service interface definitions
}
