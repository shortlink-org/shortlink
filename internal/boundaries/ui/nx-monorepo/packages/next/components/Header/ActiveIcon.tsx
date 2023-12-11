// @ts-ignore
import ListItemIcon, { ListItemIconProps } from '@mui/material/ListItemIcon'
import { useRouter } from 'next/router'
import React, { useState, useEffect, ReactElement, Children } from 'react'
import { UrlObject } from 'url'

declare type Url = string | UrlObject

type ActiveLinkProps = ListItemIconProps & {
  href: Url
  children: ReactElement
  activeClassName?: string
}

function ActiveLink({ children, activeClassName, ...props }: ActiveLinkProps) {
  const { asPath, isReady } = useRouter()

  const child = Children.only(children)
  const childClassName = child.props.className || ''
  const [className, setClassName] = useState(childClassName)

  useEffect(() => {
    // Check if the router fields are updated client-side
    if (isReady) {
      // Dynamic route will be matched via props.as
      // Static route will be matched via props.href
      const linkPathname = new URL(props.href as string, location.href).pathname

      // Using URL().pathname to get rid of query and hash
      const activePathname = new URL(asPath, location.href).pathname

      const newClassName =
        linkPathname === activePathname
          ? `${childClassName} ${activeClassName}`.trim()
          : childClassName

      if (newClassName !== className) {
        setClassName(newClassName)
      }
    }
  }, [
    asPath,
    isReady,
    props.href,
    childClassName,
    activeClassName,
    setClassName,
    className,
  ])

  return (
    <ListItemIcon {...props}>
      {React.cloneElement(child, {
        className: className || null,
      })}
    </ListItemIcon>
  )
}

export default ActiveLink
