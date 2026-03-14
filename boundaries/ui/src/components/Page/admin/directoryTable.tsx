import React from 'react'
import { DataTable, createDataTableColumnHelper } from '@shortlink-org/ui-kit'

export type DirectoryTableItem = {
  id: string
  name: string
  email: string
  title: string
  department: string
  status: 'Active' | 'Invited' | 'Suspended'
  role: string
  avatar: string
}

type DirectoryTableProps = {
  data: DirectoryTableItem[]
}

type DirectoryCellContext<TValue> = {
  row: {
    original: DirectoryTableItem
  }
  getValue: () => TValue
}

const columnHelper = createDataTableColumnHelper<DirectoryTableItem>()

function getStatusClass(status: DirectoryTableItem['status']): string {
  if (status === 'Active') return 'bg-emerald-100 text-emerald-800 dark:bg-emerald-500/15 dark:text-emerald-300'
  if (status === 'Invited') return 'bg-sky-100 text-sky-800 dark:bg-sky-500/15 dark:text-sky-300'
  return 'bg-rose-100 text-rose-800 dark:bg-rose-500/15 dark:text-rose-300'
}

const columns = [
  columnHelper.accessor('name', {
    header: 'Member',
    size: 280,
    enableColumnFilter: true,
    cell: (info: DirectoryCellContext<DirectoryTableItem['name']>) => {
      const row = info.row.original

      return (
        <div className="flex items-center gap-3">
          <img src={row.avatar} alt={row.name} className="h-10 w-10 rounded-full object-cover" />
          <div className="min-w-0">
            <div className="truncate text-sm font-semibold text-gray-900 dark:text-gray-100">{row.name}</div>
            <div className="truncate text-sm text-gray-500 dark:text-gray-400">{row.email}</div>
          </div>
        </div>
      )
    },
  }),
  columnHelper.accessor('title', {
    header: 'Title',
    size: 240,
    enableColumnFilter: true,
    cell: (info: DirectoryCellContext<DirectoryTableItem['title']>) => {
      const row = info.row.original

      return (
        <div>
          <div className="text-sm font-medium text-gray-900 dark:text-gray-100">{row.title}</div>
          <div className="text-sm text-gray-500 dark:text-gray-400">{row.department}</div>
        </div>
      )
    },
  }),
  columnHelper.accessor('status', {
    header: 'Status',
    size: 140,
    enableColumnFilter: true,
    cell: (info: DirectoryCellContext<DirectoryTableItem['status']>) => {
      const value = info.getValue()
      return <span className={`inline-flex rounded-full px-2.5 py-1 text-xs font-semibold ${getStatusClass(value)}`}>{value}</span>
    },
  }),
  columnHelper.accessor('role', {
    header: 'Role',
    size: 140,
    enableColumnFilter: true,
  }),
]

export default function DirectoryTable({ data }: DirectoryTableProps) {
  return <DataTable data={data} columns={columns} filters={true} />
}
