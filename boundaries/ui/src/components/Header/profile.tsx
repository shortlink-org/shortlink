import { Menu, Transition } from '@headlessui/react'
import { AxiosError } from 'axios'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState, useEffect, Fragment } from 'react'

import { useSession } from '@/contexts/SessionContext'
import ory from '@/pkg/sdk'

// @ts-ignore
function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export default function Profile() {
  const [logoutToken, setLogoutToken] = useState<string>('')
  const router = useRouter()
  const { session } = useSession()

  const traits: any = session?.identity?.traits ?? {}
  const firstName = traits?.name?.first ?? ''
  const lastName = traits?.name?.last ?? ''
  const email = traits?.email ?? ''
  const displayName = `${firstName} ${lastName}`.trim() || email || 'User'
  const initials =
    (firstName?.[0] ?? '') + (lastName?.[0] ?? '') || (email?.[0] ?? '').toUpperCase() || 'U'

  useEffect(() => {
    ory
      .createBrowserLogoutFlow()
      .then(({ data }) => {
        setLogoutToken(data.logout_token)
      })
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 401:
            // do nothing, the user is not logged in
            return
          default:
          // Otherwise, we nothitng - the error will be handled by the Flow component
        }

        // Something else happened!
        return Promise.reject(err)
      })
  }, [])

  const profile = [
    {
      name: 'Your Profile',
      link: '/profile',
      icon: '👤',
    },
    {
      name: 'Sign out',
      link: `#`,
      icon: '🚪',
      onClick: () => {
        ory
          .updateLogoutFlow({ token: logoutToken })
          .then(() => router.push('/auth/login'))
          .then(() => window.location.reload())
      },
    },
  ]

  // @ts-ignore
  return (
    <Menu as="div" className="relative">
      {({ open }) => (
        <>
          <div>
            <Menu.Button className="group relative flex items-center gap-2 bg-white/10 hover:bg-white/20 rounded-full p-1.5 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-white/50 focus:ring-offset-2 focus:ring-offset-indigo-600">
              <span className="sr-only">Open user menu</span>
              <div
                className="h-8 w-8 rounded-full ring-2 ring-white/20 group-hover:ring-white/40 transition-all duration-200 bg-white/20 text-white flex items-center justify-center text-xs font-semibold"
                aria-hidden="true"
              >
                {initials}
              </div>
              <svg
                className={`w-4 h-4 text-white/70 transition-transform duration-200 ${open ? 'rotate-180' : ''}`}
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
              </svg>
            </Menu.Button>
          </div>

          <Transition
            show={open}
            as="div"
            enter="transition ease-out duration-200"
            enterFrom="transform opacity-0 scale-95 translate-y-2"
            enterTo="transform opacity-100 scale-100 translate-y-0"
            leave="transition ease-in duration-150"
            leaveFrom="transform opacity-100 scale-100 translate-y-0"
            leaveTo="transform opacity-0 scale-95 translate-y-2"
          >
            <Menu.Items
              static
              className="absolute right-0 mt-3 w-56 origin-top-right rounded-xl bg-white dark:bg-gray-800 shadow-xl ring-1 ring-black/5 dark:ring-white/10 focus:outline-none overflow-hidden"
            >
              <div className="px-4 py-3 border-b border-gray-100 dark:border-gray-700">
                <div className="text-sm font-semibold text-gray-900 dark:text-gray-100 truncate">{displayName}</div>
                {email && <div className="text-xs text-gray-500 dark:text-gray-400 truncate">{email}</div>}
              </div>
              <div className="px-1 py-2">
                {profile.map((item) => (
                  <Menu.Item key={item.name}>
                    {({ active }) => (
                      <div
                        className={classNames(
                          active ? 'bg-gray-50 dark:bg-gray-700' : '',
                          'flex items-center gap-3 px-3 py-2.5 rounded-lg mx-1 cursor-pointer transition-all duration-150',
                        )}
                        onClick={item.onClick}
                      >
                        <span className="text-lg">{item.icon}</span>
                        <span className="text-sm font-medium text-gray-700 dark:text-gray-200">
                          {item.onClick ? (
                            item.name
                          ) : (
                            <Link href={item.link} passHref className="block w-full">
                              {item.name}
                            </Link>
                          )}
                        </span>
                      </div>
                    )}
                  </Menu.Item>
                ))}
              </div>
            </Menu.Items>
          </Transition>
        </>
      )}
    </Menu>
  )
}
