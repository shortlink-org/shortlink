import { injectable } from 'inversify'
import fetch from 'node-fetch'

interface Link {
  url: string,
}

@injectable()
export class LinkService {
  public async get(hash: string): Promise<string> {
    // TODO: use gRPC
    const resp = await fetch(`${process.env.API_LINK_SERVICE}/api/link/${hash}`)
    const link: Link = await resp.json()

    return link.url
  }
}
