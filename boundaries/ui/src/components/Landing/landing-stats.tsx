'use client'

import { StatCard } from '@shortlink-org/ui-kit'

import styles from './landing-stats.module.css'

export function LandingStats() {
  return (
    <section className="w-full min-w-0" aria-label="Product highlights">
      <div className={`${styles.grid} grid items-stretch gap-4 sm:grid-cols-3 sm:gap-5`}>
        <StatCard
          className="min-h-[12rem] sm:min-h-[12.5rem]"
          label="Model"
          value="Open source"
          change="MIT · auditable"
          tone="neutral"
          labelClassName="min-h-[2.75rem] flex items-end"
          changeClassName="whitespace-nowrap"
        />
        <StatCard
          className="min-h-[12rem] sm:min-h-[12.5rem]"
          label="Integrate"
          value="HTTP API"
          change="Automate links"
          tone="accent"
          labelClassName="min-h-[2.75rem] flex items-end"
          changeClassName="whitespace-nowrap"
        />
        <StatCard
          className="min-h-[12rem] sm:min-h-[12.5rem]"
          label="Privacy"
          value="You own it"
          change="No ad redirects"
          tone="success"
          labelClassName="min-h-[2.75rem] flex items-end"
          changeClassName="whitespace-nowrap"
        />
      </div>
    </section>
  )
}
