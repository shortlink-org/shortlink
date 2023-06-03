import { usePathname } from 'next/navigation'
import ListItemIcon, { ListItemIconProps } from '@mui/material/ListItemIcon'
import React, { useState, useEffect, ReactElement, Children } from 'react'
import { UrlObject } from 'url'

declare type Url = string | UrlObject

type ActiveLinkProps = ListItemIconProps & {
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
    <ListItemIcon {...props}>
      {React.cloneElement(child, {
        className: newClassName || null,
      })}
    </ListItemIcon>
  )
}

export default ActiveLink
