import React from 'react'
import { DataTable, createDataTableColumnHelper } from '@shortlink-org/ui-kit'
import { formatRelative } from 'date-fns'
import { ContentCopy } from '@mui/icons-material'
import { LinkTableItem } from '@/components/Page/user/linksTable'

type AppProps = {
  data: LinkTableItem[]
}

const columnHelper = createDataTableColumnHelper<LinkTableItem>()

const columns = [
  columnHelper.accessor('url', {
    header: 'URL',
    size: 220,
    enableColumnFilter: true,
    cell: (info: { getValue: () => unknown }) => {
      const value = info.getValue()
      const text = typeof value === 'string' ? value : String(value ?? '')
      return (
        <div className="flex items-center gap-2">
          <span className="truncate">{text}</span>
          <button
            type="button"
            className="inline-flex items-center text-gray-500 hover:text-gray-700"
            aria-label="Copy URL"
            onClick={(event) => {
              event.stopPropagation()
              if (typeof navigator !== 'undefined' && navigator.clipboard) {
                void navigator.clipboard.writeText(text)
              }
            }}
          >
            <ContentCopy fontSize="inherit" />
          </button>
        </div>
      )
    },
  }),
  columnHelper.accessor('hash', {
    header: 'Hash',
    size: 150,
    enableColumnFilter: true,
  }),
  columnHelper.accessor('describe', {
    header: 'Describe',
    size: 200,
  }),
  columnHelper.accessor('created_at', {
    header: 'Created at',
    size: 180,
    enableColumnFilter: true,
    sortingFn: 'datetime',
    cell: (info: { getValue: () => unknown }) => {
      const dateValue = info.getValue()
      if (!dateValue) return ''
      try {
        return formatRelative(new Date(String(dateValue)), new Date())
      } catch {
        return String(dateValue)
      }
    },
  }),
  columnHelper.accessor('updated_at', {
    header: 'Updated at',
    size: 180,
    enableColumnFilter: true,
    sortingFn: 'datetime',
    cell: (info: { getValue: () => unknown }) => {
      const dateValue = info.getValue()
      if (!dateValue) return ''
      try {
        return formatRelative(new Date(String(dateValue)), new Date())
      } catch {
        return String(dateValue)
      }
    },
  }),
]

export const AdminUserLinksTable = ({ data }: AppProps) => (
  <DataTable
    data={data}
    columns={columns}
    filters={true}
    enableRefresh={true}
    onRefresh={() => {}}
  />
)

export default AdminUserLinksTable
