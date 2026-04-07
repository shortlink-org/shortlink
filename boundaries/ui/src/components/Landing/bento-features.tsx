'use client'

import Image from 'next/image'

export function BentoFeatures() {
  return (
    <div className="rounded-3xl bg-zinc-50 pt-6 pb-5 sm:pt-8 sm:pb-6 lg:pt-8 lg:pb-8 dark:bg-zinc-950/60">
      <div className="mx-auto w-full max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col gap-4 lg:grid lg:grid-cols-3 lg:grid-rows-2 lg:gap-4">
          {/* Mobile — tall left (desktop col 1) */}
          <div className="relative order-2 lg:order-none lg:row-span-2">
            <div className="absolute inset-px rounded-lg bg-white lg:rounded-l-[2rem] dark:bg-zinc-900" />
            <div className="relative flex h-full flex-col overflow-hidden rounded-lg bg-white lg:rounded-l-[2rem] dark:bg-zinc-900">
              <div className="px-8 pt-8 pb-3 sm:px-10 sm:pt-10 sm:pb-0">
                <p className="mt-2 text-lg font-medium tracking-tight text-neutral-950 max-lg:text-center dark:text-white">
                  Mobile friendly
                </p>
                <p className="mt-2 max-w-lg text-sm/6 text-neutral-600 max-lg:text-center dark:text-zinc-400">
                  Manage and share links from your phone. Responsive UI and apps that stay in sync with your workspace.
                </p>
              </div>
              <div className="@container relative min-h-64 w-full grow max-lg:mx-auto max-lg:max-w-sm sm:min-h-72 lg:min-h-[24rem] xl:min-h-[26rem]">
                <div className="absolute inset-x-10 top-10 bottom-0 overflow-hidden rounded-t-[12cqw] border-x-[3cqw] border-t-[3cqw] border-zinc-700 bg-zinc-900 shadow-2xl dark:border-zinc-600">
                  <Image
                    src="https://images.unsplash.com/photo-1512941937669-90a1b58e7e9c?auto=format&fit=crop&w=800&q=80"
                    alt=""
                    width={800}
                    height={1200}
                    className="size-full object-cover object-top"
                    sizes="(min-width: 1024px) 28vw, 320px"
                  />
                </div>
              </div>
            </div>
            <div className="pointer-events-none absolute inset-px rounded-lg shadow-sm outline -outline-offset-1 outline-black/5 lg:rounded-l-[2rem] dark:outline-white/10" />
          </div>

          {/* Performance — first on mobile */}
          <div className="relative order-1 lg:order-none">
            <div className="absolute inset-px rounded-lg bg-white max-lg:rounded-t-[2rem] dark:bg-zinc-900" />
            <div className="relative flex h-full flex-col overflow-hidden rounded-lg bg-white max-lg:rounded-t-[2rem] dark:bg-zinc-900">
              <div className="px-8 pt-8 sm:px-10 sm:pt-10">
                <p className="mt-2 text-lg font-medium tracking-tight text-neutral-950 max-lg:text-center dark:text-white">
                  Performance
                </p>
                <p className="mt-2 max-w-lg text-sm/6 text-neutral-600 max-lg:text-center dark:text-zinc-400">
                  Fast redirects and low-latency APIs so every click lands quickly, at any scale you grow into.
                </p>
              </div>
              <div className="flex flex-1 items-center justify-center px-8 max-lg:pt-10 max-lg:pb-12 sm:px-10 lg:pb-2">
                <Image
                  src="https://images.unsplash.com/photo-1551288049-bebda4e38f71?auto=format&fit=crop&w=900&q=80"
                  alt=""
                  width={900}
                  height={560}
                  className="w-full max-lg:max-w-xs rounded-lg object-cover shadow-md dark:opacity-95"
                  sizes="(min-width: 1024px) 28vw, 320px"
                />
              </div>
            </div>
            <div className="pointer-events-none absolute inset-px rounded-lg shadow-sm outline -outline-offset-1 outline-black/5 max-lg:rounded-t-[2rem] dark:outline-white/10" />
          </div>

          {/* Security */}
          <div className="relative order-3 lg:order-none lg:col-start-2 lg:row-start-2">
            <div className="absolute inset-px rounded-lg bg-white dark:bg-zinc-900" />
            <div className="relative flex h-full flex-col overflow-hidden rounded-lg bg-white dark:bg-zinc-900">
              <div className="px-8 pt-8 sm:px-10 sm:pt-10">
                <p className="mt-2 text-lg font-medium tracking-tight text-neutral-950 max-lg:text-center dark:text-white">
                  Security
                </p>
                <p className="mt-2 max-w-lg text-sm/6 text-neutral-600 max-lg:text-center dark:text-zinc-400">
                  Open source you can audit, sensible defaults, and controls that respect privacy and your data.
                </p>
              </div>
              <div className="@container flex flex-1 items-center max-lg:py-6 lg:pb-2">
                <Image
                  src="https://images.unsplash.com/photo-1563986768609-322da13575f3?auto=format&fit=crop&w=800&q=80"
                  alt=""
                  width={800}
                  height={400}
                  className="h-[min(152px,40cqw)] w-full object-cover sm:h-[min(180px,42cqw)]"
                  sizes="(min-width: 1024px) 28vw, 100vw"
                />
              </div>
            </div>
            <div className="pointer-events-none absolute inset-px rounded-lg shadow-sm outline -outline-offset-1 outline-black/5 dark:outline-white/10" />
          </div>

          {/* APIs — tall right */}
          <div className="relative order-4 lg:order-none lg:row-span-2">
            <div className="absolute inset-px rounded-lg bg-white max-lg:rounded-b-[2rem] lg:rounded-r-[2rem] dark:bg-zinc-900" />
            <div className="relative flex h-full flex-col overflow-hidden rounded-lg bg-white max-lg:rounded-b-[2rem] lg:rounded-r-[2rem] dark:bg-zinc-900">
              <div className="px-8 pt-8 pb-3 sm:px-10 sm:pt-10 sm:pb-0">
                <p className="mt-2 text-lg font-medium tracking-tight text-neutral-950 max-lg:text-center dark:text-white">
                  Powerful APIs
                </p>
                <p className="mt-2 max-w-lg text-sm/6 text-neutral-600 max-lg:text-center dark:text-zinc-400">
                  Automate link creation, fetch analytics, and integrate Shortlink into your stack with a clear HTTP API.
                </p>
              </div>
              <div className="relative min-h-64 w-full grow sm:min-h-72 lg:min-h-[24rem] xl:min-h-[26rem]">
                <div className="absolute top-10 right-0 bottom-0 left-10 overflow-hidden rounded-tl-xl bg-zinc-900 shadow-2xl outline -outline-offset-1 outline-white/10 dark:bg-zinc-950">
                  <div className="flex bg-zinc-900 outline -outline-offset-1 outline-white/5 dark:bg-zinc-950">
                    <div className="-mb-px flex text-sm/6 font-medium text-zinc-400">
                      <div className="border-r border-b border-r-white/10 border-b-white/20 bg-white/5 px-4 py-2 text-white">
                        create_link.go
                      </div>
                      <div className="border-r border-zinc-600/10 px-4 py-2">client.ts</div>
                    </div>
                  </div>
                  <div className="px-4 py-4 font-mono text-[11px] leading-relaxed text-zinc-300 sm:px-6 sm:pt-6 sm:pb-10 sm:text-xs">
                    <pre className="overflow-x-auto whitespace-pre text-left">
{`POST /api/links
Content-Type: application/json

{
  "url": "https://example.com/page",
  "title": "Campaign landing"
}

→ 201 { "hash": "a1b2c3", "url": "…" }`}
                    </pre>
                  </div>
                </div>
              </div>
            </div>
            <div className="pointer-events-none absolute inset-px rounded-lg shadow-sm outline -outline-offset-1 outline-black/5 max-lg:rounded-b-[2rem] lg:rounded-r-[2rem] dark:outline-white/10" />
          </div>
        </div>
      </div>
    </div>
  )
}
