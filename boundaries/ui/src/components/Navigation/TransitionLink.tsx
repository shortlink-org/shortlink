'use client'

import type { LinkProps } from 'next/link'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import type { ReactNode } from 'react'

import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'

export type TransitionLinkProps = LinkProps & {
  children: ReactNode
}

/**
 * App Router link with shared `transitionTypes` for View Transitions.
 * Uses native `<Link>` navigation (do not wrap with `preventDefault` + `router.push` — that skips transition types).
 */
export function TransitionLink({ children, transitionTypes, href, ...props }: TransitionLinkProps) {
  const pathname = usePathname()
  return (
    <Link href={href} {...props} transitionTypes={transitionTypes ?? resolveNavTransitionTypes(pathname, href)}>
      {children}
    </Link>
  )
}

export default TransitionLink
