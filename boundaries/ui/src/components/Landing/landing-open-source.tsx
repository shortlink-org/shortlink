'use client'

import { GithubRepository } from '@shortlink-org/ui-kit'

export function LandingOpenSource() {
  return (
    <section
      className="relative w-full min-w-0 sm:-mx-2 md:-mx-3 lg:-mx-4"
      aria-labelledby="landing-open-source-heading"
    >
      <h2 id="landing-open-source-heading" className="sr-only">
        Open source
      </h2>
      <GithubRepository
        url="https://github.com/shortlink-org/shortlink"
        title="shortlink-org/shortlink"
        description="Core services, UI kit, and clients — inspect issues, PRs, and the latest implementation."
        meta="GitHub"
        ctaText="View repository"
        accentColor="#6366f1"
        hoverColor="#818cf8"
        className="shadow-[0_28px_80px_-40px_rgba(15,23,42,0.85)] dark:shadow-[0_32px_90px_-44px_rgba(0,0,0,0.65)]"
      />
    </section>
  )
}
