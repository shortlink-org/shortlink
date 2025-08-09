'use client'

import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import Statistic from 'components/Dashboard/stats'
import withAuthSync from 'components/Private'
import { fetchLinkList } from 'store'
import Header from 'components/Page/Header'
import UserLinksTable from 'components/Page/user/linksTable'

function LinkTable() {
  // @ts-ignore
  const state = useSelector((rootState) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  return (
    <>
      {/*<NextSeo title="Links" description="Links page for your account." />*/}

      <Header title="Links" />

      <Statistic count={state.list.length} />

      <UserLinksTable data={state.list} onRefresh={() => dispatch(fetchLinkList())} />
    </>
  )
}

export default withAuthSync(LinkTable)
