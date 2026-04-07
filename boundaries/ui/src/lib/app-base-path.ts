/** Mirrors `basePath` in `next.config.js` — used for pathname matching in nav. */
export const APP_BASE_PATH = '/next' as const

export function navPathToPathname(href: string): string {
  if (href === '/') return APP_BASE_PATH
  const normalized = href.startsWith('/') ? href : `/${href}`
  return `${APP_BASE_PATH}${normalized}`
}

export function navItemIsActive(pathname: string | null, itemPath: string) {
  const current = (pathname ?? '').replace(/\/$/, '')
  const target = navPathToPathname(itemPath).replace(/\/$/, '')
  if (itemPath === '/') return current === target
  return current === target || current.startsWith(`${target}/`)
}
