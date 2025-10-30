// @ts-ignore
import ListItemIcon, { ListItemIconProps } from '@mui/material/ListItemIcon'
import { usePathname } from 'next/navigation'
import React, { useState, useEffect, ReactElement, Children, isValidElement } from 'react'
import { UrlObject } from 'url'

declare type Url = string | UrlObject

type ActiveLinkProps = ListItemIconProps & {
  href: Url
  children: ReactElement
  activeClassName?: string
}

function ActiveLink({ children, activeClassName, ...props }: ActiveLinkProps) {
  const currentPath = usePathname()

  const child = Children.only(children)
  
  // Type guard to ensure child is a valid React element with props
  if (!isValidElement(child)) {
    throw new Error('ActiveLink expects a valid React element as children')
  }
  
  const childClassName = (child.props as { className?: string }).className || ''
  const [className, setClassName] = useState(childClassName)

  useEffect(() => {
    // Dynamic route will be matched via props.as
    // Static route will be matched via props.href
    const linkPathname = new URL(props.href as string, location.href).pathname

    // Using URL().pathname to get rid of query and hash
    const activePathname = new URL(currentPath, location.href).pathname

    const newClassName = linkPathname === activePathname ? `${childClassName} ${activeClassName}`.trim() : childClassName

    if (newClassName !== className) {
      setClassName(newClassName)
    }
  }, [currentPath, props.href, childClassName, activeClassName, setClassName, className])

  return (
    <ListItemIcon {...props}>
      {React.cloneElement(child, {
        className: className || null,
      } as React.HTMLAttributes<HTMLElement>)}
    </ListItemIcon>
  )
}

export default ActiveLink
