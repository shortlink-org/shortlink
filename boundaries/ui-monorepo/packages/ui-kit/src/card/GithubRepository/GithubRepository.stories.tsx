import { Meta } from '@storybook/react'
import { expect } from '@storybook/test'
import { within } from '@storybook/test'

import GithubRepository from './GithubRepository'

const meta: Meta<typeof GithubRepository> = {
  title: 'Card/GithubRepository',
  component: GithubRepository,
}

export default meta

function Template(args: any) {
  return (
    <GithubRepository
      title="GitHub Repository"
      url="https://github.com/shortlink-org/shortlink"
      {...args}
    />
  )
}

export const Default = {
  render: Template,
  args: {},
  // @ts-ignore
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    const link = canvas.getByRole('link')
    await expect(link).toHaveAttribute(
      'href',
      'https://github.com/shortlink-org/shortlink',
    )
  },
}
