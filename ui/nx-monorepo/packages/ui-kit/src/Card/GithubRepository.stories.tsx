// @ts-ignore
import React, {useContext, useEffect, useState} from 'react'
import { StoryObj, Meta, Preview } from '@storybook/react'

import GithubRepository  from './GithubRepository'

const meta: Meta<any> = {
  title: 'UI/GithubRepository',
  component: GithubRepository,
}

export default meta

const Template = (args) => (
  <GithubRepository
    title={'GitHub Repository'}
    url={'https://github.com/shortlink-org/shortlink'}
    {...args}
  />
)

export const Default = Template.bind({});
Default.args = {};
