// @ts-ignore
import React, { useState } from 'react'
import { StoryObj, Meta, Preview } from '@storybook/react'

import ToggleDarkMode  from './ToggleDarkMode'
import { ColorModeContext } from './ColorModeContext'

const meta: Meta<any> = {
  title: 'ToggleDarkMode',
  component: ToggleDarkMode,
  argTypes: {},
  decorators: [
    Story => {
      const [darkMode, setDarkMode] = useState(ColorModeContext)

      return (
        <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
          <Story />
        </ColorModeContext.Provider>
      )
    },
  ],
}

export default meta

const Template = (args) => <ToggleDarkMode {...args} />

export const Default = Template.bind({});
Default.args = {};
