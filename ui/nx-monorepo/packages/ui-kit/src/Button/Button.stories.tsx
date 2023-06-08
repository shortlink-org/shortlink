// @ts-ignore
import React, { useContext, useEffect, useState } from 'react'
import { StoryObj, Meta, Preview } from '@storybook/react'

import Button from '@mui/material/Button'

const meta: Meta<any> = {
  title: 'UI/Button',
  component: Button,
}

export default meta

function Template(args) {
  return <Button {...args}>Text</Button>
}

export const Default = Template.bind({})
Default.args = {}
