import React from 'react'

interface TabPanelProps {
  children?: React.ReactNode
  dir?: string
  index: number
  value: number
}

const TabPanel: React.FC<TabPanelProps> = ({ children, value, index, ...other }) => {
  if (value !== index) {
    return null
  }

  return (
    <div className="max-w-4xl mx-auto m-6" role="tabpanel" id={`tabpanel-${index}`} aria-labelledby={`tab-${index}`} {...other}>
      {children}
    </div>
  )
}

export default TabPanel
