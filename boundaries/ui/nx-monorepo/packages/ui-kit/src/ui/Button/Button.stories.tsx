import Button from '@mui/material-next/Button'
import { Meta } from '@storybook/react'
import { fn, expect } from '@storybook/test'
import { within, userEvent } from '@storybook/testing-library'

const meta: Meta<typeof Button> = {
  title: 'UI/Button',
  component: Button,
  args: {
    onClick: fn(),
  },
}

export default meta

function Template(args: any) {
  return <Button {...args}>Text</Button>
}

export const Default = {
  render: Template,
  args: {},
  // @ts-ignore
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    const button = canvas.getByRole('button', { name: 'Text' })
    await userEvent.click(button)
    await expect(button).toBeEnabled()
  },
}
