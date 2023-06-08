// @ts-ignore
import React, { useContext, useEffect, useState } from 'react'
import { Meta } from '@storybook/react'

import Header from './Header'

const meta: Meta<any> = {
  title: 'Page/Header',
  component: Header,
}

export default meta

const Template = (args) => <Header title="Header" {...args} />

export const Default = Template.bind({})
Default.args = {}
