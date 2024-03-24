import * as React from 'react'
import List from '@mui/material/List'
import ListItemButton from '@mui/material/ListItemButton'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemText from '@mui/material/ListItemText'
import Collapse from '@mui/material/Collapse'
import ListSubheader from '@mui/material/ListSubheader'
// eslint-disable-next-line import/order
import Tooltip from '@mui/material/Tooltip'

// Importing icons
import HttpIcon from '@mui/icons-material/Http'
import DashboardIcon from '@mui/icons-material/Dashboard'
import ListIcon from '@mui/icons-material/List'
import BarChartIcon from '@mui/icons-material/BarChart'
import PersonIcon from '@mui/icons-material/Person'
import SecurityIcon from '@mui/icons-material/Security'
import LayersIcon from '@mui/icons-material/Layers'
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet'
import AssessmentIcon from '@mui/icons-material/Assessment'
import AssignmentIcon from '@mui/icons-material/Assignment'
import PeopleIcon from '@mui/icons-material/People'
import GroupAddIcon from '@mui/icons-material/GroupAdd'
import ExpandLessIcon from '@mui/icons-material/ExpandLess'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'

import ActiveIcon from './ActiveIcon'
import ActiveLink from './ActiveLink'

// Define types for menu item and nested menu item -----------------------------
interface NestedMenuItem {
  name: string
  url: string
  icon: JSX.Element
}

interface MenuItemProps {
  name: string
  url: string
  icon: JSX.Element
  nestedList?: NestedMenuItem[]
}

// ListItem component ----------------------------------------------------------
const ListItem: React.FC<MenuItemProps> = ({ url, icon, name, nestedList }) => {
  const [open, setOpen] = React.useState(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <>
      <Tooltip title={name} followCursor enterDelay={500}>
        <ActiveLink href={url} key={url} passHref activeClassName="md:text-blue-700">
          <ListItemButton onClick={handleClick}>
            <ActiveIcon href={url} activeClassName="md:text-blue-700">
              {icon}
            </ActiveIcon>
            <ListItemText primary={name} />
            {nestedList && (open ? <ExpandLessIcon /> : <ExpandMoreIcon />)}
          </ListItemButton>
        </ActiveLink>
      </Tooltip>
      {nestedList && (
        <Collapse in={open} timeout="auto" unmountOnExit>
          <List component="div" disablePadding>
            {nestedList.map((item: NestedMenuItem) => (
              <ActiveLink href={item.url} key={item.url} passHref activeClassName="md:text-blue-700">
                <ListItemButton sx={{ pl: 4 }}>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText primary={item.name} />
                </ListItemButton>
              </ActiveLink>
            ))}
          </List>
        </Collapse>
      )}
    </>
  )
}

const mainMenuList = [
  {
    name: 'Add URL',
    url: '/user/addUrl',
    icon: <HttpIcon />,
  },
  {
    name: 'Dashboard',
    url: '/user/dashboard',
    icon: <DashboardIcon />,
  },
  {
    name: 'Links',
    url: '/user/links',
    icon: <ListIcon />,
  },
  {
    name: 'Reports',
    url: '/user/reports',
    icon: <BarChartIcon />,
  },
  {
    name: 'Profile',
    url: '/user/profile',
    icon: <PersonIcon />,
    nestedList: [
      {
        name: 'Security',
        url: '/user/security',
        icon: <SecurityIcon />,
      },
    ],
  },
  {
    name: 'Integrations',
    url: '/user/integrations',
    icon: <LayersIcon />,
  },
]

export const mainListItems: React.FC = () => (
  <List>
    {mainMenuList.map((item) => (
      <ListItem key={item.url} {...item} />
    ))}
  </List>
)

const otherMenuList = [
  {
    name: 'Billing',
    url: '/user/billing',
    icon: <AccountBalanceWalletIcon />,
  },
  {
    name: 'Audit',
    url: '/user/audit',
    icon: <AssessmentIcon />,
  },
  {
    name: 'About As',
    url: '/about',
    icon: <AssignmentIcon />,
  },
]

export const secondaryListItems = [
  <ListSubheader key="other options" inset>
    Other options
  </ListSubheader>,
  otherMenuList.map((item) => <ListItem key={item.url} {...item} />),
]

const adminMenuList = [
  {
    name: 'Groups',
    url: '/admin/groups',
    icon: <PeopleIcon />,
  },
  {
    name: 'Users',
    url: '/admin/users',
    icon: <GroupAddIcon />,
  },
  {
    name: 'Links',
    url: '/admin/links',
    icon: <ListIcon />,
  },
]

export const adminListItems = [
  <ListSubheader key="admin options" inset>
    Admin options
  </ListSubheader>,
  adminMenuList.map((item) => <ListItem key={item.url} {...item} />),
]
