import * as React from 'react'
import Collapse from '@mui/material/Collapse'

import ExpandLessIcon from '@mui/icons-material/ExpandLess'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'

const iconClassName =
  'w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
const linkClassName =
  'flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group'

// @ts-ignore
const CollapsibleMenu = ({ icon, title, children }) => {
  const [isOpen, setIsOpen] = React.useState(false)

  const toggleCollapse = () => setIsOpen(!isOpen)

  return (
    <li>
      <div
        className="flex items-center w-full p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700"
        onClick={toggleCollapse}
      >
        {React.createElement(icon, { className: iconClassName })}
        <span className="flex-1 ms-3 text-left rtl:text-right whitespace-nowrap cursor-pointer">
          {title}
        </span>
        {isOpen ? <ExpandLessIcon /> : <ExpandMoreIcon />}
      </div>
      <Collapse in={isOpen} timeout="auto" unmountOnExit>
        <ul className="py-2 px-4 space-y-2">{children}</ul>
      </Collapse>
    </li>
  )
}

export default CollapsibleMenu
