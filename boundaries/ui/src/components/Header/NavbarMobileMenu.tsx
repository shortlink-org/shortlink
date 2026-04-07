'use client'

import { Dialog, DialogBackdrop, DialogPanel } from '@headlessui/react'
import { Bars3Icon, XMarkIcon } from '@heroicons/react/24/outline'
import clsx from 'clsx'
import Link from 'next/link'
import { usePathname, useSearchParams } from 'next/navigation'
import { Suspense, useEffect, useState } from 'react'

import { navItemIsActive } from '@/lib/app-base-path'
import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'

import type { NavMenuItem } from './nav-items'
import NavbarSearch, { NavbarSearchSkeleton } from './NavbarSearch'

type NavbarMobileMenuProps = {
  menu: NavMenuItem[]
  showSearch: boolean
}

export default function NavbarMobileMenu({ menu, showSearch }: NavbarMobileMenuProps) {
  const pathname = usePathname()
  const searchParams = useSearchParams()
  const [isOpen, setIsOpen] = useState(false)

  useEffect(() => {
    const onResize = () => {
      if (window.innerWidth >= 768) setIsOpen(false)
    }
    window.addEventListener('resize', onResize)
    return () => window.removeEventListener('resize', onResize)
  }, [])

  useEffect(() => {
    setIsOpen(false)
  }, [pathname, searchParams])

  return (
    <>
      <button
        type="button"
        onClick={() => setIsOpen(true)}
        aria-label="Open menu"
        className="app-focus-ring flex h-11 w-11 items-center justify-center rounded-md border border-neutral-200 text-black transition-colors md:hidden dark:border-neutral-700 dark:text-white"
      >
        <Bars3Icon className="h-5 w-5" />
      </button>

      <Dialog open={isOpen} onClose={() => setIsOpen(false)} transition className="relative z-50 md:hidden">
        <DialogBackdrop
          transition
          className="fixed inset-0 bg-black/30 transition duration-300 ease-in-out data-closed:opacity-0"
        />
        <DialogPanel
          transition
          className="fixed inset-y-0 left-0 flex h-full w-full max-w-sm flex-col bg-white pb-6 transition duration-300 ease-in-out data-closed:-translate-x-full dark:bg-[var(--color-background)]"
        >
          <div className="p-4">
            <button
              type="button"
              className="app-focus-ring mb-4 flex h-11 w-11 items-center justify-center rounded-md border border-neutral-200 text-black transition-colors dark:border-neutral-700 dark:text-white"
              onClick={() => setIsOpen(false)}
              aria-label="Close menu"
            >
              <XMarkIcon className="h-6 w-6" />
            </button>

            {showSearch ? (
              <div className="mb-4 w-full">
                <Suspense fallback={<NavbarSearchSkeleton />}>
                  <NavbarSearch />
                </Suspense>
              </div>
            ) : null}

            {menu.length > 0 ? (
              <ul className="flex w-full flex-col gap-1">
                {menu.map((item) => (
                  <li key={item.path}>
                    <Link
                      href={item.path}
                      prefetch
                      transitionTypes={resolveNavTransitionTypes(pathname, item.path)}
                      onClick={() => setIsOpen(false)}
                      aria-current={navItemIsActive(pathname, item.path) ? 'page' : undefined}
                      className={clsx(
                        'app-focus-ring block rounded-md py-2 text-xl transition-colors',
                        navItemIsActive(pathname, item.path)
                          ? 'font-semibold text-neutral-900 dark:text-white'
                          : 'text-neutral-600 hover:text-neutral-900 dark:text-zinc-300 dark:hover:text-white',
                      )}
                    >
                      {item.title}
                    </Link>
                  </li>
                ))}
              </ul>
            ) : null}
          </div>
        </DialogPanel>
      </Dialog>
    </>
  )
}
