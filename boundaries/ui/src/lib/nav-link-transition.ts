/**
 * Labels for next/link `transitionTypes` → React.addTransitionType during App Router navigations.
 * Requires `experimental.viewTransition: true` in next.config.
 *
 * @see https://nextjs.org/docs/app/api-reference/config/next-config-js/viewTransition
 */

import type { UrlObject } from 'url'

import { APP_BASE_PATH, navPathToPathname } from '@/lib/app-base-path'

/** Deeper in the route tree (e.g. /links → /links/abc). */
export const NAV_TRANSITION_FORWARD = 'nav-forward'

/** Up the tree (e.g. /links/abc → /links or → /). */
export const NAV_TRANSITION_BACK = 'nav-back'

/** Same depth, different branch (e.g. /profile ↔ /links/search). */
export const NAV_TRANSITION_SIBLING = 'nav-sibling'

/** In-page hash navigation (`#section`). */
export const NAV_TRANSITION_ANCHOR = 'nav-anchor'

export type NavTransitionType =
  | typeof NAV_TRANSITION_FORWARD
  | typeof NAV_TRANSITION_BACK
  | typeof NAV_TRANSITION_SIBLING
  | typeof NAV_TRANSITION_ANCHOR

function normalizePath(path: string): string {
  const t = path.replace(/\/+$/, '')
  return t === '' ? '/' : t
}

function appRelative(pathname: string): string {
  const n = normalizePath(pathname)
  if (n === APP_BASE_PATH) return '/'
  const prefix = `${APP_BASE_PATH}/`
  if (n.startsWith(prefix)) {
    const rest = n.slice(APP_BASE_PATH.length)
    return rest === '' ? '/' : `/${rest.replace(/^\/+/, '')}`
  }
  return n
}

function segments(rel: string): string[] {
  return rel.split('/').filter(Boolean)
}

export function hrefPathForNav(href: string | UrlObject): string {
  if (typeof href === 'string') return href
  const pathname =
    href.pathname != null && typeof href.pathname === 'string' && href.pathname !== ''
      ? href.pathname
      : '/'
  return pathname
}

/**
 * Turnstile `href` (Next `Link` value) into app pathname (`usePathname()` shape), without query/hash.
 */
export function hrefToAppPathname(href: string | UrlObject): string {
  const pathPart = hrefPathForNav(href)
  if (pathPart.startsWith('#')) {
    return normalizePath(navPathToPathname('/'))
  }
  return normalizePath(navPathToPathname(pathPart.split('#')[0] ?? pathPart))
}

/**
 * Heuristic navigation direction from current pathname (including `basePath`) to a `Link` `href`.
 */
export function resolveNavTransitionTypes(
  currentPathname: string | null | undefined,
  href: string | UrlObject,
): NavTransitionType[] {
  const pathPart = hrefPathForNav(href)
  if (pathPart.startsWith('#')) {
    return [NAV_TRANSITION_ANCHOR]
  }

  const currentRel = appRelative(currentPathname ?? '')
  const targetRel = appRelative(hrefToAppPathname(href))

  if (currentRel === targetRel) {
    return [NAV_TRANSITION_SIBLING]
  }

  const a = segments(currentRel)
  const b = segments(targetRel)
  const minLen = Math.min(a.length, b.length)
  for (let i = 0; i < minLen; i++) {
    if (a[i] !== b[i]) {
      return [NAV_TRANSITION_SIBLING]
    }
  }
  if (b.length < a.length) {
    return [NAV_TRANSITION_BACK]
  }
  if (a.length < b.length) {
    return [NAV_TRANSITION_FORWARD]
  }
  return [NAV_TRANSITION_SIBLING]
}
