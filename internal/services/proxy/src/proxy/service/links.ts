import { inject, injectable } from 'inversify'
import fetch from 'node-fetch'

import TYPES from '../../types'

interface Link {
  url: string,
}

@injectable()
export class LinkService {
  public async get(hash: string): Promise<string> {
    // TODO: use ENV
    // TODO: use gRPC
    const resp = await fetch(`http://localhost:7070/api/link/${hash}`)
    const link: Link = await resp.json()

    return link.url
  }
}
