import React from 'react'
// @ts-ignore
import { Table } from '@shortlink-org/ui-kit'

type AppProps = {
  data: any
}

const columns = [
  {
    accessorKey: 'url',
    header: 'URL',
    size: 150,
  },
  {
    accessorKey: 'hash',
    header: 'Hash',
    size: 150,
  },
  {
    accessorKey: 'describe',
    header: 'Describe',
    size: 150,
  },
  {
    accessorKey: 'createdAt',
    header: 'Created at',
    size: 150,
  },
  {
    accessorKey: 'updatedAt',
    header: 'Updated at',
    size: 150,
  },
]

export const AdminUserLinksTable = ({ data }: AppProps) => (
  <Table data={data} columns={columns} />
)

export default AdminUserLinksTable
