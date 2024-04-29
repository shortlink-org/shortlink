import * as React from 'react'
import Collapse from '@mui/material/Collapse'

import ExpandLessIcon from '@mui/icons-material/ExpandLess'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'

const iconClassName =
  'text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
let bodyClassName =
  'flex items-center cursor-pointer w-full p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700'

// @ts-ignore
const CollapsibleMenu = ({ mode, icon, title, children }) => {
  const [isOpen, setIsOpen] = React.useState(false)

  const toggleCollapse = () => setIsOpen(!isOpen)

  if (mode === 'mini') {
    bodyClassName += ' justify-center'
  }

  return (
    <li className={'w-full'}>
      <div className={bodyClassName} onClick={toggleCollapse}>
        {React.createElement(icon, { className: iconClassName })}

        {mode === 'full' && (
          <React.Fragment>
            <span className="flex-1 ms-3 text-left rtl:text-right whitespace-nowrap">
              {title}
            </span>

            {isOpen ? <ExpandLessIcon /> : <ExpandMoreIcon />}
          </React.Fragment>
        )}
      </div>

      <Collapse in={isOpen} timeout="auto" unmountOnExit>
        <ul className="py-2 px-4 space-y-2">{children}</ul>
      </Collapse>
    </li>
  )
}

export default CollapsibleMenu
