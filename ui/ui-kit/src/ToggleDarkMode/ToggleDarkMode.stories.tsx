// @ts-ignore
import React, { useState } from 'react'
import { Story, Meta } from '@storybook/react'
import { withReactContext } from 'storybook-react-context'

import ToggleDarkMode  from './ToggleDarkMode'
import { ColorModeContext } from './ColorModeContext'
import './styles.css'

export default {
  title: 'ToggleDarkMode',
  component: ToggleDarkMode,
  argTypes: {},
  decorators: [
    withReactContext({
      Context: ColorModeContext,
      initialState: false,
    }),
  ],
} as Meta<typeof React.Component>;

const Template: Story<any> = (args) => {
  const [darkMode, setDarkMode] = useState(false)

  return (
    <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
      <ToggleDarkMode {...args} />
    </ColorModeContext.Provider>
  )
}

export const Default = Template.bind({});
Default.args = {};
