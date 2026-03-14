'use client'

import React from 'react'
import Link from 'next/link'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

import { Layout } from '@/components'
import Ready from '@/components/Landing/Ready'
import withAuthSync from '@/components/Private'

const reportHighlights = [
  {
    label: 'Sales',
    value: '89.5%',
    delta: '4.3%',
    progressClassName: 'w-1/2 bg-blue-500',
  },
  {
    label: 'Revenue',
    value: '$75,000',
    delta: 'On track',
    progressClassName: 'w-40 bg-lime-500',
  },
  {
    label: 'Customers',
    value: '3,922',
    delta: '9.1%',
    progressClassName: 'w-44 bg-yellow-500',
  },
]

function Page() {
  return (
    <Layout>
      {/*<NextSeo title="Reports" description="Reports page for your account." />*/}

      <Header title="Reports" description="Understand performance trends and key link metrics from one place." />

      <PageSection className="space-y-6 pb-10">
        <div className="rounded-2xl bg-white px-4 py-4 md:px-14 lg:px-8 lg:py-10 dark:bg-gray-800">
          <div className="flex flex-col">
            <p className="text-gray-800 dark:text-gray-300">
              Reporting is a critical part of our shortlink service. Depending on your settings, we can generate a comprehensive report on
              your vitals either daily, weekly, monthly, quarterly or yearly. This report can include things like how many people clicked on
              your links, where they came from, and what kind of device they were using. This information can be extremely valuable in
              understanding your audience and tailoring your content to them. Additionally, we can customize the reports to include only the
              information that you are interested in. Our reporting system is designed to be flexible and user-friendly, so you can get the
              most out of it.
            </p>

            <br />

            <p className="text-gray-800 dark:text-gray-300">
              Whether you're looking to track your progress over time or want to stay informed about the latest changes in your industry, our
              reporting feature will give you the information you need to make decisions that improve your bottom line.
            </p>
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-3">
          {reportHighlights.map(highlight => (
            <div key={highlight.label} className="rounded-2xl bg-white p-6 shadow dark:bg-gray-800">
              <p className="text-xs uppercase leading-none tracking-[0.18em] text-gray-500 dark:text-gray-400">{highlight.label}</p>
              <p className="mt-5 text-lg font-bold leading-3 text-gray-800 dark:text-gray-200 sm:text-xl md:text-2xl lg:text-3xl">
                {highlight.value}
              </p>
              <div className="mt-5 flex flex-col md:w-64">
                <div className="flex w-full justify-end">
                  <div className="flex items-center gap-1">
                    <svg xmlns="http://www.w3.org/2000/svg" width={16} height={16} viewBox="0 0 16 16" fill="none">
                      <path d="M8 3.33334V12.6667" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M12 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M4 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                    <p className="text-xs leading-none text-green-600">{highlight.delta}</p>
                  </div>
                </div>
                <div className="mt-2.5">
                  <div className="h-1 w-full rounded-full bg-gray-200">
                    <div className={`h-1 rounded-full ${highlight.progressClassName}`} />
                  </div>
                </div>
              </div>
              <p className="mt-1.5 text-xs leading-3 text-gray-400 dark:text-gray-300">Yearly target</p>
            </div>
          ))}
        </div>

        <div className="rounded-2xl border border-[var(--color-border)] bg-white p-6 shadow-sm dark:bg-gray-800">
          <div className="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
            <div className="max-w-2xl">
              <p className="text-xs font-semibold uppercase tracking-[0.2em] text-sky-600 dark:text-sky-300">Next step</p>
              <h2 className="mt-3 text-2xl font-semibold text-gray-900 dark:text-gray-100">Build reports that match your workflow</h2>
              <p className="mt-3 text-sm leading-7 text-gray-600 dark:text-gray-400">
                The current dashboard is still a lightweight preview. The next milestone is saved report views, time ranges,
                workspace filters and exports that teams can schedule or share.
              </p>
            </div>
            <div className="flex flex-wrap gap-3">
              <Link
                href="/links"
                className="inline-flex items-center justify-center rounded-full bg-purple-700 px-5 py-3 text-sm font-semibold text-white transition-colors hover:bg-purple-800 dark:bg-indigo-500 dark:hover:bg-indigo-600"
              >
                Review link data
              </Link>
              <Link
                href="/contact"
                className="inline-flex items-center justify-center rounded-full border border-[var(--color-border)] bg-[var(--color-surface)] px-5 py-3 text-sm font-semibold text-[var(--color-foreground)] transition-colors hover:bg-[var(--color-muted)]"
              >
                Request a report
              </Link>
            </div>
          </div>
        </div>

        <Ready />
      </PageSection>
    </Layout>
  )
}

export default withAuthSync(Page)
