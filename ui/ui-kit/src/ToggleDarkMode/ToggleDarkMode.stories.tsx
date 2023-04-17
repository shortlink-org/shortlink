// @ts-ignore
import React, { useState } from 'react'
import { StoryObj, Meta, Preview } from '@storybook/react'

import ToggleDarkMode  from './ToggleDarkMode'

const meta: Meta<any> = {
  title: 'UI/ToggleDarkMode',
  component: ToggleDarkMode,
}

export default meta

const Template = (args) => <ToggleDarkMode {...args} />

export const Default = Template.bind({});
Default.args = {};
