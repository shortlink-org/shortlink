// @ts-ignore
import React from 'react'
import { Meta } from '@storybook/react'

import ToggleDarkMode  from './ToggleDarkMode'
import Header  from '../Header/Header'

const meta: Meta<any> = {
  title: 'UI/ToggleDarkMode',
  component: ToggleDarkMode,
}

export default meta

const Template = (args) => <ToggleDarkMode {...args} />

export const Default = Template.bind({});
Default.args = {};

export const WithHeader = () => {
  return (
    <>
      <Header title={'Header'} />
      <ToggleDarkMode />
    </>
  )
}
