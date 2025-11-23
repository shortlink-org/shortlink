'use client'

import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { LoadingSpinner, ErrorAlert } from '@/components/common'

import Statistic from '@/components/Dashboard/stats'
import withAuthSync from '@/components/Private'
import { fetchLinkList } from '@/store'
import AdminUserLinksTable from '@/components/Page/admin/user/linksTable'
import Header from '@/components/Page/Header'
import { LinkState } from '@/store/reducers/link'

function LinkTable() {
  const state = useSelector((rootState: { link: LinkState }) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  // Преобразуем данные для таблицы (конвертируем TimestamppbTimestamp в строки)
  const tableData = state.list.map((link) => ({
    ...link,
    created_at: link.created_at
      ? new Date((link.created_at.seconds || 0) * 1000 + (link.created_at.nanos || 0) / 1000000).toISOString()
      : '',
    updated_at: link.updated_at
      ? new Date((link.updated_at.seconds || 0) * 1000 + (link.updated_at.nanos || 0) / 1000000).toISOString()
      : '',
  }))

  return (
    <>
      {/*<NextSeo title="Links" description="Admin links page" />*/}

      <Header title="Admin links" />

      {state.loading && <LoadingSpinner minHeight="200px" />}

      <ErrorAlert error={state.error} />

      {!state.loading && (
        <>
          <Statistic count={state.list.length} />

          <AdminUserLinksTable data={tableData} />
        </>
      )}
    </>
  )
}

export default withAuthSync(LinkTable)
