'use client'

import Link from 'next/link'

const faqs = [
  {
    question: 'What is ShortLink?',
    answer:
      'ShortLink is a link management platform that helps you shorten, organize, and track links across campaigns and teams.',
  },
  {
    question: 'How do I create a short link?',
    answer:
      'Open the Add Link page, paste your long URL, pick your custom slug (optional), and save. Your short link is ready to share.',
  },
  {
    question: 'Can I track analytics for my links?',
    answer:
      'Yes. You can view clicks, referrers, and other metrics on your Links page and in reports.',
  },
  {
    question: 'Do you support custom domains?',
    answer:
      'Custom domains are supported for teams and enterprise plans. Contact support to connect your domain.',
  },
  {
    question: 'Is there a free plan?',
    answer:
      'Yes, the free plan includes core shortening features and basic analytics.',
  },
]

export default function FaqPage() {
  return (
    <div className="min-h-full bg-slate-50 dark:bg-gray-950">
      <section className="mx-auto max-w-5xl px-6 py-16 md:py-24">
        <div className="flex flex-col gap-10">
          <div className="rounded-3xl border border-indigo-100/60 bg-white/80 p-10 shadow-[0_30px_80px_-60px_rgba(79,70,229,0.6)] backdrop-blur dark:border-indigo-900/50 dark:bg-gray-900/70">
            <p className="text-xs uppercase tracking-[0.3em] text-indigo-500">Need help?</p>
            <h1 className="mt-4 text-4xl font-semibold text-gray-900 dark:text-white md:text-5xl">
              Frequently Asked Questions
            </h1>
            <p className="mt-4 max-w-2xl text-base text-gray-600 dark:text-gray-300 md:text-lg">
              Everything you need to know about ShortLink, from creating links to managing analytics and domains.
            </p>
          </div>

          <div className="grid gap-4">
            {faqs.map((item) => (
              <details
                key={item.question}
                className="group rounded-2xl border border-gray-200/70 bg-white/90 p-6 shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-lg dark:border-gray-800 dark:bg-gray-900/70"
              >
                <summary className="flex cursor-pointer list-none items-center justify-between gap-4 text-lg font-semibold text-gray-900 dark:text-white">
                  {item.question}
                  <span className="inline-flex h-8 w-8 items-center justify-center rounded-full border border-indigo-200 text-indigo-600 transition-transform group-open:rotate-45 dark:border-indigo-800 dark:text-indigo-300">
                    +
                  </span>
                </summary>
                <p className="mt-4 text-base leading-relaxed text-gray-600 dark:text-gray-300">{item.answer}</p>
              </details>
            ))}
          </div>

          <div className="flex flex-col items-start justify-between gap-6 rounded-2xl border border-indigo-100/70 bg-indigo-50/80 p-8 text-gray-800 shadow-inner dark:border-indigo-900/60 dark:bg-indigo-950/40 dark:text-gray-200 md:flex-row md:items-center">
            <div>
              <h2 className="text-2xl font-semibold">Still have questions?</h2>
              <p className="mt-2 text-sm text-gray-600 dark:text-gray-300">
                Reach out to our team and we’ll get back within one business day.
              </p>
            </div>
            <Link
              href="/contact"
              className="inline-flex items-center justify-center rounded-xl border border-indigo-200 px-6 py-3 text-sm font-semibold text-indigo-600 transition hover:border-indigo-300 hover:bg-indigo-100/60 hover:text-indigo-700 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 dark:border-indigo-800 dark:text-indigo-200 dark:hover:border-indigo-600 dark:hover:bg-indigo-900/40 dark:hover:text-indigo-100"
            >
              Contact support
            </Link>
          </div>
        </div>
      </section>
    </div>
  )
}
