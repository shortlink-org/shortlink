'use client'

function AppleGlyph({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 305 305" aria-hidden>
      <path d="M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z" />
      <path d="M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z" />
    </svg>
  )
}

function PlayGlyph({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 512 512" aria-hidden>
      <path d="M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z" />
    </svg>
  )
}

function PhoneOutline({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24" aria-hidden>
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"
      />
    </svg>
  )
}

function PhoneMockup({
  letter,
  innerClassName,
  rotationClass,
  className,
}: {
  letter: string
  innerClassName: string
  rotationClass: string
  className?: string
}) {
  return (
    <div
      className={`relative shrink-0 ${rotationClass} ${className ?? ''}`}
      style={{ width: 'min(11rem, 42vw)', aspectRatio: '9 / 19' }}
    >
      {/* Bezel */}
      <div
        className="absolute inset-0 rounded-[1.85rem] p-[3px] shadow-[inset_0_1px_0_rgb(255_255_255/0.35),0_12px_40px_-12px_rgb(0_0_0/0.45)] dark:shadow-[inset_0_1px_0_rgb(255_255_255/0.12),0_16px_48px_-8px_rgb(0_0_0/0.85)]"
        style={{
          background: 'linear-gradient(145deg, #d4d4d8 0%, #71717a 45%, #3f3f46 100%)',
        }}
      >
        <div
          className={`flex h-full w-full items-center justify-center rounded-[1.65rem] ${innerClassName}`}
        >
          <span className="text-3xl font-bold tracking-tight text-white drop-shadow-sm sm:text-4xl">{letter}</span>
        </div>
      </div>
    </div>
  )
}

export default function MobileApps() {
  return (
    <section className="w-full min-w-0 py-14 lg:py-20" aria-labelledby="landing-mobile-heading">
      {/* Один скруглённый блок: фон + обрезка контента без «полосы» снизу */}
      <div
        className="relative isolate overflow-hidden rounded-[1.75rem] border border-neutral-200/80 bg-gradient-to-r from-neutral-200 via-neutral-300 to-neutral-500 shadow-sm dark:border-white/10 dark:from-zinc-900 dark:via-zinc-800 dark:to-zinc-950 dark:shadow-[0_24px_80px_-40px_rgb(0_0_0/0.9)]"
      >
        <div
          className="pointer-events-none absolute inset-0 opacity-40 dark:opacity-30"
          style={{
            background:
              'radial-gradient(ellipse 80% 60% at 85% 50%, rgb(99 102 241 / 0.35), transparent 55%), radial-gradient(ellipse 60% 50% at 15% 80%, rgb(168 85 247 / 0.2), transparent 50%)',
          }}
          aria-hidden
        />

        <div className="relative grid gap-12 p-8 sm:p-10 lg:grid-cols-2 lg:items-center lg:gap-16 lg:p-12 xl:p-14">
          <div className="min-w-0">
            <div className="mb-5 flex items-center gap-2">
              <PhoneOutline className="h-5 w-5 shrink-0 text-indigo-600 dark:text-indigo-400" />
              <span className="text-xs font-semibold uppercase tracking-[0.2em] text-indigo-700 dark:text-indigo-400">
                Mobile apps
              </span>
            </div>

            <h2
              id="landing-mobile-heading"
              className="text-balance text-3xl font-bold tracking-tight text-neutral-900 sm:text-4xl lg:text-[2.5rem] lg:leading-tight dark:text-white"
            >
              Shortlink on the go
            </h2>

            <p className="mt-5 max-w-xl text-pretty text-base leading-relaxed text-neutral-700 sm:text-lg dark:text-zinc-300">
              Access your links anywhere, anytime. Our mobile apps sync seamlessly across all your devices — iOS, Android,
              and desktop.
            </p>

            <div className="mt-10 flex flex-col gap-4 sm:flex-row sm:flex-wrap">
              <button
                type="button"
                className="inline-flex min-h-[3.25rem] items-center gap-3 rounded-xl border border-neutral-800/10 bg-neutral-900 px-5 py-3 text-left text-white shadow-lg transition hover:-translate-y-0.5 hover:bg-neutral-800 hover:shadow-xl dark:border-white/15 dark:bg-white dark:text-zinc-900 dark:hover:bg-zinc-100"
              >
                <AppleGlyph className="h-8 w-8 shrink-0 opacity-95" />
                <span className="leading-tight">
                  <span className="block text-[11px] font-medium uppercase tracking-wide text-neutral-400 dark:text-zinc-600">
                    Download on the
                  </span>
                  <span className="block text-base font-semibold">App Store</span>
                </span>
              </button>

              <button
                type="button"
                className="inline-flex min-h-[3.25rem] items-center gap-3 rounded-xl border border-neutral-800/10 bg-neutral-900 px-5 py-3 text-left text-white shadow-lg transition hover:-translate-y-0.5 hover:bg-neutral-800 hover:shadow-xl dark:border-white/15 dark:bg-white dark:text-zinc-900 dark:hover:bg-zinc-100"
              >
                <PlayGlyph className="h-8 w-8 shrink-0 opacity-95" />
                <span className="leading-tight">
                  <span className="block text-[11px] font-medium uppercase tracking-wide text-neutral-400 dark:text-zinc-600">
                    Get it on
                  </span>
                  <span className="block text-base font-semibold">Google Play</span>
                </span>
              </button>
            </div>
          </div>

          {/* Макеты обрезаются скруглением родителя (overflow-hidden на карточке) */}
          <div className="relative flex min-h-[17rem] items-center justify-center overflow-visible lg:min-h-[22rem]">
            <div
              className="pointer-events-none absolute left-1/2 top-1/2 h-[120%] w-[120%] -translate-x-1/2 -translate-y-1/2 rounded-full blur-3xl dark:opacity-70"
              style={{
                background: 'radial-gradient(circle, rgb(99 102 241 / 0.25) 0%, transparent 65%)',
              }}
              aria-hidden
            />
            <div className="relative flex items-center justify-center gap-0 pl-4 pr-2 sm:pl-8 sm:pr-4">
              <PhoneMockup
                letter="S"
                rotationClass="-rotate-6 sm:-rotate-[8deg]"
                innerClassName="bg-gradient-to-br from-violet-400 via-indigo-500 to-indigo-700 dark:from-violet-500 dark:via-indigo-600 dark:to-indigo-900"
              />
              <PhoneMockup
                letter="L"
                rotationClass="z-[1] -ml-10 rotate-6 sm:-ml-14 sm:rotate-[8deg]"
                innerClassName="bg-gradient-to-br from-fuchsia-500 via-purple-600 to-indigo-800 dark:from-fuchsia-600 dark:via-purple-700 dark:to-indigo-950"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
