import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet'
import AssessmentIcon from '@mui/icons-material/Assessment'
import AssignmentIcon from '@mui/icons-material/Assignment'
import BarChartIcon from '@mui/icons-material/BarChart'
import DashboardIcon from '@mui/icons-material/Dashboard'
import GroupAddIcon from '@mui/icons-material/GroupAdd'
import HttpIcon from '@mui/icons-material/Http'
import LayersIcon from '@mui/icons-material/Layers'
import ListIcon from '@mui/icons-material/List'
import PeopleIcon from '@mui/icons-material/People'
import PersonIcon from '@mui/icons-material/Person'
import ListItemButton from '@mui/material/ListItemButton'
import ListItemText from '@mui/material/ListItemText'
import ListSubheader from '@mui/material/ListSubheader'
import Tooltip from '@mui/material/Tooltip'
import * as React from 'react'

import ActiveIcon from './ActiveIcon'
import ActiveLink from './ActiveLink'

function ListItem({ url, icon, name }: any) {
  return (
    <ActiveLink
      href={url}
      key={url}
      passHref
      activeClassName="md:text-blue-700"
    >
      <Tooltip title={name} followCursor enterDelay={500}>
        <ListItemButton>
          <ActiveIcon href={url} activeClassName="md:text-blue-700">
            {icon}
          </ActiveIcon>
          <ListItemText primary={name} />
        </ListItemButton>
      </Tooltip>
    </ActiveLink>
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
  },
  {
    name: 'Integrations',
    url: '/user/integrations',
    icon: <LayersIcon />,
  },
]

export const mainListItems = mainMenuList.map((item) => (
  <ListItem key={item.url} {...item} />
))

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
