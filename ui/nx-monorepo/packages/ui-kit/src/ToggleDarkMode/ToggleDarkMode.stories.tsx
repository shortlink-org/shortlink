// @ts-ignore
import * as React from 'react'
import { Meta } from '@storybook/react'

import ToggleDarkMode from './ToggleDarkMode'
import Header from '../Header/Header'

const meta: Meta<any> = {
  title: 'UI/ToggleDarkMode',
  component: ToggleDarkMode,
}

export default meta

function Template(args) {
  return <ToggleDarkMode {...args} />
}

export const Default = Template.bind({})
Default.args = {}

export function WithHeader() {
  return (
    <>
      <Header title="Header" />
      <ToggleDarkMode />
    </>
  )
}
