import { Popover, Transition } from '@headlessui/react'
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import Link from 'next/link'
import { Fragment } from 'react'

const solutions = [
  {
    name: 'Pricing',
    description: 'Measure actions your users take',
    href: '/pricing',
    icon: IconOne,
  },
  {
    name: 'Contacts',
    description: 'Send us an email',
    href: '/contact',
    icon: IconTwo,
  },
  {
    name: 'Reports',
    description: 'Keep track of your growth',
    href: '/user/reports',
    icon: IconThree,
  },
]

export default function Example() {
  return (
    <div className="top-16 max-w-sm px-4">
      <Popover className="relative">
        {({ open }) => (
          <>
            <Popover.Button
              className={`
                ${open ? '' : 'text-opacity-90'}
                group inline-flex items-center rounded-md bg-indigo-500 px-3 py-2 text-base font-medium text-white hover:text-opacity-100 focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 dark:bg-indigo-700 dark:text-gray-200 dark:focus-visible:ring-gray-400`}
            >
              <span>Solutions</span>
              <ChevronDownIcon
                className={`${open ? '' : 'text-opacity-70'}
                  ml-2 h-5 w-5 text-indigo-300 transition duration-150 ease-in-out group-hover:text-opacity-80 dark:text-indigo-500`}
                aria-hidden="true"
              />
            </Popover.Button>
            <Transition
              as="div"
              // ...
            >
              <Popover.Panel className="absolute left-1/2 z-10 mt-3 w-screen max-w-sm -translate-x-1/2 transform overflow-hidden rounded-lg bg-white/95 shadow-lg ring-1 ring-black ring-opacity-5 backdrop-blur dark:bg-gray-800/95 px-4 sm:px-0 lg:max-w-3xl">
                <div>
                  <div className="relative grid gap-8 bg-white dark:bg-gray-800 p-7 lg:grid-cols-2">
                    {solutions.map((item) => (
                      <Link key={item.href} href={item.href}>
                        <div className="-m-3 cursor-pointer flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 dark:hover:bg-gray-700 focus:outline-none focus-visible:ring focus-visible:ring-indigo-500 dark:focus-visible:ring-gray-400">
                          <div className="flex h-10 w-10 shrink-0 items-center justify-center text-white sm:h-12 sm:w-12 dark:text-gray-300">
                            <item.icon aria-hidden="true" />
                          </div>
                          <div className="ml-4">
                            <p className="text-sm font-medium text-gray-900 dark:text-gray-200">{item.name}</p>
                            <p className="text-sm text-gray-500 dark:text-gray-400">{item.description}</p>
                          </div>
                        </div>
                      </Link>
                    ))}
                  </div>
                  <div className="bg-gray-50 dark:bg-gray-800 p-4">
                    <Link href="/faq">
                      <span className="flow-root cursor-pointer rounded-md px-2 py-2 transition duration-150 ease-in-out hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus-visible:ring focus-visible:ring-indigo-500 dark:focus-visible:ring-gray-400">
                        <span className="flex items-center">
                          <span className="text-sm font-medium text-gray-900 dark:text-gray-200">Documentation</span>
                        </span>
                        <span className="block text-sm text-gray-500 dark:text-gray-400">Start integrating products and tools</span>
                      </span>
                    </Link>
                  </div>
                </div>
              </Popover.Panel>
            </Transition>
          </>
        )}
      </Popover>
    </div>
  )
}

function IconOne() {
  return (
    <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect width="48" height="48" rx="8" fill="#FFEDD5" />
      <path d="M24 11L35.2583 17.5V30.5L24 37L12.7417 30.5V17.5L24 11Z" stroke="#FB923C" strokeWidth="2" />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M16.7417 19.8094V28.1906L24 32.3812L31.2584 28.1906V19.8094L24 15.6188L16.7417 19.8094Z"
        stroke="#FDBA74"
        strokeWidth="2"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M20.7417 22.1196V25.882L24 27.7632L27.2584 25.882V22.1196L24 20.2384L20.7417 22.1196Z"
        stroke="#FDBA74"
        strokeWidth="2"
      />
    </svg>
  )
}

function IconTwo() {
  return (
    <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect width="48" height="48" rx="8" fill="#FFEDD5" />
      <path
        d="M28.0413 20L23.9998 13L19.9585 20M32.0828 27.0001L36.1242 34H28.0415M19.9585 34H11.8755L15.9171 27"
        stroke="#FB923C"
        strokeWidth="2"
      />
      <path fillRule="evenodd" clipRule="evenodd" d="M18.804 30H29.1963L24.0001 21L18.804 30Z" stroke="#FDBA74" strokeWidth="2" />
    </svg>
  )
}

function IconThree() {
  return (
    <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect width="48" height="48" rx="8" fill="#FFEDD5" />
      <rect x="13" y="32" width="2" height="4" fill="#FDBA74" />
      <rect x="17" y="28" width="2" height="8" fill="#FDBA74" />
      <rect x="21" y="24" width="2" height="12" fill="#FDBA74" />
      <rect x="25" y="20" width="2" height="16" fill="#FDBA74" />
      <rect x="29" y="16" width="2" height="20" fill="#FB923C" />
      <rect x="33" y="12" width="2" height="24" fill="#FB923C" />
    </svg>
  )
}
