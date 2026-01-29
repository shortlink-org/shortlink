/**
 * Layout - Updated with NavigationProvider and SessionContext
 * 
 * Changes:
 * - ✅ Added NavigationProvider for global progress bar
 * - ✅ Content fades slightly during navigation
 * - ✅ Progress bar shows at top during page transitions
 * - ✅ Removed duplicate session check (uses SessionContext)
 */

import CssBaseline from '@mui/material/CssBaseline'
import AddLinkIcon from '@mui/icons-material/AddLink'
import AdminPanelSettingsIcon from '@mui/icons-material/AdminPanelSettings'
import GroupsIcon from '@mui/icons-material/Groups'
import ListIcon from '@mui/icons-material/List'
import PersonIcon from '@mui/icons-material/Person'
import GroupIcon from '@mui/icons-material/Group'
import TravelExploreIcon from '@mui/icons-material/TravelExplore'
import React, { useState } from 'react'
// @ts-ignore
import { ScrollToTopButton, Sidebar } from '@shortlink-org/ui-kit'

// import PushNotificationLayout from '@/components/PushNotificationLayout'
import Footer from '@/components/Footer'
import Header from '@/components/Header'
import { NavigationProvider } from '@/components/Navigation'
import { useSession } from '@/contexts/SessionContext'

// @ts-ignore
export function Layout({ children }: { children: React.ReactNode }) {
  const [open, setOpen] = useState(false)
  const { hasSession, isLoading: isSessionLoading } = useSession()

  const sidebarSections = [
    {
      type: 'simple',
      items: [
        {
          url: '/add-link',
          icon: <AddLinkIcon />,
          name: 'Add URL',
        },
        {
          url: '/links',
          icon: <ListIcon />,
          name: 'Links',
        },
        {
          url: '/profile',
          icon: <PersonIcon />,
          name: 'Profile',
        },
      ],
    },
    {
      type: 'collapsible',
      icon: AdminPanelSettingsIcon,
      title: 'Admin',
      items: [
        {
          url: '/admin/links',
          icon: <ListIcon />,
          name: 'Links',
        },
        {
          url: '/admin/sitemap',
          icon: <TravelExploreIcon />,
          name: 'Sitemaps',
        },
        {
          url: '/admin/users',
          icon: <GroupIcon />,
          name: 'Users',
        },
        {
          url: '/admin/groups',
          icon: <GroupsIcon />,
          name: 'Groups',
        },
      ],
    },
  ]
  

  return (
    <NavigationProvider>
      <CssBaseline />

      <div className="grid grid-rows-[auto_1fr] h-screen overflow-hidden">
        <Header hasSession={hasSession} isSessionLoading={isSessionLoading} setOpen={() => setOpen(!open)} />

        <main className={'grid grid-cols-[auto_1fr] min-h-0'}>
          <div className={'h-full overflow-auto'}>
            {hasSession && <Sidebar mode={open ? 'full' : 'mini'} sections={sidebarSections} />}
          </div>

          <div className="overflow-auto h-full">
            {children}

            <Footer />

            <ScrollToTopButton />
          </div>
        </main>
      </div>
    </NavigationProvider>
  )
}
