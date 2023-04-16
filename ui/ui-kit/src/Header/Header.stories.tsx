// @ts-ignore
import React, {useContext, useEffect, useState} from 'react'
import { StoryObj, Meta, Preview } from '@storybook/react'

import Header  from './Header'
import { ColorModeContext } from '../theme/ColorModeContext'
import {useTheme as nextUseTheme} from "next-themes";

const meta: Meta<any> = {
  title: 'Header',
  component: Header,
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

const Template = (args) => <Header title={'Header'} {...args} />

export const Default = Template.bind({});
Default.args = {};
