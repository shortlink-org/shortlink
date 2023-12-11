import { injectable } from 'inversify'

interface Link {
  url: string,
}

@injectable()
export class LinkService {
  public async get(hash: string): Promise<string> {
    // TODO: use gRPC
    try {
      const resp = await fetch(`${process.env.API_LINK_SERVICE}/api/links/${encodeURIComponent(hash)}`)
      if (!resp.ok) {
        throw new Error(`${resp.status} ${resp.statusText}`)
      }

      // @ts-ignore
      const link: Link = await resp.json()
      return link.url
    } catch (err) {
      // @ts-ignore
      throw new Error(err)
    }
  }
}
