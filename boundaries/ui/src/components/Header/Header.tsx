'use client'

import { Bars3Icon, MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import clsx from 'clsx'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { Suspense, useEffect, useState } from 'react'

import { ThemeToggle } from '@/components/ThemeToggle'
import { navItemIsActive } from '@/lib/app-base-path'

import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'

import { LogoSquare } from './LogoSquare'
import { desktopNavItems, mobileNavItems } from './nav-items'
import NavbarMobileMenu from './NavbarMobileMenu'
import NavbarSearch, { NavbarSearchSkeleton } from './NavbarSearch'
import ProfileMenu from './ProfileMenu'

const SITE_LABEL = 'ShortLink'

interface HeaderProps {
  hasSession: boolean
  isSessionLoading?: boolean
  setOpen: () => void
}

function SidebarToggle({ onClick }: { onClick: () => void }) {
  return (
    <button
      type="button"
      onClick={onClick}
      aria-label="Expand navigation"
      className="app-focus-ring flex h-11 w-11 items-center justify-center rounded-md border border-neutral-200 text-black transition-colors dark:border-neutral-700 dark:text-white"
    >
      <Bars3Icon className="h-5 w-5" />
    </button>
  )
}

function useMainScrollFolded() {
  const [scrolled, setScrolled] = useState(false)
  useEffect(() => {
    const region = document.querySelector<HTMLElement>('[data-app-scroll-region]')
    const read = () => {
      const y = region ? region.scrollTop : window.scrollY
      setScrolled(y > 10)
    }
    read()
    const target = region ?? window
    target.addEventListener('scroll', read, { passive: true })
    return () => target.removeEventListener('scroll', read)
  }, [])
  return scrolled
}

export default function Header({ hasSession, setOpen }: HeaderProps) {
  const pathname = usePathname()
  const scrolled = useMainScrollFolded()
  const desktop = desktopNavItems(hasSession)
  const mobile = mobileNavItems(hasSession)

  return (
    <header
      className={clsx(
        'sticky top-0 z-40 border-b border-[var(--color-border)] bg-[var(--color-surface)]/95 backdrop-blur-md transition-[box-shadow,background-color] duration-200',
        scrolled &&
          'shadow-[0_10px_36px_-14px_rgb(15_23_42/0.22)] dark:shadow-[0_12px_40px_-16px_rgb(0_0_0/0.55)]',
      )}
    >
      <nav
        className={clsx(
          'relative mx-auto flex max-w-[1600px] items-center justify-between px-4 transition-[padding] duration-200 ease-out lg:px-6',
          scrolled ? 'py-2.5' : 'py-4',
        )}
      >
        <div className="block flex-none md:hidden">
          {hasSession ? (
            <SidebarToggle onClick={setOpen} />
          ) : (
            <Suspense
              fallback={
                <div
                  className="app-focus-ring flex h-11 w-11 items-center justify-center rounded-md border border-neutral-200 opacity-50 dark:border-neutral-700"
                  aria-hidden
                />
              }
            >
              <NavbarMobileMenu menu={mobile} showSearch={false} />
            </Suspense>
          )}
        </div>

        <div className="flex w-full min-w-0 items-center">
          <div className="flex w-full min-w-0 items-center md:w-1/3">
            {hasSession ? (
              <div className="mr-2 hidden shrink-0 md:block">
                <SidebarToggle onClick={setOpen} />
              </div>
            ) : null}
            <Link
              href="/"
              prefetch
              transitionTypes={resolveNavTransitionTypes(pathname, '/')}
              className="app-focus-ring mr-2 flex w-full min-w-0 items-center justify-center gap-2 rounded-md md:w-auto md:flex-none lg:mr-6"
            >
              <LogoSquare />
              <div className="ml-2 flex-none text-sm font-medium uppercase text-neutral-900 md:hidden lg:block dark:text-white">
                {SITE_LABEL}
              </div>
            </Link>
            {desktop.length > 0 ? (
              <ul className="hidden gap-6 text-sm md:flex md:items-center">
                {desktop.map((item) => (
                  <li key={item.path}>
                    <Link
                      href={item.path}
                      prefetch
                      transitionTypes={resolveNavTransitionTypes(pathname, item.path)}
                      aria-current={navItemIsActive(pathname, item.path) ? 'page' : undefined}
                      className={clsx(
                        'app-focus-ring rounded-md px-1.5 py-1 text-sm underline-offset-4 transition-colors',
                        navItemIsActive(pathname, item.path)
                          ? 'font-semibold text-[var(--color-foreground)] dark:text-zinc-50'
                          : 'text-neutral-500 hover:text-neutral-900 hover:underline dark:text-neutral-400 dark:hover:text-zinc-200',
                      )}
                    >
                      {item.title}
                    </Link>
                  </li>
                ))}
              </ul>
            ) : null}
          </div>

          <div className="hidden min-w-0 justify-center md:flex md:w-1/3">
            {hasSession ? (
              <Suspense fallback={<NavbarSearchSkeleton />}>
                <NavbarSearch />
              </Suspense>
            ) : null}
          </div>

          <div className="flex w-auto min-w-0 flex-1 items-center justify-end gap-2 md:w-1/3 md:gap-4">
            {hasSession ? (
              <Link
                href="/links/search"
                prefetch
                transitionTypes={resolveNavTransitionTypes(pathname, '/links/search')}
                aria-label="Search links"
                className="app-focus-ring flex h-11 w-11 items-center justify-center rounded-md border border-neutral-200 text-neutral-700 transition-colors hover:bg-neutral-50 md:hidden dark:border-neutral-700 dark:text-neutral-200 dark:hover:bg-neutral-800"
              >
                <MagnifyingGlassIcon className="h-5 w-5" />
              </Link>
            ) : null}
            <ThemeToggle />
            <ProfileMenu />
          </div>
        </div>
      </nav>
    </header>
  )
}
