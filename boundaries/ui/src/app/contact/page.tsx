import Link from 'next/link'

import { Button } from '@shortlink-org/ui-kit'

import styles from './page.module.css'

const channels = [
  {
    title: 'Email support',
    description: 'Reach our team directly for account or billing questions.',
    action: 'support@shortlink.best',
    href: 'mailto:support@shortlink.best?subject=ShortLink%20Support',
  },
  {
    title: 'Product docs',
    description: 'Find guides and integration details for ShortLink.',
    action: 'Read the FAQ',
    href: '/faq',
  },
]

function ArrowIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" aria-hidden>
      <path strokeLinecap="round" strokeLinejoin="round" d="M13.5 4.5 21 12m0 0-7.5 7.5M21 12H3" />
    </svg>
  )
}

export default function ContactPage() {
  return (
    <div className="min-h-full bg-zinc-50 dark:bg-neutral-950">
      <main className={`${styles.contactMain} py-14 sm:py-16 md:py-16 lg:py-20`}>
        {/* Hero */}
        <div className="@container relative isolate overflow-hidden rounded-[2rem] border border-zinc-200/90 bg-zinc-50/80 shadow-sm ring-1 ring-zinc-950/5 backdrop-blur-md dark:border-white/10 dark:bg-zinc-950/80 dark:shadow-[0_1px_0_0_rgb(255_255_255/0.06)_inset] dark:ring-white/10">
          <div
            className="pointer-events-none absolute -top-28 left-1/2 h-80 w-[min(100%,42rem)] -translate-x-1/2 rounded-full bg-[radial-gradient(closest-side,rgb(99_102_241/0.22),transparent)] blur-3xl dark:bg-[radial-gradient(closest-side,rgb(129_140_248/0.16),transparent)]"
            aria-hidden
          />
          <div className="pointer-events-none absolute inset-x-0 bottom-0 h-px bg-gradient-to-r from-transparent via-zinc-300/70 to-transparent dark:via-white/10" aria-hidden />
          <div className="relative px-8 py-10 sm:px-10 sm:py-12 lg:px-12 lg:py-14">
            <p className="text-[11px] font-semibold uppercase tracking-[0.22em] text-indigo-600 dark:text-indigo-400">
              Contact
            </p>
            <h1 className="mt-4 text-balance text-3xl font-semibold tracking-tight text-zinc-900 sm:text-4xl md:text-5xl dark:text-white">
              Talk to our support team
            </h1>
            <p className="mt-4 max-w-2xl text-pretty text-base leading-relaxed text-zinc-600 sm:text-lg dark:text-zinc-400">
              We typically respond within one business day. Choose the channel that works best for you.
            </p>
          </div>
        </div>

        {/* Full-bleed band: background to viewport edges; cards align with hero/CTA via subgrid */}
        <section
          className={`${styles.contactFullBleed} py-8 md:py-10 bg-zinc-200/35 dark:bg-white/[0.04]`}
          aria-label="Contact channels"
        >
          <div className="grid gap-4 sm:grid-cols-2 sm:gap-5 lg:gap-6">
            {channels.map((channel) => (
              <div
                key={channel.title}
                className="group relative overflow-hidden rounded-2xl border border-zinc-200/90 bg-white/90 p-7 shadow-xs ring-1 ring-zinc-950/[0.04] transition-[transform,box-shadow] duration-300 ease-out hover:-translate-y-1 hover:shadow-lg dark:border-white/10 dark:bg-zinc-900/70 dark:ring-white/[0.06] dark:hover:bg-zinc-900/85 sm:p-8"
              >
                <div
                  className="pointer-events-none absolute -right-12 -top-12 h-40 w-40 rounded-full bg-gradient-to-br from-indigo-400/15 to-violet-500/10 blur-2xl transition-opacity duration-300 group-hover:opacity-100 dark:from-indigo-500/20 dark:to-violet-600/10"
                  aria-hidden
                />
                <div className="relative">
                  <h2 className="text-lg font-semibold tracking-tight text-zinc-900 dark:text-white">{channel.title}</h2>
                  <p className="mt-3 text-sm leading-relaxed text-zinc-600 dark:text-zinc-400">{channel.description}</p>
                  {channel.href.startsWith('/') ? (
                    <Link
                      href={channel.href}
                      className="mt-6 inline-flex items-center gap-2 text-sm font-semibold text-indigo-600 transition hover:gap-2.5 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
                    >
                      {channel.action}
                      <ArrowIcon className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
                    </Link>
                  ) : (
                    <a
                      href={channel.href}
                      className="mt-6 inline-flex items-center gap-2 text-sm font-semibold text-indigo-600 transition hover:gap-2.5 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
                    >
                      {channel.action}
                      <ArrowIcon className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
                    </a>
                  )}
                </div>
              </div>
            ))}
          </div>
        </section>

        {/* CTA */}
        <div className="relative isolate overflow-hidden rounded-[2rem] border border-indigo-200/90 bg-gradient-to-br from-indigo-50/95 via-white to-violet-50/90 p-8 shadow-sm ring-1 ring-indigo-500/10 dark:border-indigo-500/25 dark:from-indigo-950/60 dark:via-zinc-950 dark:to-violet-950/40 dark:ring-indigo-400/15 sm:p-10">
          <div
            className="pointer-events-none absolute -bottom-20 right-0 h-56 w-56 rounded-full bg-[radial-gradient(closest-side,rgb(139_92_246/0.2),transparent)] blur-3xl dark:bg-[radial-gradient(closest-side,rgb(167_139_250/0.12),transparent)]"
            aria-hidden
          />
          <div className="relative">
            <h2 className="text-xl font-semibold tracking-tight text-zinc-900 sm:text-2xl dark:text-white">
              Need a quicker response?
            </h2>
            <p className="mt-3 max-w-2xl text-sm leading-relaxed text-zinc-600 sm:text-base dark:text-zinc-400">
              Include your account email, relevant link IDs, and a clear description of the issue to help us triage faster.
            </p>
            <Button
              as="a"
              asProps={{
                href: 'mailto:support@shortlink.best?subject=ShortLink%20Support',
              }}
              className="mt-7 bg-indigo-600 shadow-lg shadow-indigo-500/25 hover:bg-indigo-500 hover:shadow-indigo-500/35 focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 dark:bg-indigo-500 dark:shadow-indigo-950/50 dark:hover:bg-indigo-400 dark:focus-visible:ring-offset-zinc-950"
            >
              Email support
              <ArrowIcon className="ml-1 h-4 w-4" />
            </Button>
          </div>
        </div>
      </main>
    </div>
  )
}
