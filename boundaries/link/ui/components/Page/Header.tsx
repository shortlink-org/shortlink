type AppProps = {
  title: string
}

function Header({ title }: AppProps) {
  return (
    <div className="my-6 lg:my-12 container px-6 mx-auto dark:bg-gray-400 flex flex-col md:flex-row items-start md:items-center justify-between pb-4 border-b border-gray-300">
      <div>
        <h4 className="text-2xl font-bold leading-tight text-gray-800 dark:text-white">{title}</h4>
      </div>
    </div>
  )
}

export default Header
