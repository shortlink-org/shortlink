'use client'

import { useTransition, ReactNode, MouseEvent } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

interface TransitionLinkProps {
  href: string
  children: ReactNode
  className?: string
  /** Show pending state on the link itself */
  showPending?: boolean
  /** Prefetch the page on hover */
  prefetch?: boolean
  /** Callback when transition starts */
  onTransitionStart?: () => void
  /** Callback when transition ends */
  onTransitionEnd?: () => void
  [key: string]: any
}

/**
 * Link с useTransition для плавной навигации
 * 
 * Старый контент остается видимым пока загружается новая страница
 * 
 * @example
 * ```tsx
 * <TransitionLink href="/profile">
 *   Go to Profile
 * </TransitionLink>
 * ```
 */
export function TransitionLink({
  href,
  children,
  className = '',
  showPending = true,
  prefetch = true,
  onTransitionStart,
  onTransitionEnd,
  ...props
}: TransitionLinkProps) {
  const [isPending, startTransition] = useTransition()
  const router = useRouter()

  const handleClick = (e: MouseEvent<HTMLAnchorElement>) => {
    // Если есть модификаторы (Ctrl, Cmd, Shift) - используем стандартное поведение
    if (e.metaKey || e.ctrlKey || e.shiftKey) {
      return
    }

    e.preventDefault()
    
    onTransitionStart?.()
    
    startTransition(() => {
      router.push(href)
      // В next.js router.push возвращает void, но transition ждёт пока страница загрузится
    })
    
    // onTransitionEnd вызовется когда transition закончится
    if (onTransitionEnd) {
      setTimeout(onTransitionEnd, 0)
    }
  }

  const handleMouseEnter = () => {
    // Prefetch on hover for instant navigation
    if (prefetch) {
      router.prefetch(href)
    }
  }

  return (
    <Link
      href={href}
      onClick={handleClick}
      onMouseEnter={handleMouseEnter}
      className={`${className} ${showPending && isPending ? 'opacity-70 transition-opacity' : ''}`}
      {...props}
    >
      {children}
    </Link>
  )
}

export default TransitionLink

