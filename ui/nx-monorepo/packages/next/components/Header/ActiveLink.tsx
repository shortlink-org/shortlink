import { usePathname } from 'next/navigation'
import Link, { LinkProps } from 'next/link'
import React, { ReactElement, Children } from 'react'
import { UrlObject } from 'url'

declare type Url = string | UrlObject

type ActiveLinkProps = LinkProps & {
  href: Url
  children: ReactElement
  activeClassName?: string
}

const ActiveLink = ({
  children,
  activeClassName,
  ...props
}: ActiveLinkProps) => {
  const pathname = usePathname()

  const child = Children.only(children)
  const childClassName = child.props.className || ''

  const isActive = pathname.startsWith(props.href as string)
  const newClassName = isActive
    ? `${childClassName} ${activeClassName}`.trim()
    : childClassName

  return (
    <Link {...props} legacyBehavior>
      {React.cloneElement(child, {
        className: newClassName || null,
      })}
    </Link>
  )
}

export default ActiveLink
