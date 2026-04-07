'use client'

import { useState } from 'react'

const testimonials = [
  {
    name: 'Viktor Login',
    role: 'Open source contributor',
    image: 'https://www.tailwind-kit.com/images/person/1.jpg',
    quote:
      'Shortlink has transformed how I manage my URLs. The analytics are incredibly detailed and the interface is intuitive.',
  },
  {
    name: 'Alex Chen',
    role: 'Product engineer',
    image: 'https://www.tailwind-kit.com/images/person/2.jpg',
    quote:
      'Finally a URL shortener that respects privacy and gives you full control. Open source for the win!',
  },
  {
    name: 'Sarah Miller',
    role: 'Automation lead',
    image: 'https://www.tailwind-kit.com/images/person/3.jpg',
    quote: 'The API is fantastic for automation. I integrated it into my workflow in minutes.',
  },
]

export default function Testimonials() {
  const [active, setActive] = useState(0)
  const t = testimonials[active]

  return (
    <section className="w-full min-w-0 py-14 lg:py-20">
      <div className="mb-12 text-center sm:mb-14">
        <span className="text-xs font-semibold uppercase tracking-widest text-indigo-600 dark:text-indigo-400">
          Testimonials
        </span>
        <h2 className="mt-3 text-3xl font-bold text-gray-900 sm:text-4xl lg:text-5xl dark:text-white">
          Loved by developers
        </h2>
        <p className="mx-auto mt-4 max-w-2xl text-base text-gray-600 sm:text-lg dark:text-gray-400">
          See what our users have to say about Shortlink
        </p>
      </div>

      <div
        className="overflow-hidden rounded-[1.25rem] bg-[#0b0f18] px-6 py-8 shadow-[0_24px_64px_-36px_rgba(0,0,0,0.55)] ring-1 ring-white/[0.06] sm:rounded-3xl sm:px-9 sm:py-10 md:px-10 md:py-11 lg:px-12 lg:py-12 xl:px-14 xl:py-14"
        style={{ backgroundImage: 'radial-gradient(120% 80% at 100% 0%, rgba(99,102,241,0.08), transparent 55%)' }}
      >
        <div className="flex flex-col items-stretch gap-10 lg:flex-row lg:items-center lg:gap-12 xl:gap-14">
          <div className="flex shrink-0 justify-center lg:w-[min(100%,280px)] lg:justify-start xl:w-[min(100%,300px)]">
            {/* eslint-disable-next-line @next/next/no-img-element -- remote host not in next/image patterns */}
            <img
              key={t.image}
              src={t.image}
              alt=""
              className="aspect-[3/4] w-full max-w-[260px] rounded-2xl object-cover shadow-[0_20px_50px_-28px_rgba(0,0,0,0.85)] ring-1 ring-white/10 sm:max-w-[280px] lg:max-w-none"
            />
          </div>

          <div className="relative flex min-w-0 flex-1 flex-col justify-center">
            <div
              className="pointer-events-none absolute left-0 top-0 font-serif text-[4rem] leading-none tracking-tight text-slate-500/35 select-none sm:text-[4.5rem] lg:text-[5rem]"
              aria-hidden
            >
              “
            </div>

            <blockquote className="relative z-[1] border-none p-0 pt-6 sm:pt-7 md:pt-8">
              <p className="text-pretty text-lg font-normal leading-relaxed text-white sm:text-xl lg:text-2xl lg:leading-snug">
                {t.quote}
              </p>
            </blockquote>

            <footer className="relative z-[1] mt-8 sm:mt-9 lg:mt-10">
              <cite className="not-italic">
                <div className="text-base font-semibold text-white sm:text-lg">{t.name}</div>
                <div className="mt-1 text-sm text-slate-400 sm:text-base">{t.role}</div>
              </cite>
            </footer>
          </div>
        </div>

        <div
          className="mt-10 flex justify-center gap-2.5 sm:mt-11 lg:mt-12"
          role="tablist"
          aria-label="Choose testimonial"
        >
          {testimonials.map((_, index) => (
            <button
              key={index}
              type="button"
              role="tab"
              aria-selected={index === active}
              aria-label={`Testimonial ${index + 1}`}
              className={`rounded-full transition-[width,background-color] duration-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-400 ${
                index === active ? 'h-2.5 w-8 bg-white' : 'h-2.5 w-2.5 bg-slate-600 hover:bg-slate-500'
              }`}
              onClick={() => setActive(index)}
            />
          ))}
        </div>
      </div>
    </section>
  )
}
