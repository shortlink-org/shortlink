import { Meta } from '@storybook/react'

import Sidebar from './Sidebar'

const meta: Meta<typeof Sidebar> = {
  title: 'Page/Sidebar',
  component: Sidebar,
}

export default meta

function Template(args: any) {
  return <Sidebar mode={'full'} {...args} />
}

export const Default = {
  render: Template,
  args: {},
}
