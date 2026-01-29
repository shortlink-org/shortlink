'use client'

function LinkIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={2}
        d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
      />
    </svg>
  )
}

interface StatisticProps {
  count: number
}

export default function Statistic({ count }: StatisticProps) {
  const getMessage = () => {
    if (count === 0) {
      return 'You don\'t have any saved links yet. Create your first link using the "Add URL" button in the menu.'
    }
    if (count === 1) {
      return 'You have one saved link. You can manage it through the table: view, edit, or delete it.'
    }
    return `You have ${count} saved links. Use the table to manage your links: filter, sort, edit, and delete them.`
  }

  return (
    <div className="my-6 px-4">
      <div className="rounded-2xl bg-gradient-to-br from-indigo-500 to-purple-600 dark:from-indigo-600 dark:to-purple-700 p-6 shadow-lg">
        <div className="flex flex-col lg:flex-row lg:items-center gap-6">
          {/* Icon and Count */}
          <div className="flex items-center gap-4">
            <div className="flex items-center justify-center w-16 h-16 rounded-2xl bg-white/20 backdrop-blur-sm">
              <LinkIcon className="w-8 h-8 text-white" />
            </div>
            <div>
              <p className="text-5xl lg:text-6xl font-bold text-white">{count}</p>
              <p className="text-sm text-indigo-200 font-medium">{count === 1 ? 'Link' : 'Links'}</p>
            </div>
          </div>

          {/* Message */}
          <div className="lg:flex-1 lg:border-l lg:border-white/20 lg:pl-6">
            <p className="text-white/90 leading-relaxed">{getMessage()}</p>
          </div>
        </div>
      </div>
    </div>
  )
}
