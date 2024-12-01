interface StatisticProps {
  count: number
}

export default function Statistic({ count }: StatisticProps) {
  return (
    <div className="px-4 py-4 my-3 rounded mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-14 lg:px-8 lg:py-10 bg-white dark:bg-gray-800">
      <div className="flex flex-col lg:items-center lg:flex-row">
        <div className="flex items-center mb-6 lg:w-1/2 lg:mb-0">
          <div className="flex items-center justify-center w-16 h-16 mr-5 rounded-full bg-indigo-50 dark:bg-indigo-900 sm:w-24 sm:h-24 xl:mr-10 xl:w-28 xl:h-28">
            <svg
              className="w-12 h-12 text-deep-purple-accent-400 dark:text-deep-purple-accent-200 sm:w-16 sm:h-16 xl:w-20 xl:h-20"
              stroke="currentColor"
              viewBox="0 0 52 52"
            >
              <polygon
                strokeWidth="3"
                strokeLinecap="round"
                strokeLinejoin="round"
                fill="none"
                points="29 13 14 29 25 29 23 39 38 23 27 23"
              />
            </svg>
          </div>
          <h3 className="text-4xl font-extrabold text-gray-800 dark:text-white sm:text-5xl xl:text-6xl">{count}</h3>
        </div>
        <div className="lg:w-1/2">
          <p className="text-gray-800 dark:text-gray-300">
            A table can be a great way to keep track of links. You can add links to the table, delete links from the table, and update links
            in the table. This can be a great way to organize your links and keep track of them. You can also use the table to share links
            with others. This can be a great way to share links with friends or family.
          </p>
        </div>
      </div>
    </div>
  )
}
