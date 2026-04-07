export type NavMenuItem = {
  title: string
  path: string
}

const guestNav: NavMenuItem[] = [
  { title: 'Home', path: '/' },
  { title: 'Contact', path: '/contact' },
  { title: 'FAQ', path: '/faq' },
]

const authedDesktopNav: NavMenuItem[] = [
  { title: 'Home', path: '/' },
  { title: 'Links', path: '/links' },
  { title: 'Add link', path: '/add-link' },
]

export function desktopNavItems(hasSession: boolean): NavMenuItem[] {
  return hasSession ? authedDesktopNav : guestNav
}

export function mobileNavItems(hasSession: boolean): NavMenuItem[] {
  if (!hasSession) return guestNav
  return [
    ...authedDesktopNav,
    { title: 'Reports', path: '/user/reports' },
    { title: 'Contact', path: '/contact' },
    { title: 'FAQ', path: '/faq' },
  ]
}
