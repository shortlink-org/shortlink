type TimelineItemProps = {
  name: string
  date: string
  action: string
  content: string
  icon: React.ReactNode
}

const TimelineItem: React.FC<TimelineItemProps> = ({
  name,
  date,
  action,
  content,
  icon,
}) => (
  <div className="relative">
    <div className="md:flex items-center md:space-x-4 mb-3">
      <div className="flex items-center space-x-4 md:space-x-2 md:space-x-reverse">
        {/* Icon */}
        <div className="flex items-center justify-center w-10 h-10 rounded-full bg-white dark:bg-gray-800 shadow md:order-1">
          {icon}
        </div>
        {/* Date */}
        <time className="font-caveat font-medium text-xl text-indigo-500 dark:text-indigo-300 md:w-28">
          {date}
        </time>
      </div>
      {/* Title */}
      <div className="text-slate-500 ml-14 dark:text-gray-300">
        <span className="text-slate-900 font-bold dark:text-gray-100">
          {name}
        </span>{' '}
        {action}
      </div>
    </div>
    {/* Card */}
    <div className="bg-white dark:bg-gray-800 p-4 rounded border border-slate-200 dark:border-gray-700 text-slate-500 dark:text-gray-300 shadow ml-14 md:ml-44">
      {content}
    </div>
  </div>
)

export type TimelineProps = {
  items: TimelineItemProps[]
}

export const Timeline: React.FC<TimelineProps> = ({ items }) => (
  <div className="space-y-8 relative before:absolute before:inset-0 before:ml-5 before:-translate-x-px md:before:ml-[8.75rem] md:before:translate-x-0 before:h-full before:w-0.5 before:bg-gradient-to-b before:from-transparent before:via-slate-300 before:to-transparent">
    {items.map((item, index) => (
      <TimelineItem
        key={index}
        date={item.date}
        name={item.name}
        action={item.action}
        content={item.content}
        icon={item.icon}
      />
    ))}
  </div>
)

export default Timeline
