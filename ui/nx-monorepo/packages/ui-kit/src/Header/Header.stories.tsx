import { Meta } from '@storybook/react'

import Header from './Header'

const meta: Meta<any> = {
  title: 'Page/Header',
  component: Header,
}

export default meta

function Template(args: any) {
  return <Header title="Header" {...args} />
}

export const Default = Template.bind({})
// @ts-ignore
Default.args = {}
