import { Meta, StoryFn } from '@storybook/react'
import { expect } from '@storybook/test'
import { within } from '@storybook/test'

import GithubRepository, { GithubRepositoryProps } from './GithubRepository'
import { JSX } from 'react/jsx-runtime'

const meta: Meta<typeof GithubRepository> = {
  title: 'Card/GithubRepository',
  component: GithubRepository,
  argTypes: {
    title: {
      control: 'text',
      description: 'Title of the GitHub repository',
      defaultValue: 'GitHub Repository',
    },
    url: {
      control: 'text',
      description: 'URL of the GitHub repository',
      defaultValue: 'https://github.com/shortlink-org/shortlink',
    },
  },
  parameters: {
    docs: {
      description: {
        component:
          'A card component that links to a GitHub repository with enhanced accessibility and styling.',
      },
    },
  },
}

export default meta

const Template: StoryFn<GithubRepositoryProps> = (
  args: JSX.IntrinsicAttributes & GithubRepositoryProps,
) => <GithubRepository {...args} />

export const Default = Template.bind({})
Default.args = {
  title: 'GitHub Repository',
  url: 'https://github.com/shortlink-org/shortlink',
}

Default.parameters = {
  docs: {
    storyDescription: 'Default state of the GithubRepository component.',
  },
}

// @ts-ignore
Default.play = async ({ canvasElement }) => {
  const canvas = within(canvasElement)

  // Check if the link has the correct href attribute
  const link = canvas.getByRole('link', {
    name: /visit github repository github repository/i,
  })
  await expect(link).toHaveAttribute(
    'href',
    'https://github.com/shortlink-org/shortlink',
  )

  // Check if the title is rendered correctly
  const title = canvas.getByText('GitHub Repository')
  expect(title).toBeInTheDocument()

  // Check if the displayed URL is correct
  const displayedUrl = canvas.getByText('/shortlink-org/shortlink')
  expect(displayedUrl).toBeInTheDocument()
}
