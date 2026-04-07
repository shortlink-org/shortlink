import Link from 'next/link'

const sections = [
  {
    title: 'What we collect',
    body: 'We collect account details you provide, basic usage analytics, and link performance data needed to operate the service.',
  },
  {
    title: 'How we use data',
    body: 'We use your data to provide link management, analytics, and security features. We do not sell your personal information.',
  },
  {
    title: 'Cookies',
    body: 'Cookies keep you signed in and help us understand product performance. You can disable cookies, but core features may stop working.',
  },
  {
    title: 'Data retention',
    body: 'We retain data while your account is active. You can request deletion by contacting support.',
  },
]

const cardSurface =
  'rounded-2xl border border-zinc-300/80 bg-zinc-200/95 shadow-sm dark:rounded-[1.35rem] dark:border-white/[0.08] dark:bg-zinc-800/90 dark:shadow-[0_24px_64px_-40px_rgba(0,0,0,0.55)]'

/* Explicit hex + important: light surfaces use rgb() indigo tokens elsewhere — guarantees contrast on zinc-200 cards */
const btnPrimary =
  'app-focus-ring inline-flex items-center justify-center rounded-xl border-2 border-[#4338ca] !bg-[#4f46e5] px-5 py-2.5 text-sm font-semibold !text-white no-underline shadow-sm transition-colors hover:border-[#4338ca] hover:!bg-[#4338ca] focus-visible:ring-2 focus-visible:ring-[#6366f1] focus-visible:ring-offset-2 focus-visible:ring-offset-zinc-200 dark:border-indigo-400 dark:!bg-indigo-500 dark:hover:border-sky-300 dark:hover:!bg-indigo-400 dark:focus-visible:ring-offset-zinc-900'

const btnSecondary =
  'app-focus-ring inline-flex items-center justify-center rounded-xl border-2 border-indigo-600/70 bg-indigo-50 px-5 py-2.5 text-sm font-semibold text-indigo-800 transition-colors hover:bg-indigo-100 dark:border-white/40 dark:bg-transparent dark:text-white dark:hover:bg-white/10 dark:focus-visible:ring-white/40 dark:focus-visible:ring-offset-zinc-900'

export default function PrivacyPage() {
  return (
    <div className="min-h-full bg-zinc-100 text-zinc-900 dark:bg-[#060a12] dark:text-white">
      <section className="mx-auto max-w-3xl px-4 py-12 sm:px-6 sm:py-16 md:max-w-4xl md:py-20 lg:px-8">
        <div className="flex flex-col gap-8 sm:gap-10">
          <div className={`${cardSurface} p-8 sm:p-9 md:p-10`}>
            <p className="text-[11px] font-semibold uppercase tracking-[0.22em] text-indigo-700 dark:text-sky-400">
              Privacy
            </p>
            <h1 className="mt-5 text-3xl font-bold tracking-tight sm:text-4xl md:text-[2.75rem] md:leading-tight">
              Privacy policy
            </h1>
            <p className="mt-4 max-w-2xl text-base leading-relaxed text-zinc-600 dark:text-zinc-400 md:text-lg">
              This page explains what data ShortLink collects and how it is used. For questions, reach out to our support team.
            </p>
          </div>

          <div className="flex flex-col gap-3 sm:gap-3.5">
            {sections.map((section) => (
              <div key={section.title} className={`${cardSurface} px-5 py-5 sm:px-6 sm:py-6`}>
                <h2 className="text-lg font-semibold text-zinc-900 dark:text-white sm:text-xl">{section.title}</h2>
                <p className="mt-3 text-sm leading-relaxed text-zinc-600 dark:text-zinc-400 sm:text-base">{section.body}</p>
              </div>
            ))}
          </div>

          <div
            className={`${cardSurface} flex flex-col items-start justify-between gap-6 p-8 sm:flex-row sm:items-center sm:p-9`}
          >
            <div>
              <h2 className="text-xl font-semibold text-zinc-900 sm:text-2xl dark:text-white">Have a privacy question?</h2>
              <p className="mt-2 max-w-md text-sm leading-relaxed text-zinc-600 dark:text-zinc-400">
                Contact support or review the FAQ for more details.
              </p>
            </div>
            <div className="flex w-full flex-col gap-3 sm:w-auto sm:flex-row sm:flex-wrap">
              <Link href="/contact" className={`${btnPrimary} w-full justify-center sm:w-auto`}>
                Contact support
              </Link>
              <Link href="/faq" className={`${btnSecondary} w-full justify-center sm:w-auto`}>
                View FAQ
              </Link>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}
