import { injectable } from 'inversify'
// @ts-ignore
const fetch = (...args) => import('node-fetch').then(({default: fetch}) => fetch(...args))

interface Link {
  url: string,
}

@injectable()
export class LinkService {
  public async get(hash: string): Promise<string> {
    // TODO: use gRPC
    const resp = await fetch(`${process.env.API_LINK_SERVICE}/api/links/${hash}`)
    const link: unknown = await resp.json()

    // @ts-ignore
    return link.url
  }
}
