import React from 'react'
import { Menu, Transition } from '@headlessui/react'
import Link from 'next/link'

const profile = [
  {
    name: 'Your Profile',
    link: '#',
  },
  {
    name: 'Sign out',
    link: `http://127.0.0.1:4433/self-service/browser/flows/logout`,
  },
]

function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export default function Profile() {
  // @ts-ignore
  return (
    <Menu as="div" className="ml-3 relative">
      {({ open }) => (
        <React.Fragment>
          <div>
            <Menu.Button className="max-w-xs bg-gray-800 rounded-full flex items-center text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white">
              <span className="sr-only">Open user menu</span>
              <img
                className="h-8 w-8 rounded-full"
                src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                alt=""
              />
            </Menu.Button>
          </div>
          <Transition
            show={open}
            as={React.Fragment}
            enter="transition ease-out duration-100"
            enterFrom="transform opacity-0 scale-95"
            enterTo="transform opacity-100 scale-100"
            leave="transition ease-in duration-75"
            leaveFrom="transform opacity-100 scale-100"
            leaveTo="transform opacity-0 scale-95"
          >
            <Menu.Items
              static
              className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
            >
              {profile.map((item) => (
                <Menu.Item key={item.name}>
                  {({ active }) => (
                    <span
                      className={classNames(
                        active ? 'bg-gray-100' : '',
                        'block px-4 py-2 text-sm text-gray-700',
                      )}
                    >
                      <Link
                        href={item.link}
                      >
                        <p>{item.name}</p>
                      </Link>
                    </span>
                  )}
                </Menu.Item>
              ))}
            </Menu.Items>
          </Transition>
        </React.Fragment>
      )}
    </Menu>
  )
}
