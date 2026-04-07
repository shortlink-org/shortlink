'use client'

import NextLink from 'next/link'
import { usePathname } from 'next/navigation'
import { forwardRef } from 'react'
import type { ComponentPropsWithRef } from 'react'

import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'

/**
 * Footer links need client pathname to pick `nav-forward` / `nav-back` / `nav-sibling`.
 */
export const FooterNavLink = forwardRef<HTMLAnchorElement, ComponentPropsWithRef<typeof NextLink>>(
  function FooterNavLink(props, ref) {
    const pathname = usePathname()
    const transitionTypes = props.transitionTypes ?? resolveNavTransitionTypes(pathname, props.href)
    return <NextLink ref={ref} {...props} transitionTypes={transitionTypes} />
  },
)
