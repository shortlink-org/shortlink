import { Meta } from '@storybook/react'
import { action } from '@storybook/addon-actions'

import SearchForm from './SearchForm'

const meta: Meta<typeof SearchForm> = {
  title: 'UI/Input/SearchForm',
  component: SearchForm,
  args: {
    // @ts-ignore
    onClick: action('clicked'),
  },
}

export default meta

function Template(args: any) {
  return (
    <div style={{ height: '300vh', position: 'relative' }}>
      {/* SearchForm will be positioned relative to this container */}
      <SearchForm {...args} />
    </div>
  )
}

export const Default = {
  render: Template,
  args: {},
}
