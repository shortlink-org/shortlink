import Balancer from 'react-wrap-balancer'

type AppProps = {
  title: string
}

export function Header({ title }: AppProps) {
  return (
    <div className="my-6 lg:my-10 container px-6 mx-auto flex flex-col md:flex-row items-center justify-between pb-4 border-b border-gray-300 dark:border-gray-700 transition-colors duration-500">
      <div>
        <h4 className="text-2xl md:text-3xl font-bold leading-tight dark:text-white transition-colors duration-500">
          <Balancer>{title}</Balancer>
        </h4>
      </div>
      <div className="mt-6 md:mt-0">
        <button
          type="button"
          className="mr-3 bg-gray-200 dark:bg-gray-700 focus:outline-none transition duration-150 ease-in-out rounded hover:bg-gray-300 dark:hover:bg-gray-600 text-indigo-700 dark:text-indigo-300 px-5 py-2 text-sm transition-colors duration-200"
        >
          Back
        </button>
        <button
          type="button"
          className="transition-colors focus:outline-none duration-150 ease-in-out hover:bg-indigo-500 bg-indigo-700 dark:bg-indigo-800 dark:hover:bg-indigo-600 rounded text-white px-8 py-2 text-sm"
        >
          Edit Profile
        </button>
      </div>
    </div>
  )
}

export default Header
