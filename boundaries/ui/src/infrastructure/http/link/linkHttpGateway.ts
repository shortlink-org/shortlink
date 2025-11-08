import { CreateLinkCommand, CreateLinkResult, LinkGateway } from '@/domain/link/link.types'
import { invalidResponseError, mapRestErrorToDomainError, networkError } from '@/domain/link/link.errors'

type RestSuccessPayload = {
  hash?: string
  link?: {
    hash?: string
    url?: string
    describe?: string
  }
  url?: string
  describe?: string
}

const ENDPOINT = '/api/links'

const parseJsonSafe = async (response: Response) => {
  try {
    return await response.json()
  } catch {
    return null
  }
}

const extractHash = (payload: RestSuccessPayload | null): string | null => {
  if (!payload) {
    return null
  }

  if (payload.hash) {
    return payload.hash
  }

  if (payload.link?.hash) {
    return payload.link.hash
  }

  return null
}

export class LinkHttpGateway implements LinkGateway {
  async createLink(command: CreateLinkCommand): Promise<CreateLinkResult> {
    try {
      const response = await fetch(ENDPOINT, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          url: command.url,
          describe: command.describe,
        }),
      })

      const payload = await parseJsonSafe(response)

      if (response.ok) {
        const hash = extractHash(payload)

        if (!hash) {
          return {
            kind: 'failure',
            error: invalidResponseError(),
          }
        }

        return {
          kind: 'success',
          link: {
            hash,
            url: command.url,
            describe: command.describe,
          },
        }
      }

      return {
        kind: 'failure',
        error: mapRestErrorToDomainError(payload),
      }
    } catch (error) {
      console.error('Failed to create link', error)
      return {
        kind: 'failure',
        error: networkError(),
      }
    }
  }
}
