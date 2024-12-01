import { Meta } from '@storybook/react'
import { action } from '@storybook/addon-actions'

import ScrollToTopButton from './ScrollToTopButton'

const meta: Meta<typeof ScrollToTopButton> = {
  title: 'UI/ScrollToTopButton',
  component: ScrollToTopButton,
  args: {
    // @ts-ignore
    onClick: action('clicked'),
  },
}

export default meta

function Template(args: any) {
  return (
    <div style={{ height: '300vh', position: 'relative' }}>
      {/* ScrollToTopButton will be positioned relative to this container */}
      <ScrollToTopButton {...args} />
    </div>
  )
}

export const Default = {
  render: Template,
  args: {},
}
