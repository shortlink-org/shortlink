import { inject, injectable } from 'inversify'
import { PrismaClient } from '@prisma/client'

import { Stats } from "../../proto/domain/proxy/v1/proxy_pb";
import TYPES from '../../types'
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";

@injectable()
export class StatsService {
  constructor(
    @inject(TYPES.REPOSITORY.StatsRepositoryImpl) private readonly prisma: PrismaClient,
  ) {}

  public get(hash: string): Promise<Stats> {
    return new Promise<Stats>((resolve, reject) => {
      let stats = new Stats()

      const resp = this.prisma.stats.findUnique({
        where: {
          hash: hash,
        }
      }).then(data => {
        // @ts-ignore
        stats.setCountRedirect(data.count_redirect)

        const timestamp = new Timestamp()
        // @ts-ignore
        timestamp.fromDate(data.updated_at)
        stats.setUpdatedAt(timestamp)
      }).catch(
        // TODO: error
      )
    })
  }
}
