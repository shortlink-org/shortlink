import Link from 'next/link'

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

export default function ContactPage() {
  return (
    <div className="min-h-full bg-slate-50 dark:bg-gray-950">
      <section className="mx-auto max-w-5xl px-6 py-16 md:py-24">
        <div className="flex flex-col gap-10">
          <div className="rounded-3xl border border-indigo-100/60 bg-white/80 p-10 shadow-[0_30px_80px_-60px_rgba(79,70,229,0.6)] backdrop-blur dark:border-indigo-900/50 dark:bg-gray-900/70">
            <p className="text-xs uppercase tracking-[0.3em] text-indigo-500">Contact</p>
            <h1 className="mt-4 text-4xl font-semibold text-gray-900 dark:text-white md:text-5xl">
              Talk to our support team
            </h1>
            <p className="mt-4 max-w-2xl text-base text-gray-600 dark:text-gray-300 md:text-lg">
              We typically respond within one business day. Choose the channel that works best for you.
            </p>
          </div>

          <div className="grid gap-4 md:grid-cols-2">
            {channels.map((channel) => (
              <div
                key={channel.title}
                className="flex h-full flex-col justify-between rounded-2xl border border-gray-200/70 bg-white/90 p-6 shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-lg dark:border-gray-800 dark:bg-gray-900/70"
              >
                <div>
                  <h2 className="text-lg font-semibold text-gray-900 dark:text-white">{channel.title}</h2>
                  <p className="mt-3 text-sm text-gray-600 dark:text-gray-300">{channel.description}</p>
                </div>
                {channel.href.startsWith('/') ? (
                  <Link
                    href={channel.href}
                    className="mt-6 inline-flex items-center text-sm font-semibold text-indigo-600 hover:text-indigo-500 dark:text-indigo-300"
                  >
                    {channel.action}
                  </Link>
                ) : (
                  <a
                    href={channel.href}
                    className="mt-6 inline-flex items-center text-sm font-semibold text-indigo-600 hover:text-indigo-500 dark:text-indigo-300"
                  >
                    {channel.action}
                  </a>
                )}
              </div>
            ))}
          </div>

          <div className="rounded-2xl border border-indigo-100/70 bg-indigo-50/70 p-8 text-gray-800 shadow-inner dark:border-indigo-900/60 dark:bg-indigo-950/40 dark:text-gray-200">
            <h2 className="text-2xl font-semibold">Need a quicker response?</h2>
            <p className="mt-2 text-sm text-gray-600 dark:text-gray-300">
              Include your account email, relevant link IDs, and a clear description of the issue to help us triage faster.
            </p>
            <a
              href="mailto:support@shortlink.best?subject=ShortLink%20Support"
              className="mt-6 inline-flex items-center justify-center rounded-xl border border-indigo-200 px-6 py-3 text-sm font-semibold text-indigo-600 transition hover:border-indigo-300 hover:bg-indigo-100/60 hover:text-indigo-700 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2 dark:border-indigo-800 dark:text-indigo-200 dark:hover:border-indigo-600 dark:hover:bg-indigo-900/40 dark:hover:text-indigo-100"
            >
              Email support
            </a>
          </div>
        </div>
      </section>
    </div>
  )
}
