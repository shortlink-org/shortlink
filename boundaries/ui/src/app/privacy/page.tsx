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

export default function PrivacyPage() {
  return (
    <div className="min-h-full bg-slate-50 dark:bg-gray-950">
      <section className="mx-auto max-w-5xl px-6 py-16 md:py-24">
        <div className="flex flex-col gap-10">
          <div className="rounded-3xl border border-indigo-100/60 bg-white/80 p-10 shadow-[0_30px_80px_-60px_rgba(79,70,229,0.6)] backdrop-blur dark:border-indigo-900/50 dark:bg-gray-900/70">
            <p className="text-xs uppercase tracking-[0.3em] text-indigo-500">Privacy</p>
            <h1 className="mt-4 text-4xl font-semibold text-gray-900 dark:text-white md:text-5xl">Privacy policy</h1>
            <p className="mt-4 max-w-2xl text-base text-gray-600 dark:text-gray-300 md:text-lg">
              This page explains what data ShortLink collects and how it is used. For questions, reach out to our support team.
            </p>
          </div>

          <div className="grid gap-4">
            {sections.map((section) => (
              <div
                key={section.title}
                className="rounded-2xl border border-gray-200/70 bg-white/90 p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900/70"
              >
                <h2 className="text-lg font-semibold text-gray-900 dark:text-white">{section.title}</h2>
                <p className="mt-3 text-sm text-gray-600 dark:text-gray-300">{section.body}</p>
              </div>
            ))}
          </div>

          <div className="flex flex-col items-start justify-between gap-6 rounded-2xl border border-indigo-100/70 bg-indigo-50/80 p-8 text-gray-800 shadow-inner dark:border-indigo-900/60 dark:bg-indigo-950/40 dark:text-gray-200 md:flex-row md:items-center">
            <div>
              <h2 className="text-2xl font-semibold">Have a privacy question?</h2>
              <p className="mt-2 text-sm text-gray-600 dark:text-gray-300">
                Contact support or review the FAQ for more details.
              </p>
            </div>
            <div className="flex flex-wrap gap-3">
              <Link
                href="/contact"
                className="inline-flex items-center justify-center rounded-xl border border-indigo-200 px-5 py-2.5 text-sm font-semibold text-indigo-600 transition hover:border-indigo-300 hover:bg-indigo-100/60 hover:text-indigo-700 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 dark:border-indigo-800 dark:text-indigo-200 dark:hover:border-indigo-600 dark:hover:bg-indigo-900/40 dark:hover:text-indigo-100"
              >
                Contact support
              </Link>
              <Link
                href="/faq"
                className="inline-flex items-center justify-center rounded-xl border border-indigo-200 px-5 py-2.5 text-sm font-semibold text-indigo-600 transition hover:border-indigo-300 hover:bg-indigo-100/60 hover:text-indigo-700 dark:border-indigo-800 dark:text-indigo-200 dark:hover:border-indigo-600 dark:hover:bg-indigo-900/40 dark:hover:text-indigo-100"
              >
                View FAQ
              </Link>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}
