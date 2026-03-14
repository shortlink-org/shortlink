import clsx from 'clsx'
import React from 'react'

type PageSectionSize = 'default' | 'narrow' | 'wide' | 'full'

const sizeClassMap: Record<PageSectionSize, string> = {
  default: 'max-w-7xl',
  narrow: 'max-w-4xl',
  wide: 'max-w-[90rem]',
  full: 'max-w-none',
}

type PageSectionProps = {
  children: React.ReactNode
  className?: string
  size?: PageSectionSize
  as?: React.ElementType
}

export default function PageSection({
  children,
  className,
  size = 'default',
  as: Component = 'section',
}: PageSectionProps) {
  return <Component className={clsx('mx-auto w-full px-4 sm:px-6 lg:px-8', sizeClassMap[size], className)}>{children}</Component>
}
