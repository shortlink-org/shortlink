import Box from '@mui/material/Box'
import Typography from '@mui/material/Typography'
import * as React from 'react'

interface TabPanelProps {
  children?: React.ReactNode
  dir?: string
  index: number
  value: number
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props

  if (value === index) {
    return <div className="w-4xl m-auto">{children}</div>
  }

  return null
}

export default TabPanel
