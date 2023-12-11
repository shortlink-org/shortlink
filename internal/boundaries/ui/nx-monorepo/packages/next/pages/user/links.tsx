import { NextSeo } from 'next-seo'
import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { Layout } from 'components'
import Statistic from 'components/Dashboard/stats'
import withAuthSync from 'components/Private'
import { fetchLinkList } from 'store'

import Header from '../../components/Page/Header'
import UserLinksTable from '../../components/Page/user/linksTable'

export function LinkTable() {
  // @ts-ignore
  const state = useSelector((rootState) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  return (
    <Layout>
      <NextSeo title="Links" description="Links page for your account." />

      <Header title="Links" />

      <Statistic count={state.list.length} />

      <UserLinksTable
        data={state.list}
        onRefresh={() => dispatch(fetchLinkList())}
      />
    </Layout>
  )
}

export default withAuthSync(LinkTable)
