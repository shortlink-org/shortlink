'use client'

import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react'
import { ArrowRightOnRectangleIcon, ChartBarIcon, UserIcon } from '@heroicons/react/24/outline'
import { AxiosError } from 'axios'
import clsx from 'clsx'
import Link from 'next/link'
import { useCallback, useEffect, useState } from 'react'

import { useSession } from '@/contexts/SessionContext'
import { usePathname } from 'next/navigation'

import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'
import ory from '@/pkg/sdk'

export default function ProfileMenu() {
  const pathname = usePathname()
  const [logoutToken, setLogoutToken] = useState('')
  const { session, hasSession, isLoading } = useSession()

  const traits = (session?.identity?.traits as Record<string, unknown> | undefined) ?? {}
  const nameTraits = traits.name as Record<string, string> | undefined
  const firstName = nameTraits?.first ?? ''
  const lastName = nameTraits?.last ?? ''
  const email = (traits.email as string) ?? ''
  const displayName = `${firstName} ${lastName}`.trim() || email || 'User'
  const initials =
    (firstName?.[0] ?? '') + (lastName?.[0] ?? '') || (email?.[0] ?? 'U').toUpperCase() || 'U'

  useEffect(() => {
    if (!hasSession) return

    ory
      .createBrowserLogoutFlow()
      .then(({ data }) => setLogoutToken(data.logout_token))
      .catch((err: AxiosError) => {
        if (err.response?.status === 401) return
        return Promise.reject(err)
      })
  }, [hasSession])

  const handleLogout = useCallback(async () => {
    try {
      const token = logoutToken || (await ory.createBrowserLogoutFlow()).data.logout_token
      await ory.updateLogoutFlow({ token })
      window.location.assign('/auth/login')
    } catch (err) {
      console.error('Logout failed', err)
    }
  }, [logoutToken])

  if (isLoading) {
    return <div className="h-8 w-8 shrink-0 animate-pulse rounded-full bg-neutral-200 dark:bg-neutral-700" />
  }

  if (!hasSession) {
    return (
      <Link
        href="/auth/login"
        prefetch
        transitionTypes={resolveNavTransitionTypes(pathname, '/auth/login')}
        className="app-focus-ring rounded-md px-4 py-2 text-sm font-medium text-neutral-900 transition-opacity hover:opacity-80 dark:text-white"
      >
        Log in
      </Link>
    )
  }

  return (
    <Menu>
      <MenuButton className="app-focus-ring flex items-center gap-2 rounded-full p-1.5 transition-all hover:bg-neutral-100 dark:hover:bg-neutral-800">
        <span className="sr-only">Open user menu</span>
        <span className="flex h-8 w-8 items-center justify-center rounded-full bg-teal-600 text-xs font-semibold text-white dark:bg-teal-500">
          {initials}
        </span>
      </MenuButton>

      <MenuItems
        transition
        anchor="bottom end"
        className="z-50 mt-2 w-56 origin-top-right rounded-xl border border-neutral-200 bg-white shadow-lg outline-none transition duration-100 ease-out data-closed:scale-95 data-closed:opacity-0 dark:border-neutral-700 dark:bg-neutral-800"
      >
        <div className="border-b border-neutral-100 px-4 py-3 dark:border-neutral-700">
          <p className="truncate text-sm font-semibold text-neutral-900 dark:text-white">{displayName}</p>
          {email ? (
            <p className="truncate text-xs text-neutral-500 dark:text-neutral-400">{email}</p>
          ) : null}
        </div>
        <div className="p-1">
          <MenuItem>
            {({ focus }) => (
              <Link
                href="/profile"
                transitionTypes={resolveNavTransitionTypes(pathname, '/profile')}
                className={clsx(
                  focus ? 'bg-neutral-50 dark:bg-neutral-700' : '',
                  'flex items-center gap-3 rounded-lg px-3 py-2 text-sm text-neutral-700 dark:text-neutral-200',
                )}
              >
                <UserIcon className="h-5 w-5 text-neutral-500 dark:text-neutral-400" />
                Your profile
              </Link>
            )}
          </MenuItem>
          <MenuItem>
            {({ focus }) => (
              <Link
                href="/user/reports"
                transitionTypes={resolveNavTransitionTypes(pathname, '/user/reports')}
                className={clsx(
                  focus ? 'bg-neutral-50 dark:bg-neutral-700' : '',
                  'flex items-center gap-3 rounded-lg px-3 py-2 text-sm text-neutral-700 dark:text-neutral-200',
                )}
              >
                <ChartBarIcon className="h-5 w-5 text-neutral-500 dark:text-neutral-400" />
                Reports
              </Link>
            )}
          </MenuItem>
          <MenuItem>
            {({ focus }) => (
              <button
                type="button"
                onClick={handleLogout}
                className={clsx(
                  focus ? 'bg-neutral-50 dark:bg-neutral-700' : '',
                  'flex w-full items-center gap-3 rounded-lg px-3 py-2 text-left text-sm text-neutral-700 dark:text-neutral-200',
                )}
              >
                <ArrowRightOnRectangleIcon className="h-5 w-5 text-neutral-500 dark:text-neutral-400" />
                Sign out
              </button>
            )}
          </MenuItem>
        </div>
      </MenuItems>
    </Menu>
  )
}
