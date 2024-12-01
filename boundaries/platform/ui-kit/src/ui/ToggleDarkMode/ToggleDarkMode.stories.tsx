import { Meta } from '@storybook/react'
import Button from '@mui/material/Button'

import ToggleDarkMode from './ToggleDarkMode'
import Header from '../../page/Header/Header'

const meta: Meta<typeof ToggleDarkMode> = {
  title: 'UI/ToggleDarkMode',
  component: ToggleDarkMode,
}

export default meta

function Template(args: any) {
  return <ToggleDarkMode {...args} />
}

export const Default = {
  render: Template,
  args: {},
}

export function WithHeader() {
  return (
    <>
      <Header title="Header" />
      <ToggleDarkMode id="toggle" />
      <Button variant="contained">Example MUI button</Button>
    </>
  )
}
