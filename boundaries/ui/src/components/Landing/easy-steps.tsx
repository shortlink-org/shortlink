'use client'

import type { ComponentType } from 'react'
import {
  CodeBracketIcon,
  PaintBrushIcon,
  ShieldCheckIcon,
  SparklesIcon,
} from '@heroicons/react/24/outline'
import Balancer from 'react-wrap-balancer'

type StepIcon = ComponentType<{ className?: string }>

const steps: {
  title: string
  icon: StepIcon
  accent: string
  iconWrap: string
}[] = [
  {
    title: "We're a friendly, open source project that respects your privacy",
    icon: ShieldCheckIcon,
    accent: 'text-emerald-600 dark:text-emerald-400',
    iconWrap: 'bg-emerald-500/10 ring-emerald-500/20 dark:bg-emerald-400/10 dark:ring-emerald-400/25',
  },
  {
    title: 'Coded by developers for developers',
    icon: CodeBracketIcon,
    accent: 'text-violet-600 dark:text-violet-400',
    iconWrap: 'bg-violet-500/10 ring-violet-500/20 dark:bg-violet-400/10 dark:ring-violet-400/25',
  },
  {
    title: 'We use cutting edge technologies to give you the best experience',
    icon: SparklesIcon,
    accent: 'text-amber-600 dark:text-amber-400',
    iconWrap: 'bg-amber-500/10 ring-amber-500/20 dark:bg-amber-400/10 dark:ring-amber-400/25',
  },
  {
    title: 'High quality UI you can rely on',
    icon: PaintBrushIcon,
    accent: 'text-sky-600 dark:text-sky-400',
    iconWrap: 'bg-sky-500/10 ring-sky-500/20 dark:bg-sky-400/10 dark:ring-sky-400/25',
  },
]

export function EasySteps() {
  return (
    <section
      className="w-full min-w-0 pt-8 pb-14 sm:pt-10 sm:pb-16 lg:pt-10 lg:pb-20"
      aria-labelledby="easy-steps-heading"
    >
      <div className="@container relative isolate overflow-hidden rounded-[2rem] border border-zinc-200/90 bg-zinc-50/80 shadow-sm ring-1 ring-zinc-950/5 backdrop-blur-md dark:border-white/10 dark:bg-zinc-950/80 dark:shadow-[0_1px_0_0_rgb(255_255_255/0.06)_inset] dark:ring-white/10">
        {/* Mesh / noise-style glow */}
        <div
          className="pointer-events-none absolute -top-24 left-1/2 h-72 w-[min(100%,48rem)] -translate-x-1/2 rounded-full bg-[radial-gradient(closest-side,rgb(99_102_241/0.22),transparent)] blur-3xl dark:bg-[radial-gradient(closest-side,rgb(129_140_248/0.18),transparent)]"
          aria-hidden
        />

        <div className="relative px-5 py-12 sm:px-8 sm:py-16 lg:px-12 lg:py-20 xl:px-16 xl:py-24">
          <header className="mx-auto max-w-3xl text-center">
            <p className="inline-flex items-center justify-center rounded-full border border-indigo-200/90 bg-indigo-50/90 px-4 py-1.5 text-[11px] font-semibold uppercase tracking-[0.22em] text-indigo-800 shadow-xs dark:border-indigo-400/25 dark:bg-indigo-500/10 dark:text-indigo-200">
              In a few easy steps
            </p>
            <h2
              id="easy-steps-heading"
              className="mt-6 text-balance bg-gradient-to-b from-zinc-900 to-zinc-600 bg-clip-text text-3xl font-semibold tracking-tight text-transparent sm:mt-7 sm:text-4xl lg:text-[2.85rem] lg:leading-[1.08] dark:from-white dark:to-zinc-400"
            >
              <Balancer>Create beautiful short links &amp; powerful link management</Balancer>
            </h2>
            <p className="mx-auto mt-5 max-w-xl text-pretty text-sm leading-relaxed text-zinc-600 dark:text-zinc-400 sm:text-base lg:text-lg">
              Everything you need to shorten, organize, and understand your traffic — without friction.
            </p>
          </header>

          <ol className="mt-12 grid gap-5 sm:mt-16 sm:grid-cols-2 sm:gap-6 xl:mt-20 xl:grid-cols-4 xl:gap-5">
            {steps.map((step, index) => {
              const Icon = step.icon
              const n = index + 1
              return (
                <li key={step.title} className="min-w-0">
                  <div
                    className="group relative flex h-full flex-col rounded-3xl border border-zinc-200/90 bg-white/90 p-6 shadow-xs ring-1 ring-zinc-950/[0.04] transition-[transform,box-shadow] duration-300 ease-out hover:-translate-y-1 hover:shadow-lg sm:p-7 lg:p-8 dark:border-white/10 dark:bg-zinc-900/60 dark:ring-white/[0.06] dark:hover:bg-zinc-900/80"
                  >
                    <div className="flex items-start justify-between gap-3">
                      <span
                        className={`inline-flex h-11 w-11 shrink-0 items-center justify-center rounded-2xl ring-1 ${step.iconWrap}`}
                      >
                        <Icon className={`h-5 w-5 ${step.accent}`} />
                      </span>
                      <span className="tabular-nums text-3xl font-semibold tracking-tight text-zinc-200 dark:text-zinc-700">
                        {String(n).padStart(2, '0')}
                      </span>
                    </div>
                    <p className="mt-5 text-sm font-medium leading-snug text-zinc-900 sm:text-[0.95rem] dark:text-zinc-100">
                      {step.title}
                    </p>
                    <span
                      className="mt-4 h-0.5 w-8 rounded-full bg-gradient-to-r from-indigo-500 to-violet-500 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
                      aria-hidden
                    />
                  </div>
                </li>
              )
            })}
          </ol>
        </div>
      </div>
    </section>
  )
}
