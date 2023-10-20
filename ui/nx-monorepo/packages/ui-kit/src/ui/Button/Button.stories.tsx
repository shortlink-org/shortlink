import Button from '@mui/material/Button'
import { Meta } from '@storybook/react'

const meta: Meta<any> = {
  title: 'UI/Button',
  component: Button,
}

export default meta

function Template(args: any) {
  return <Button {...args}>Text</Button>
}

export const Default = {
  render: Template,
  args: {},
}
