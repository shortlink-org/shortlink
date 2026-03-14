'use client'

import Link from 'next/link'

import Ready from '@/components/Landing/Ready'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

const plannedIntegrations = [
  {
    name: 'Slack',
    description: 'Ship alerts for new links, suspicious traffic and workspace events directly to your channels.',
    status: 'Design review',
  },
  {
    name: 'Webhook delivery',
    description: 'Push link-created and click events into your own automation pipelines.',
    status: 'In planning',
  },
  {
    name: 'Google Analytics',
    description: 'Attach campaign metadata and compare shortlink traffic with downstream conversion data.',
    status: 'Research',
  },
  {
    name: 'Zapier',
    description: 'Connect shortlink workflows with CRM, email and project tools without custom code.',
    status: 'Backlog',
  },
]

function Intargrations() {
  return (
    <>
      {/*<NextSeo title="Intargrations" description="Intargrations page for your account." />*/}
      <Header title="Integration" description="Connect external services and prepare your workspace for automation." />

      <PageSection className="space-y-6 pb-10">
        <div className="rounded-2xl border border-[var(--color-border)] bg-white p-6 shadow-sm dark:bg-gray-800">
          <div className="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
            <div className="max-w-2xl">
              <p className="text-xs font-semibold uppercase tracking-[0.2em] text-sky-600 dark:text-sky-300">Coming soon</p>
              <h2 className="mt-3 text-2xl font-semibold text-gray-900 dark:text-gray-100">Integration workspace</h2>
              <p className="mt-3 text-sm leading-7 text-gray-600 dark:text-gray-400">
                Native integrations are not enabled yet, but the product direction is clear. We are prioritising notifications,
                analytics connectors and event delivery so teams can automate link operations safely.
              </p>
            </div>
            <div className="flex flex-wrap gap-3">
              <Link
                href="/contact"
                className="inline-flex items-center justify-center rounded-full bg-purple-700 px-5 py-3 text-sm font-semibold text-white transition-colors hover:bg-purple-800 dark:bg-indigo-500 dark:hover:bg-indigo-600"
              >
                Request integration
              </Link>
              <Link
                href="/faq"
                className="inline-flex items-center justify-center rounded-full border border-[var(--color-border)] bg-[var(--color-surface)] px-5 py-3 text-sm font-semibold text-[var(--color-foreground)] transition-colors hover:bg-[var(--color-muted)]"
              >
                Read setup notes
              </Link>
            </div>
          </div>

          <div className="mt-6 grid gap-4 md:grid-cols-2">
            {plannedIntegrations.map((integration) => (
              <div
                key={integration.name}
                className="rounded-2xl border border-[var(--color-border)] bg-[var(--color-background)]/70 p-5"
              >
                <div className="flex items-center justify-between gap-3">
                  <h3 className="text-base font-semibold text-gray-900 dark:text-gray-100">{integration.name}</h3>
                  <span className="rounded-full bg-sky-500/10 px-2.5 py-1 text-[11px] font-semibold uppercase tracking-[0.14em] text-sky-700 dark:text-sky-300">
                    {integration.status}
                  </span>
                </div>
                <p className="mt-3 text-sm leading-6 text-gray-600 dark:text-gray-400">{integration.description}</p>
              </div>
            ))}
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-3">
          <div className="rounded-2xl border border-[var(--color-border)] bg-white p-5 shadow-sm dark:bg-gray-800">
            <p className="text-sm font-semibold text-gray-900 dark:text-gray-100">Delivery model</p>
            <p className="mt-2 text-sm leading-6 text-gray-600 dark:text-gray-400">
              Expect API keys, webhook signing and workspace-scoped permissions from day one.
            </p>
          </div>
          <div className="rounded-2xl border border-[var(--color-border)] bg-white p-5 shadow-sm dark:bg-gray-800">
            <p className="text-sm font-semibold text-gray-900 dark:text-gray-100">Security baseline</p>
            <p className="mt-2 text-sm leading-6 text-gray-600 dark:text-gray-400">
              Integrations will inherit audit logs, revocation controls and least-privilege defaults.
            </p>
          </div>
          <div className="rounded-2xl border border-[var(--color-border)] bg-white p-5 shadow-sm dark:bg-gray-800">
            <p className="text-sm font-semibold text-gray-900 dark:text-gray-100">Feedback loop</p>
            <p className="mt-2 text-sm leading-6 text-gray-600 dark:text-gray-400">
              Share the workflows you need most so we can prioritise the first production adapters.
            </p>
            </div>
        </div>

        <Ready />
      </PageSection>
    </>
  )
}

export default withAuthSync(Intargrations)
