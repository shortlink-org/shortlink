export type ShortLink = {
  hash: string
  url: string
  describe?: string
}

export type CreateLinkCommand = {
  url: string
  describe?: string
  allowed_emails?: string[]
}

export type LinkErrorCode =
  | 'SESSION_NOT_FOUND'
  | 'USER_NOT_IDENTIFIED'
  | 'SESSION_METADATA_MISSING'
  | 'UNKNOWN'
  | 'NETWORK_ERROR'
  | 'INVALID_RESPONSE'

export type LinkAction = 'LOGIN' | 'RETRY' | 'NONE'

export type LinkDomainError = {
  code: LinkErrorCode
  title: string
  detail: string
  action: LinkAction
}

export type CreateLinkResult =
  | {
      kind: 'success'
      link: ShortLink
    }
  | {
      kind: 'failure'
      error: LinkDomainError
    }

export interface LinkGateway {
  createLink(command: CreateLinkCommand): Promise<CreateLinkResult>
}
