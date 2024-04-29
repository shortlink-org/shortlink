import { Meta } from '@storybook/react'

import Sidebar from './Sidebar'

const meta: Meta<typeof Sidebar> = {
  title: 'Page/Sidebar',
  component: Sidebar,
}

export default meta

export const Default = {
  render: (args: any) => {
    let className = 'h-screen grid grid-cols-4 gap-4'

    if (args.mode === 'mini') {
      className = 'h-screen grid grid-cols-12 gap-4'
    }

    return (
      <div className={className}>
        <Sidebar {...args} />
      </div>
    )
  },
  argTypes: {
    mode: {
      control: {
        type: 'select',
        options: ['full', 'mini'],
      },
    },
  },
}
