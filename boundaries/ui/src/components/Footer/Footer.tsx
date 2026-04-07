/* tslint:disable */

import { Footer as UiKitFooter } from '@shortlink-org/ui-kit'
import type { FC } from 'react'

import Copyright from '../Copyright'

import { FooterNavLink } from './FooterNavLink'

// Define the type for the props
interface FooterProps {
  className?: string // Optional prop
}

const Footer: FC<FooterProps> = ({ className = '' }) => (
  <UiKitFooter
    className={className}
    contained={false}
    contentClassName="!max-w-none w-full px-5 sm:px-8 lg:px-10"
    LinkComponent={FooterNavLink}
    copyright={<Copyright />}
    description="Create, share, and track short links in one place."
    socialLinks={[]}
  />
)

export default Footer
