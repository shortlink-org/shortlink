'use client'

import { MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import { useRouter, useSearchParams } from 'next/navigation'
import { Suspense, useCallback } from 'react'

function NavbarSearchForm() {
  const router = useRouter()
  const searchParams = useSearchParams()

  const onSubmit = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault()
      const form = e.target as HTMLFormElement
      const input = form.search as HTMLInputElement
      const q = input.value.trim()
      const params = new URLSearchParams()
      if (q) params.set('q', q)
      const qs = params.toString()
      router.push(qs ? `/links/search?${qs}` : '/links/search')
    },
    [router],
  )

  return (
    <form onSubmit={onSubmit} className="relative w-full max-w-[550px] lg:w-80 xl:w-full">
      <input
        key={searchParams?.get('q') ?? ''}
        type="text"
        name="search"
        placeholder="Search links..."
        autoComplete="off"
        defaultValue={searchParams?.get('q') ?? ''}
        className="text-md w-full rounded-lg border border-neutral-200 bg-white px-4 py-2 pr-10 text-black placeholder:text-neutral-500 md:text-sm dark:border-neutral-800 dark:bg-transparent dark:text-white dark:placeholder:text-neutral-400"
      />
      <div className="pointer-events-none absolute right-0 top-0 mr-3 flex h-full items-center">
        <MagnifyingGlassIcon className="h-4 w-4 text-neutral-500 dark:text-neutral-400" aria-hidden />
      </div>
    </form>
  )
}

export function NavbarSearchSkeleton() {
  return (
    <div className="relative w-full max-w-[550px] lg:w-80 xl:w-full">
      <div className="w-full rounded-lg border border-neutral-200 bg-white px-4 py-2 pr-10 text-sm text-transparent dark:border-neutral-800 dark:bg-transparent">
        Search links...
      </div>
      <div className="pointer-events-none absolute right-0 top-0 mr-3 flex h-full items-center">
        <MagnifyingGlassIcon className="h-4 w-4 text-neutral-500 dark:text-neutral-400" aria-hidden />
      </div>
    </div>
  )
}

export default function NavbarSearch() {
  return (
    <Suspense fallback={<NavbarSearchSkeleton />}>
      <NavbarSearchForm />
    </Suspense>
  )
}
