import clsx from 'clsx'
import { LinkIcon } from '@heroicons/react/24/solid'

export function LogoSquare({ size }: { size?: 'sm' }) {
  return (
    <div
      className={clsx(
        'flex flex-none items-center justify-center border border-neutral-200 bg-white dark:border-neutral-700 dark:bg-neutral-950',
        {
          'h-10 w-10 rounded-xl': !size,
          'h-[30px] w-[30px] rounded-lg': size === 'sm',
        },
      )}
    >
      <LinkIcon
        className={clsx({
          'h-4 w-4 text-neutral-900 dark:text-neutral-100': !size,
          'h-2.5 w-2.5 text-neutral-900 dark:text-neutral-100': size === 'sm',
        })}
      />
    </div>
  )
}
