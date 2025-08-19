import React from 'react'
// @ts-ignore
import { Table } from '@shortlink-org/ui-kit'
import { formatRelative } from 'date-fns'
import { ContentCopy } from '@mui/icons-material'

type AppProps = {
  data: Array<{
    url: string
    hash: string
    describe?: string
    created_at: string
    updated_at: string
  }>

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
    Cell: ({ cell }: CellProps) =>
      formatRelative(new Date(cell.getValue()), new Date(), {
        // @ts-ignore
        addSuffix: true,
      }),
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
    Cell: ({ cell }: CellProps) =>
      formatRelative(new Date(cell.getValue()), new Date(), {
        // @ts-ignore
        addSuffix: true,
      }),
    muiFilterTextFieldProps: {
      sx: {
        minWidth: '250px',
      },
    },
  },
]

export const UserLinksTable = ({ data, onRefresh }: AppProps) => <Table data={data} columns={columns} onRefresh={onRefresh} />

export default UserLinksTable
