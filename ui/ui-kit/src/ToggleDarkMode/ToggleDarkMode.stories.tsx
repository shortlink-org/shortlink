// @ts-ignore
import React from 'react'
import { Story, Meta } from '@storybook/react'

import ToggleDarkMode  from './ToggleDarkMode'

export default {
  title: 'ToggleDarkMode',
  component: ToggleDarkMode,
  argTypes: {},
} as Meta<typeof React.Component>;

const Template: Story<any> = (args) => <ToggleDarkMode {...args} />;

export const DarkMode = Template.bind({});
DarkMode.args = {};
