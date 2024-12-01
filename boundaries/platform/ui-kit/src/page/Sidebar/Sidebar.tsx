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
import CollapsibleMenu from './CollapsibleMenu'
import Footer from './Footer'
import getItem from './Item'

type AppProps = {
  mode: 'full' | 'mini'
}

export function Sidebar({ mode }: AppProps) {
  return (
    <aside
      className="w-full h-full bg-gray-50 dark:bg-gray-800 flex justify-between flex-col min-h-0"
      aria-label="Sidebar"
    >
      <ul className="space-y-2 font-medium flex-grow w-full h-full px-2 py-4 overflow-y-auto">
        {getItem({
          mode,
          url: '/user/addUrl',
          icon: <AddLinkIcon />,
          name: 'Add URL',
        })}

        <CollapsibleMenu icon={HttpIcon} title="Links" mode={mode}>
          {getItem({
            mode,
            url: '/user/dashboard',
            icon: <DashboardIcon />,
            name: 'Dashboard',
          })}
          {getItem({
            mode,
            url: '/user/links',
            icon: <ListIcon />,
            name: 'Links',
          })}
          {getItem({
            mode,
            url: '/user/reports',
            icon: <BarChartIcon />,
            name: 'Reports',
          })}
        </CollapsibleMenu>

        <CollapsibleMenu icon={SettingsIcon} title="Setting" mode={mode}>
          {getItem({
            mode,
            url: '/user/profile',
            icon: <PersonIcon />,
            name: 'Profile',
          })}
          {getItem({
            mode,
            url: '/user/security',
            icon: <SecurityIcon />,
            name: 'Security',
          })}
          {getItem({
            mode,
            url: '/user/integrations',
            icon: <LayersIcon />,
            name: 'Integrations',
          })}
          {getItem({
            mode,
            url: '/user/audit',
            icon: <AssessmentIcon />,
            name: 'Audit',
          })}
        </CollapsibleMenu>

        <CollapsibleMenu icon={ShoppingCartIcon} title="E-commerce" mode={mode}>
          {getItem({
            mode,
            url: '/user/billing',
            icon: <AccountBalanceWalletIcon />,
            name: 'Billing',
          })}
          {getItem({
            mode,
            url: '/user/invoice',
            icon: <ReceiptIcon />,
            name: 'Invoice',
          })}
        </CollapsibleMenu>

        <CollapsibleMenu
          icon={AdminPanelSettingsIcon}
          title="Admin"
          mode={mode}
        >
          {getItem({
            mode,
            url: '/admin/links',
            icon: <ListIcon />,
            name: 'Links',
          })}
          {getItem({
            mode,
            url: '/admin/users',
            icon: <GroupAddIcon />,
            name: 'Users',
          })}
          {getItem({
            mode,
            url: '/admin/groups',
            icon: <PeopleIcon />,
            name: 'Groups',
          })}
        </CollapsibleMenu>

        {getItem({
          mode,
          url: '/about',
          icon: <AssignmentIcon />,
          name: 'About As',
        })}
      </ul>

      <Footer mode={mode} />
    </aside>
  )
}

export default Sidebar
