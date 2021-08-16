import { inject, injectable } from 'inversify'
import { PrismaClient } from '@prisma/client'
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {Logger} from "tslog"
import * as Amqp from "amqp-ts"

import AMQPController from '../infrastructure/amqp/amqp'
import { Stats } from "../../proto/domain/proxy/v1/proxy_pb";
import { Link } from "../../proto/domain/link/v1/link_pb";
import TYPES from '../../types'
import StatsRepository from "../infrastructure/store";
import {MQ_EVENT_LINK_NEW} from "../domain/event";

const log: Logger = new Logger()

@injectable()
export class StatsService {
  constructor(
    @inject(TYPES.REPOSITORY.StatsRepository) private readonly store: StatsRepository,
    @inject(TYPES.TAGS.AMQPController) private readonly amqp: AMQPController,
  ) {
    const exchange = this.amqp.connection.declareExchange(MQ_EVENT_LINK_NEW, "fanout", {durable: false})
    let queue = this.amqp.connection.declareQueue("shortlink-proxy")
    queue.bind(exchange)
    queue.activateConsumer((message: Amqp.Message) => {
      let link = Link.deserializeBinary(message.content)
      log.info(link)
      this.create(link)
    })
  }

  public get(hash: string): Promise<Stats> {
    return new Promise<Stats>((resolve, reject) => {
      let stats = new Stats()

      // const resp = this.prisma.stats.findUnique({
      //   where: {
      //     hash: hash,
      //   }
      // }).then(data => {
      //   // @ts-ignore
      //   stats.setCountRedirect(data.count_redirect)
      //
      //   const timestamp = new Timestamp()
      //   // @ts-ignore
      //   timestamp.fromDate(data.updated_at)
      //   stats.setUpdatedAt(timestamp)
      // }).catch(
      //   // TODO: error
      // )
    })
  }

  public create(payload: Link) {
    try {
      this.store.create(payload)
    } catch (error) {
      log.error(error)
    }
  }
}
