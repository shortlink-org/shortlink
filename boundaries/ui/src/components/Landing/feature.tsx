'use client'

import type { ComponentType } from 'react'
import {
  ComputerDesktopIcon,
  MagnifyingGlassIcon,
  PaintBrushIcon,
} from '@heroicons/react/24/solid'
import { FeatureCard } from '@shortlink-org/ui-kit'

import { BentoFeatures } from './bento-features'
import { EasySteps } from './easy-steps'

type IconComp = ComponentType<{ className?: string }>
type LinkColor = 'indigo' | 'purple' | 'pink'

const services: {
  title: string
  body: string
  icon: IconComp
  iconGradient: { from: string; to: string }
  decorationGradient: { from: string; to: string }
  linkColor: LinkColor
}[] = [
  {
    title: 'Website design',
    body: "Today's web patterns, integrated into solutions that fit your product.",
    icon: ComputerDesktopIcon,
    iconGradient: { from: 'indigo-500', to: 'blue-600' },
    decorationGradient: { from: 'indigo-500/10', to: 'cyan-500/5' },
    linkColor: 'indigo',
  },
  {
    title: 'Branding',
    body: 'Clear brand stories that connect with your audience.',
    icon: PaintBrushIcon,
    iconGradient: { from: 'purple-500', to: 'pink-600' },
    decorationGradient: { from: 'purple-500/10', to: 'pink-500/5' },
    linkColor: 'purple',
  },
  {
    title: 'SEO marketing',
    body: 'Grow visibility with search-focused content and structure.',
    icon: MagnifyingGlassIcon,
    iconGradient: { from: 'pink-500', to: 'rose-600' },
    decorationGradient: { from: 'rose-500/10', to: 'pink-500/5' },
    linkColor: 'pink',
  },
]

export default function Feature() {
  return (
    <>
      <section
        id="features"
        className="w-full min-w-0 scroll-mt-[var(--app-scroll-margin)] pt-16 pb-2 sm:pb-3 lg:pt-20 lg:pb-4"
        aria-labelledby="features-heading"
      >
        <header className="mx-auto mb-14 max-w-2xl text-center sm:mb-16 lg:mb-12">
          <span className="mb-5 inline-block rounded-full border border-indigo-200/80 bg-indigo-50 px-4 py-1.5 text-xs font-semibold uppercase tracking-wider text-indigo-700 dark:border-indigo-500/30 dark:bg-indigo-950/50 dark:text-indigo-300">
            Features
          </span>
          <h2
            id="features-heading"
            className="text-balance text-3xl font-extrabold tracking-tight text-neutral-900 dark:text-white sm:text-4xl lg:text-5xl xl:text-[3.25rem] xl:leading-[1.1]"
          >
            Powerful link management
          </h2>
          <p className="mt-5 text-pretty text-lg text-neutral-600 sm:text-xl dark:text-zinc-400">
            Everything you need to create, manage, and track your short links
          </p>
        </header>

        <BentoFeatures />
      </section>

      <EasySteps />

      <section className="w-full min-w-0 py-14 lg:py-20" aria-labelledby="services-heading">
        <h2 id="services-heading" className="sr-only">
          Services
        </h2>
        <ul className="grid gap-6 sm:grid-cols-2 sm:gap-8 lg:grid-cols-3 lg:gap-8">
          {services.map((s) => {
            const Icon = s.icon
            return (
              <li key={s.title} className="min-h-0">
                <FeatureCard
                  icon={<Icon className="h-8 w-8 text-white" aria-hidden />}
                  title={s.title}
                  description={s.body}
                  eyebrow="Service"
                  href="/faq"
                  linkText="Learn more"
                  linkColor={s.linkColor}
                  iconGradient={s.iconGradient}
                  decorationGradient={s.decorationGradient}
                  className="h-full mx-2 my-3 sm:mx-3 sm:my-4 lg:mx-4 lg:my-5"
                />
              </li>
            )
          })}
        </ul>
      </section>
    </>
  )
}
