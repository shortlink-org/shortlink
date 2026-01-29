import { LinkAction, LinkDomainError, LinkErrorCode } from './link.types'

type RestErrorMessage = {
  code?: string
  title?: string
  detail?: string
  action?: string
  desc?: string
  message?: string
}

type RestErrorPayload = RestErrorMessage & {
  error?: string
  messages?: RestErrorMessage[]
}

const FALLBACK_TITLE = 'Unable to create link'
const FALLBACK_DETAIL = 'Please try again later.'

const mapCode = (code?: string): LinkErrorCode => {
  switch (code) {
    case 'SESSION_NOT_FOUND':
      return 'SESSION_NOT_FOUND'
    case 'USER_NOT_IDENTIFIED':
      return 'USER_NOT_IDENTIFIED'
    case 'SESSION_METADATA_MISSING':
      return 'SESSION_METADATA_MISSING'
    case 'PERMISSION_DENIED':
      return 'PERMISSION_DENIED'
    case 'INVALID_TOKEN':
      return 'INVALID_TOKEN'
    case 'SERVICE_UNAVAILABLE':
      return 'SERVICE_UNAVAILABLE'
    case 'UNKNOWN':
    default:
      return 'UNKNOWN'
  }
}

const mapAction = (action?: string): LinkAction => {
  switch (action) {
    case 'LOGIN':
      return 'LOGIN'
    case 'RETRY':
      return 'RETRY'
    default:
      return 'NONE'
  }
}

const coerceMessage = (candidate: RestErrorMessage | null | undefined): RestErrorMessage | null => {
  if (!candidate) {
    return null
  }

  const detail = candidate.detail ?? candidate.message ?? candidate.desc
  if (!detail && !candidate.title) {
    return null
  }

  return {
    code: candidate.code,
    title: candidate.title,
    detail,
    action: candidate.action,
  }
}

const extractPrimaryMessage = (payload: RestErrorPayload | null | undefined): RestErrorMessage | null => {
  if (!payload) {
    return null
  }

  const candidates: (RestErrorMessage | null)[] = [
    coerceMessage({
      code: payload.code,
      title: payload.title,
      detail: payload.detail ?? payload.message ?? payload.error ?? payload.desc,
      action: payload.action,
    }),
  ]

  if (Array.isArray(payload.messages)) {
    candidates.push(...payload.messages.map((message) => coerceMessage(message)))
  }

  return candidates.find((candidate) => candidate) ?? null
}

export const mapRestErrorToDomainError = (payload: RestErrorPayload | null | undefined): LinkDomainError => {
  try {
    const primary = extractPrimaryMessage(payload)

    return {
      code: mapCode(primary?.code),
      title: primary?.title ?? FALLBACK_TITLE,
      detail: primary?.detail ?? FALLBACK_DETAIL,
      action: mapAction(primary?.action),
    }
  } catch (error) {
    console.error('Failed to map link error payload', error)
    return {
      code: 'UNKNOWN',
      title: FALLBACK_TITLE,
      detail: FALLBACK_DETAIL,
      action: 'NONE',
    }
  }
}

export const networkError = (): LinkDomainError => ({
  code: 'NETWORK_ERROR',
  title: 'No network connection',
  detail: 'Could not reach the server. Check your connection and try again.',
  action: 'RETRY',
})

export const invalidResponseError = (): LinkDomainError => ({
  code: 'INVALID_RESPONSE',
  title: FALLBACK_TITLE,
  detail: FALLBACK_DETAIL,
  action: 'NONE',
})
