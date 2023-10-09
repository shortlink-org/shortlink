import { NextSeo } from 'next-seo'
import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { Layout } from 'components'
import Statistic from 'components/Dashboard/stats'
import withAuthSync from 'components/Private'
import { fetchLinkList } from 'store'

import AdminUserLinksTable from '../../components/Page/admin/user/linksTable'
import Header from '../../components/Page/Header'

// const columns = [
//   {
//     title: 'URL',
//     field: 'url',
//     render: (rowData) => (
//       <Link
//         href={rowData.url}
//         target="_blank"
//         rel="noopener"
//         variant="p"
//         underline="hover"
//       >
//         {rowData.url}
//       </Link>
//     ),
//   },
//   { title: 'hash', field: 'hash' },
//   { title: 'Describe', field: 'describe' },
//   {
//     title: 'Created at',
//     field: 'createdAt',
//     render: (rowData) => (
//       <Tooltip arrow title={rowData.createdAt} interactive>
//         <span>
//           {formatRelative(new Date(rowData.createdAt), new Date(), {
//             addSuffix: true,
//           })}
//         </span>
//       </Tooltip>
//     ),
//   },
//   {
//     title: 'Updated at',
//     field: 'updatedAt',
//     render: (rowData) => (
//       <Tooltip arrow title={rowData.updatedAt} interactive>
//         <span>
//           {formatRelative(new Date(rowData.updatedAt), new Date(), {
//             addSuffix: true,
//           })}
//         </span>
//       </Tooltip>
//     ),
//   },
// ]

export function LinkTable() {
  // @ts-ignore
  const state = useSelector((rootState) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  return (
    <Layout>
      <NextSeo title="Links" description="Admin links page" />

      <Header title="Admin links" />

      <Statistic count={state.list.length} />

      <AdminUserLinksTable data={state.list || []} />
    </Layout>
  )
}

// <MaterialTable
//   icons={tableIcons}
//   columns={columns}
//   data={state.list || []}
//   actions={[
//     {
//       icon: tableIcons.Add,
//       tooltip: 'Add link',
//       isFreeAction: true,
//       onClick: () => alert('You want to add a new row'),
//     },
//     {
//       icon: tableIcons.Update,
//       tooltip: 'Update link',
//       isFreeAction: true,
//       onClick: () => dispatch(fetchLinkList()),
//     },
//     {
//       icon: tableIcons.Save,
//       tooltip: 'Save link',
//       onClick: (event, rowData) => alert(`You saved ${rowData.name}`),
//     },
//     {
//       icon: tableIcons.Delete,
//       tooltip: 'Delete link',
//       onClick: (event, rowData) =>
//         confirm(`You want to delete ${rowData.name}`),
//     },
//   ]}
//   options={{
//     exportButton: true,
//   }}
//   title="Link Table"
// />

export default withAuthSync(LinkTable)
