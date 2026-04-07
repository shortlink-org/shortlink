import { Suspense } from 'react'

import LinksSearchPageClient from './LinksSearchPageClient'

export default function LinksSearchRoutePage() {
  return (
    <Suspense fallback={<div className="min-h-[40vh]" aria-busy aria-label="Loading search" />}>
      <LinksSearchPageClient />
    </Suspense>
  )
}
