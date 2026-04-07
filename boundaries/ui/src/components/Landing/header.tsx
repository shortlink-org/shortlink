'use client'

import { Button } from '@shortlink-org/ui-kit'
import Image from 'next/image'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import Balancer from 'react-wrap-balancer'

import { resolveNavTransitionTypes } from '@/lib/nav-link-transition'

import styles from './header.module.css'

export default function Header() {
  const pathname = usePathname()

  return (
    <section
      aria-labelledby="landing-hero-title"
      className={`${styles.root} relative w-full min-w-0 overflow-hidden rounded-[1.75rem] border border-[var(--color-border)]/90 bg-gradient-to-br from-white via-zinc-50/90 to-violet-50/40 shadow-[0_20px_50px_-38px_rgb(15_23_42/0.28)] dark:border-white/[0.09] dark:from-zinc-900/95 dark:via-zinc-950 dark:to-indigo-950/50 dark:shadow-[0_24px_60px_-40px_rgb(0_0_0/0.65)]`}
    >
      <div className={`${styles.grid} w-full min-w-0`}>
        <div className={`${styles.content} relative z-10`}>
          <div className={styles.body}>
            <p className="text-sm font-semibold uppercase tracking-[0.22em] text-indigo-600 dark:text-indigo-300">
              Link intelligence
            </p>
            <h1
              id="landing-hero-title"
              className={`${styles.title} mt-4 tracking-tight font-extrabold sm:mt-5`}
            >
              <Balancer>
                <span className="block text-gray-900 xl:inline dark:text-zinc-50">Shorten your links</span>{' '}
                <span className="block text-indigo-600 xl:inline dark:text-indigo-300">in one click</span>
              </Balancer>
            </h1>
            <p className={`${styles.lead} mt-4 text-gray-600 sm:mt-5 dark:text-zinc-300`}>
              Get full control over your links in one place. <br />
              Easily create, manage, and track your links. <br />
              Get started today!
            </p>
            <div className={`mt-6 ${styles.actions} sm:mt-7`}>
              <Button
                as={Link}
                asProps={{
                  href: '/auth/login',
                  transitionTypes: resolveNavTransitionTypes(pathname, '/auth/login'),
                }}
                size="lg"
                className="group relative justify-center shadow-lg ring-1 ring-violet-500/25 sm:min-w-[12.5rem] sm:px-8 md:py-4 md:text-lg lg:min-w-[14rem] lg:px-10 dark:ring-indigo-400/20"
              >
                <span className="flex items-center gap-2">
                  Get started
                  <svg
                    className="h-5 w-5 transition-transform motion-reduce:transform-none group-hover:translate-x-0.5 group-hover:sm:translate-x-1"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    aria-hidden
                  >
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
                  </svg>
                </span>
              </Button>
              <Button
                as={Link}
                asProps={{
                  href: '#features',
                  transitionTypes: resolveNavTransitionTypes(pathname, '#features'),
                }}
                variant="secondary"
                size="lg"
                className="group justify-center max-md:min-h-[2.75rem] max-md:py-2.5 max-md:text-sm sm:px-8 md:py-3.5 md:text-base lg:px-10"
              >
                <span className="flex items-center gap-2 font-medium text-[var(--color-muted-foreground)] group-hover:text-[var(--color-foreground)]">
                  Explore features
                  <svg
                    className="h-4 w-4 shrink-0 transition-transform motion-reduce:transform-none sm:h-5 sm:w-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    aria-hidden
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                    />
                  </svg>
                </span>
              </Button>
            </div>
          </div>
        </div>
        <div className={styles.media}>
          <div className={styles.mediaCrop}>
            <Image
              className={styles.image}
              src="https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80"
              alt="Team collaborating with laptops in a bright office"
              width={1425}
              height={950}
              priority
              sizes="(min-width: 1024px) 42vw, 100vw"
            />
          </div>
        </div>
      </div>
    </section>
  )
}
