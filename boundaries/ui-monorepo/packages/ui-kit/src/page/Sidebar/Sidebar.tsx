import * as React from 'react'

// Importing icons
import AddLinkIcon from '@mui/icons-material/AddLink'
import SettingsIcon from '@mui/icons-material/Settings'
import ReceiptIcon from '@mui/icons-material/Receipt'
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
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart'
import AdminPanelSettingsIcon from '@mui/icons-material/AdminPanelSettings'

// Components
import ActiveLink from './ActiveLink'
import CollapsibleMenu from './CollapsibleMenu'
import Footer from './Footer'

type AppProps = {
  mode: 'full' | 'mini'
}

const iconClassName =
  'w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
const linkClassName =
  'flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group'

function getLink(url: string, icon: JSX.Element, name: string) {
  return (
    <li>
      <ActiveLink
        href={url}
        key={url}
        passHref
        activeClassName="md:text-blue-700"
      >
        <div className={linkClassName}>
          {React.cloneElement(icon, { className: iconClassName })}
          <span className="ms-3">{name}</span>
        </div>
      </ActiveLink>
    </li>
  )
}

export function Sidebar({ mode }: AppProps) {
  return (
    <aside
      id="default-sidebar"
      className="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
      aria-label="Sidebar"
    >
      <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
        <ul className="space-y-2 font-medium">
          {getLink('/user/addUrl', <AddLinkIcon />, 'Add URL')}

          <CollapsibleMenu icon={HttpIcon} title="Links">
            {getLink('/user/dashboard', <DashboardIcon />, 'Dashboard')}
            {getLink('/user/links', <ListIcon />, 'Links')}
            {getLink('/user/reports', <BarChartIcon />, 'Reports')}
          </CollapsibleMenu>

          <CollapsibleMenu icon={SettingsIcon} title="Setting">
            {getLink('/user/profile', <PersonIcon />, 'Profile')}
            {getLink('/user/security', <SecurityIcon />, 'Security')}
            {getLink('/user/integrations', <LayersIcon />, 'Integrations')}
            {getLink('/user/audit', <AssessmentIcon />, 'Audit')}
          </CollapsibleMenu>

          <CollapsibleMenu icon={ShoppingCartIcon} title="E-commerce">
            {getLink('/user/billing', <AccountBalanceWalletIcon />, 'Billing')}
            {getLink('/user/invoice', <ReceiptIcon />, 'Invoice')}
          </CollapsibleMenu>

          <CollapsibleMenu icon={AdminPanelSettingsIcon} title="Admin">
            {getLink('/admin/links', <ListIcon />, 'Links')}
            {getLink('/admin/users', <GroupAddIcon />, 'Users')}
            {getLink('/admin/groups', <PeopleIcon />, 'Groups')}
          </CollapsibleMenu>

          {getLink('/about', <AssignmentIcon />, 'About As')}
        </ul>

        <Footer />
      </div>
    </aside>
  )
}

export default Sidebar
