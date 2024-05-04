import * as React from 'react'
import Link from 'next/link'

const ProfileLink = () => {
  return (
    <Link href="/user/profile">
      <button
        aria-label="visit"
        className="cursor-pointer focus:ring-2 focus:outline-none hover:bg-indigo-700 p-2 bg-indigo-600 rounded-full"
      >
        <svg
          width={20}
          height={20}
          viewBox="0 0 20 20"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M4.16666 10H15.8333"
            stroke="white"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M10.8333 15L15.8333 10"
            stroke="white"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M10.8333 5L15.8333 10"
            stroke="white"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </button>
    </Link>
  )
}

// @ts-ignore
function Footer({ mode }) {
  return (
    <div className="flex bg-indigo-500 items-center justify-center space-x-2 py-4 px-3 w-full h-16 mt-auto">
      {mode === 'full' ? (
        <React.Fragment>
          <div className={'m-1'}>
            <img src="https://i.ibb.co/fxrbS6p/Ellipse-2-2.png" alt="avatar" />
          </div>

          <div className="flex flex-col justify-start items-start space-y-2 m-1">
            <p className="cursor-pointer text-base leading-4 text-white">
              Alexis Enache
            </p>
            <p className="cursor-pointer text-xs leading-3 text-gray-200">
              alexis _enache@gmail.com
            </p>
          </div>

          <ProfileLink />
        </React.Fragment>
      ) : (
        <ProfileLink />
      )}
    </div>
  )
}

export default Footer
