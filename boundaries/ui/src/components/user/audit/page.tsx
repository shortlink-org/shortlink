'use client'

import Link from 'next/link'

import Ready from '@/components/Landing/Ready'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

const auditEvents = [
  {
    id: 'role-change',
    title: 'Workspace role updated',
    time: 'Mar 14, 2026',
    description: 'An administrator changed workspace permissions for the Growth Team group.',
    href: '/profile',
  },
  {
    id: 'login-challenge',
    title: 'New sign-in challenge created',
    time: 'Mar 12, 2026',
    description: 'A new browser session triggered an identity verification challenge before access was granted.',
    href: '/security',
  },
  {
    id: 'api-access',
    title: 'Integration token viewed',
    time: 'Mar 10, 2026',
    description: 'A privileged workspace member accessed integration settings and reviewed API credentials.',
    href: '/integrations',
  },
  {
    id: 'bulk-update',
    title: 'Bulk link metadata edited',
    time: 'Mar 08, 2026',
    description: 'A large batch edit changed descriptions and campaign labels across multiple links.',
    href: '/links',
  },
]

function Page() {
  return (
    <>
      {/*<NextSeo title="Audit" description="Audit your account" />*/}

      <Header title="Audit" description="Review account activity and maintain an auditable trail of critical events." />

      <PageSection className="space-y-6 pb-10">
        <div className="rounded-2xl bg-white px-4 py-4 md:px-14 lg:px-8 lg:py-10 dark:bg-gray-800">
          <div className="flex flex-col">
            <p className="text-gray-800 dark:text-gray-300">
              Keep your business safe with shortlink. We add an audit trail so you can see what's happening in your account in real-time, and
              review past activity whenever you need to.
            </p>
          </div>
        </div>

        <Ready />

        <ul className="rounded-2xl bg-coolGray-100 p-4 text-coolGray-800 dark:bg-gray-900/40 dark:text-gray-100 lg:p-8">
          {auditEvents.map(event => (
            <li key={event.id}>
              <article>
                <Link
                  href={event.href}
                  className="grid overflow-hidden rounded-xl p-4 transition-colors hover:bg-white dark:hover:bg-gray-800 md:grid-cols-5 lg:p-6 xl:grid-cols-12"
                >
                  <h3 className="mb-1 ml-8 font-semibold md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                    {event.title}
                  </h3>
                  <time dateTime={event.time} className="row-start-1 mb-1 text-coolGray-600 md:col-start-1 xl:col-span-2">
                    {event.time}
                  </time>
                  <p className="ml-8 text-coolGray-700 md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                    {event.description}
                  </p>
                </Link>
              </article>
            </li>
          ))}
        </ul>
      </PageSection>
    </>
  )
}

export default withAuthSync(Page)
