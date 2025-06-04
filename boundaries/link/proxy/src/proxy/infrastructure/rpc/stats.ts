import * as grpc from '@grpc/grpc-js'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';

import {Stats} from "../../../proto/domain/proxy/v1/proxy_pb";
import { StatsRequest, StatsResponse } from '../../../proto/infrastructure/rpc/proxy/v1/proxy_pb'
import { StatsServiceClient } from '../../../proto/infrastructure/rpc/proxy/v1/proxy_grpc_pb'
import {injectable} from "inversify";

@injectable()
class StatsServer {
  [name: string]: grpc.UntypedHandleCall;
  /**
   * Return stats by use URL
   * @param url
   */
  stats = (url: grpc.ServerUnaryCall<StatsRequest, StatsResponse>): StatsResponse => {
    const resp: StatsResponse = new StatsResponse()

    console.info(`hash: ${url.request.hash}`)

    let stats = new Stats()
    // @ts-ignore
    stats.countRedirect = 0

    const timestamp = new Timestamp();
    timestamp.fromDate(new Date());
    // @ts-ignore
    stats.updatedAt = timestamp

    resp.stats = stats
    return resp
  }
}

export default {
  service: StatsServiceClient.service, // Service interface
  handler: new StatsServer(),  // Service interface definitions
}
