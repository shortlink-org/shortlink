import * as React from 'react'
import ActiveLink from './ActiveLink'

type AppProps = {
  mode: 'full' | 'mini'
  url: string
  icon: JSX.Element
  name: string
}

function getItem({ mode, url, icon, name }: AppProps) {
  const iconClassName =
    'text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
  let linkClassName =
    'flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group'

  if (mode === 'mini') {
    linkClassName += ' item-center justify-center'
  }

  return (
    <li key={url} className={'w-full'}>
      <ActiveLink href={url} passHref activeClassName="md:text-blue-700">
        <div className={linkClassName}>
          {React.cloneElement(icon, { className: iconClassName })}

          {mode === 'full' && <span className="ms-3">{name}</span>}
        </div>
      </ActiveLink>
    </li>
  )
}

export default getItem
