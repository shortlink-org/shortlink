import React, { useEffect } from 'react'
import { useSelector } from 'react-redux'
import Router from 'next/router'

export default function withAuthSync(Child: any) {
  return (props?: any) => {
    // checks whether we are on client / browser or server.
    if (typeof window !== 'undefined') {
      // @ts-ignore
      const state = useSelector(state => state.session)

      if (!state.token) {
        Router.push('/auth/login')
        return null
      }

      // If this is an token we just render the component that was passed with all its props
      return <Child {...props} />
    }

    // If we are on server, return null
    return null
  }
}
