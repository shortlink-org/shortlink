import { Meta } from '@storybook/react'

import Sidebar from './Sidebar'

const meta: Meta<typeof Sidebar> = {
  title: 'Page/Sidebar',
  component: Sidebar,
}

export default meta

export const Default = {
  render: (args: any) => <Sidebar {...args} />,
  argTypes: {
    mode: {
      control: {
        type: 'select',
        options: ['full', 'mini'],
      },
    },
  },
}
