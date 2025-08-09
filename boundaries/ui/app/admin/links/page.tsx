'use client'

import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import Statistic from 'components/Dashboard/stats'
import withAuthSync from 'components/Private'
import { fetchLinkList } from 'store'
import AdminUserLinksTable from 'components/Page/admin/user/linksTable'
import Header from 'components/Page/Header'

function LinkTable() {
  // @ts-ignore
  const state = useSelector((rootState) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  return (
    <>
      {/*<NextSeo title="Links" description="Admin links page" />*/}

      <Header title="Admin links" />

      <Statistic count={state.list.length} />

      <AdminUserLinksTable data={state.list || []} />
    </>
  )
}

export default withAuthSync(LinkTable)
