/**
 * Data layer для API запросов (TanStack Query управляет кешем)
 */

import { Session } from '@ory/client'
import ory from '@/pkg/sdk'

/**
 * Fetch user session
 * Кешируется на 5 минут
 */
export async function fetchSession(): Promise<Session> {
  const { data } = await ory.toSession()
  return data
}

/**
 * Fetch user profile
 * Кешируется на 2 минуты
 */
export async function fetchProfile(userId: string, signal?: AbortSignal) {
  // Здесь будет ваш API call
  const response = await fetch(`/api/profile/${userId}`, { signal })
  if (!response.ok) {
    throw new Error('Failed to fetch profile')
  }
  return response.json()
}

/**
 * Fetch user links
 * Кешируется на 1 минуту
 */
export async function fetchLinks(userId?: string, filter?: string, signal?: AbortSignal) {
  const params = new URLSearchParams()
  if (userId) params.set('userId', userId)
  if (filter) params.set('filter', filter)

  const response = await fetch(`/api/links?${params}`, { signal })
  if (!response.ok) {
    throw new Error('Failed to fetch links')
  }
  return response.json()
}

/**
 * Search links
 * Короткий TTL для поиска
 */
export async function searchLinks(query: string, signal?: AbortSignal) {
  if (!query.trim()) {
    return []
  }

  const response = await fetch(`/api/links/search?q=${encodeURIComponent(query)}`, { signal })
  if (!response.ok) {
    throw new Error('Failed to search links')
  }
  return response.json()
}

/**
 * Fetch links list
 * Кешируется на 1 минуту
 */
export async function fetchLinksList(userId?: string, signal?: AbortSignal) {
  const params = new URLSearchParams()
  if (userId) params.set('userId', userId)

  const response = await fetch(`/api/links?${params}`, { signal })
  if (!response.ok) {
    throw new Error('Failed to fetch links')
  }
  const data = await response.json()
  return data.links || []
}

/**
 * Mutations (не кешируются, но возвращают промисы)
 */

export async function updateProfile(userId: string, data: any) {
  const response = await fetch(`/api/profile/${userId}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  })
  
  if (!response.ok) {
    throw new Error('Failed to update profile')
  }
  
  return response.json()
}

export async function createLink(data: any) {
  const response = await fetch('/api/links', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  })
  
  if (!response.ok) {
    throw new Error('Failed to create link')
  }
  
  return response.json()
}

export async function deleteLink(linkId: string) {
  const response = await fetch(`/api/links/${linkId}`, {
    method: 'DELETE'
  })
  
  if (!response.ok) {
    throw new Error('Failed to delete link')
  }
  
  return response.json()
}

export async function updateLink(linkId: string, data: any) {
  const response = await fetch(`/api/links/${linkId}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  })
  
  if (!response.ok) {
    throw new Error('Failed to update link')
  }
  
  return response.json()
}
