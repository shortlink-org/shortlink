/* tslint:disable */

// @ts-ignore
import { Footer as UiKitFooter } from '@shortlink-org/ui-kit'
import Link from 'next/link'

import Copyright from '../Copyright'

// Define the type for the props
interface FooterProps {
  className?: string // Optional prop
}

const Footer: React.FC<FooterProps> = ({ className = '' }) => (
  <UiKitFooter className={className} LinkComponent={Link} copyright={<Copyright />} />
)

export default Footer
