import { injectable, inject } from "inversify";
import { PrismaClient } from '@prisma/client'
import { Stats } from "../../../proto/domain/proxy/v1/proxy_pb";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";

import {Link} from "../../../proto/domain/link/v1/link_pb";

const prisma = new PrismaClient()

@injectable()
class StatsRepository {

  /**
   * Return stats by use URL
   * @param hash
   */
  get = async (hash: string): Promise<Stats> => {
    let resp = await prisma.stats.findUnique({
      where: {
        hash: hash,
      },
    })

    let stats = new Stats()

    if (resp !== null) {
      stats.setHash(resp.hash)
      stats.setCountRedirect(resp.count_redirect)

      const timestamp = new Timestamp();
      timestamp.fromDate(resp.updated_at);
      stats.setUpdatedAt(timestamp)

      stats.setUpdatedAt(timestamp)
    }

    return stats
  }

  list = async (): Promise<Array<Stats>> => {
    let list: Array<Stats> = []
    const resp = await prisma.stats.findMany()
    return list
  }

  create = async (payload: Link): Promise<Stats> => {
    // let resp = await prisma.stats.create({
    //   data: {
    //     hash: payload.getHash(),
    //   },
    // })

    let stats = new Stats()
    // stats.setHash(resp.hash)

    return stats
  }

  update = async (payload: Stats): Promise<Stats> => {
    await prisma.stats.update({
      where: {
        hash: payload.getHash(),
      },
      data: payload,
    })

    return payload
  }

  delete = async (hash: string): Promise<boolean> => {
    const resp = await prisma.stats.delete({
      where: {
        hash: hash,
      },
    })

    return true
  }
}

export default StatsRepository
