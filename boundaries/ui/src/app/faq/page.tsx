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

const cardSurface =
  'rounded-2xl border border-zinc-300/80 bg-zinc-200/95 shadow-sm dark:rounded-[1.35rem] dark:border-white/[0.08] dark:bg-zinc-800/90 dark:shadow-[0_24px_64px_-40px_rgba(0,0,0,0.55)]'

export default function FaqPage() {
  return (
    <div className="min-h-full bg-zinc-100 text-zinc-900 dark:bg-[#060a12] dark:text-white">
      <section className="mx-auto max-w-3xl px-4 py-12 sm:px-6 sm:py-16 md:max-w-4xl md:py-20 lg:px-8">
        <div className="flex flex-col gap-8 sm:gap-10">
          <div className={`${cardSurface} p-8 sm:p-9 md:p-10`}>
            <p className="text-[11px] font-semibold uppercase tracking-[0.22em] text-indigo-700 dark:text-sky-400">
              Need help?
            </p>
            <h1 className="mt-5 text-3xl font-bold tracking-tight sm:text-4xl md:text-[2.75rem] md:leading-tight">
              Frequently Asked Questions
            </h1>
            <p className="mt-4 max-w-2xl text-base leading-relaxed text-zinc-600 dark:text-zinc-400 md:text-lg">
              Everything you need to know about ShortLink, from creating links to managing analytics and domains.
            </p>
          </div>

          <div className="flex flex-col gap-3 sm:gap-3.5">
            {faqs.map((item) => (
              <details
                key={item.question}
                className={`${cardSurface} group px-5 py-4 transition-[box-shadow] open:shadow-md open:dark:shadow-[0_28px_72px_-42px_rgba(0,0,0,0.65)] sm:px-6 sm:py-5`}
              >
                <summary className="app-focus-ring flex cursor-pointer list-none items-center justify-between gap-4 rounded-lg text-base font-medium text-zinc-900 outline-offset-4 dark:text-white sm:text-lg [&::-webkit-details-marker]:hidden">
                  <span className="min-w-0 pr-2">{item.question}</span>
                  <span
                    className="inline-flex h-10 w-10 shrink-0 items-center justify-center rounded-full border-2 border-indigo-600/40 bg-indigo-50 text-xl font-medium leading-none text-indigo-700 transition-transform duration-200 group-open:rotate-45 group-open:border-indigo-600/50 dark:border-white/35 dark:bg-white/15 dark:text-white dark:shadow-[inset_0_1px_0_rgba(255,255,255,0.12)] dark:group-open:bg-white/20"
                    aria-hidden
                  >
                    +
                  </span>
                </summary>
                <p className="mt-4 border-t border-zinc-300/70 pt-4 text-base leading-relaxed text-zinc-600 dark:border-white/[0.08] dark:text-zinc-400">
                  {item.answer}
                </p>
              </details>
            ))}
          </div>

          <div
            className={`${cardSurface} flex flex-col items-start justify-between gap-6 p-8 sm:flex-row sm:items-center sm:p-9`}
          >
            <div>
              <h2 className="text-xl font-semibold text-zinc-900 sm:text-2xl dark:text-white">Still have questions?</h2>
              <p className="mt-2 max-w-md text-sm leading-relaxed text-zinc-600 dark:text-zinc-400">
                Reach out to our team and we’ll get back within one business day.
              </p>
            </div>
            <Link
              href="/contact"
              className="app-focus-ring inline-flex w-full shrink-0 items-center justify-center rounded-xl border-2 border-[#4338ca] !bg-[#4f46e5] px-6 py-3 text-sm font-semibold !text-white no-underline shadow-sm transition-colors hover:border-[#4338ca] hover:!bg-[#4338ca] focus-visible:ring-2 focus-visible:ring-[#6366f1] focus-visible:ring-offset-2 focus-visible:ring-offset-zinc-200 dark:border-indigo-400 dark:!bg-indigo-500 dark:hover:border-sky-300 dark:hover:!bg-indigo-400 dark:focus-visible:ring-offset-zinc-900 sm:w-auto"
            >
              Contact support
            </Link>
          </div>
        </div>
      </section>
    </div>
  )
}
