import { useQuery, useQueryClient, type QueryKey, type UseQueryResult } from '@tanstack/react-query'
import type { Session } from '@ory/client'

import { fetchLinks, fetchLinksList, fetchProfile, fetchSession, searchLinks } from './api'

export const queryKeys = {
  session: ['session'] as const,
  profile: (userId: string) => ['profile', userId] as const,
  links: (userId?: string, filter?: string) =>
    ['links', { userId: userId ?? 'all', filter: filter ?? 'none' }] as const,
  linksList: (userId?: string) => ['links', 'list', userId ?? 'all'] as const,
  searchLinks: (query: string) => ['links', 'search', query] as const,
}

export function useSessionQuery(): UseQueryResult<Session> {
  return useQuery<Session>({
    queryKey: queryKeys.session,
    queryFn: () => fetchSession(),
    staleTime: 5 * 60 * 1000,
    retry: 1,
  }) as UseQueryResult<Session>
}

export function useOptionalSessionQuery(): UseQueryResult<Session | null> {
  return useQuery<Session | null>({
    queryKey: queryKeys.session,
    queryFn: async () => {
      try {
        return await fetchSession()
      } catch {
        return null
      }
    },
    staleTime: 5 * 60 * 1000,
    retry: false,
  }) as UseQueryResult<Session | null>
}

export function useProfileQuery(userId: string, enabled: boolean = true): UseQueryResult<any> {
  return useQuery<any>({
    queryKey: queryKeys.profile(userId),
    queryFn: ({ signal }) => fetchProfile(userId, signal),
    enabled: Boolean(userId) && enabled,
    staleTime: 2 * 60 * 1000,
    retry: 1,
  }) as UseQueryResult<any>
}

export function useLinksQuery(userId?: string, filter?: string, enabled: boolean = true): UseQueryResult<any> {
  return useQuery<any>({
    queryKey: queryKeys.links(userId, filter),
    queryFn: ({ signal }) => fetchLinks(userId, filter, signal),
    enabled,
    staleTime: 60 * 1000,
    retry: 1,
  }) as UseQueryResult<any>
}

export function useLinksListQuery(userId?: string, enabled: boolean = true): UseQueryResult<any[]> {
  return useQuery<any[]>({
    queryKey: queryKeys.linksList(userId),
    queryFn: ({ signal }) => fetchLinksList(userId, signal),
    enabled,
    staleTime: 60 * 1000,
    retry: 1,
  }) as UseQueryResult<any[]>
}

export function useSearchLinksQuery(query: string, enabled: boolean = true): UseQueryResult<any[]> {
  const trimmed = query.trim()

  return useQuery<any[]>({
    queryKey: queryKeys.searchLinks(trimmed),
    queryFn: ({ signal }) => searchLinks(trimmed, signal),
    enabled: Boolean(trimmed) && enabled,
    staleTime: 30 * 1000,
    retry: 1,
  }) as UseQueryResult<any[]>
}

export function useInvalidateCache() {
  const queryClient = useQueryClient()

  return {
    invalidate: (queryKey: QueryKey) => queryClient.invalidateQueries({ queryKey }),
    invalidatePattern: (pattern: RegExp) =>
      queryClient.invalidateQueries({
        predicate: (query) => pattern.test(JSON.stringify(query.queryKey)),
      }),
    clear: () => queryClient.clear(),
  }
}
