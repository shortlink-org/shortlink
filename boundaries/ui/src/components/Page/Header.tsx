type AppProps = {
  title: string
  description?: string
}

function Header({ title, description }: AppProps) {
  return (
    <div className="my-6 lg:my-10 px-4 lg:px-6">
      <div className="pb-4 border-b border-gray-200 dark:border-gray-700">
        <h1 className="text-2xl lg:text-3xl font-bold text-gray-900 dark:text-white">{title}</h1>
        {description && <p className="mt-2 text-sm text-gray-600 dark:text-gray-400">{description}</p>}
      </div>
    </div>
  )
}

export default Header
