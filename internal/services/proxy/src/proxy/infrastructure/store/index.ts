import { injectable, inject } from "inversify";
import { Prisma, PrismaClient } from '@prisma/client'
import { Stats } from "../../../proto/domain/proxy/v1/proxy_pb";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";

import TYPES from '../../../types'

const prisma = new PrismaClient()

@injectable()
class StatsRepositoryImpl {
  // private statsRepository: StatsRepositoryImpl;
  // constructor(@inject(TYPES.StatsRepositoryImpl) statsRepository: StatsRepositoryImpl) {
  //   this.statsRepository = statsRepository;
  // }

  /**
   * Return stats by use URL
   * @param url
   */
  get = (url: string): Stats => {
    // const resp = await prisma.stats.findFirst()
    //
    let resp = new Stats()
    // stats.setCountRedirect(0)
    //
    // const timestamp = new Timestamp();
    // timestamp.fromDate(new Date());
    // stats.setUpdatedAt(timestamp)

    // TODO:
    // 1. return stats by hash

    return resp
  }

  list = (): Array<Stats> => {
    let list: Array<Stats> = []
    // TODO:
    // 1. return Array<Stats>
    return list
  }

  create = (payload: Stats): boolean => {
    // TODO:
    // 1. create Stats
    return true
  }

  update = (payload: Stats): boolean => {
    // TODO:
    // 1. update Stats
    return true
  }

  delete = (hash: String): boolean => {
    // TODO:
    // 1. delete Stats
    return true
  }
}
