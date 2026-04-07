'use client'

// @ts-ignore — package ships without TypeScript declarations
import { Newsletter } from '@shortlink-org/ui-kit'

export default function Subscribe() {
  return (
    <div className="w-full min-w-0 pb-1 [&_section]:py-14 sm:[&_section]:py-16 lg:[&_section]:py-20 xl:[&_section]:py-24">
      <Newsletter
        eyebrow="Stay updated"
        heading="Subscribe to our newsletter"
        description="Be the first to know about new features, updates, and tips for making the most of Shortlink."
        disclaimer="No spam, unsubscribe at any time."
        emailPlaceholder="Enter your email"
        buttonText="Subscribe"
        onSubmit={async (_email: string) => {
          await new Promise((resolve) => setTimeout(resolve, 900))
        }}
      />
    </div>
  )
}
