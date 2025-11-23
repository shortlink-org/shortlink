import React from 'react'
// @ts-ignore
import { Table } from '@shortlink-org/ui-kit'
import { formatRelative } from 'date-fns'
import { ContentCopy } from '@mui/icons-material'

export interface LinkTableItem {
  url: string
  hash: string
  describe?: string
  created_at: string
  updated_at: string
}

type AppProps = {
  data: LinkTableItem[]
  onRefresh?: () => void
}

type CellProps = {
  cell: {
    getValue: () => string
  }
}

const columns = [
  {
    accessorKey: 'url',
    header: 'URL',
    size: 150,
    enableClickToCopy: true,
    muiCopyButtonProps: {
      fullWidth: true,
      startIcon: <ContentCopy />,
      sx: { justifyContent: 'flex-start' },
    },
    filterVariant: 'autocomplete',
    enableEditing: false,
  },
  {
    accessorKey: 'hash',
    header: 'Hash',
    size: 150,
    filterVariant: 'autocomplete',
    enableEditing: false,
  },
  {
    accessorKey: 'describe',
    header: 'Describe',
    size: 150,
  },
  {
    accessorKey: 'created_at',
    header: 'Created at',
    size: 150,
    filterVariant: 'date',
    filterFn: 'lessThan',
    sortingFn: 'datetime',
    Cell: ({ cell }: CellProps) => {
      const dateValue = cell.getValue()
      if (!dateValue) return ''
      try {
        return formatRelative(new Date(dateValue), new Date(), {
          addSuffix: true,
        })
      } catch {
        return dateValue
      }
    },
    muiFilterTextFieldProps: {
      sx: {
        minWidth: '250px',
      },
    },
  },
  {
    accessorKey: 'updated_at',
    header: 'Updated at',
    size: 150,
    filterVariant: 'date',
    filterFn: 'lessThan',
    sortingFn: 'datetime',
    Cell: ({ cell }: CellProps) => {
      const dateValue = cell.getValue()
      if (!dateValue) return ''
      try {
        return formatRelative(new Date(dateValue), new Date(), {
          addSuffix: true,
        })
      } catch {
        return dateValue
      }
    },
    muiFilterTextFieldProps: {
      sx: {
        minWidth: '250px',
      },
    },
  },
]

export const UserLinksTable = ({ data, onRefresh }: AppProps) => <Table data={data} columns={columns} onRefresh={onRefresh} />

export default UserLinksTable
