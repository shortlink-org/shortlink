import { Delete, Edit, FileDownload, Update } from '@mui/icons-material'
import { Box, Button, IconButton, Tooltip } from '@mui/material'
import { mkConfig, generateCsv, download } from 'export-to-csv'
import {
  MaterialReactTable,
  useMaterialReactTable,
  // eslint-disable-next-line camelcase
  MRT_ToggleDensePaddingButton,
  // eslint-disable-next-line camelcase
  MRT_ToggleFullScreenButton,
  // eslint-disable-next-line camelcase
  MRT_ShowHideColumnsButton,
  // eslint-disable-next-line camelcase
  MRT_ToggleFiltersButton,
  type MaterialReactTable as MaterialReactTableProps,
  // eslint-disable-next-line camelcase
  type MRT_Row,
} from 'material-react-table'
import React, { useState, useCallback } from 'react'

import CreateNewItemModal from './CreateNewItemModal/CreateNewItemModal'

const csvConfig = mkConfig({
  fieldSeparator: ',',
  decimalSeparator: '.',
  useKeysAsHeaders: true,
})

type TableProps = {
  columns: any
  data: any

  onCreate?: () => void
  onUpdate?: () => void
  onDelete?: () => void
  onRefresh?: () => void
}

export const Table: React.FC<TableProps> = ({ columns, data, onRefresh }) => {
  // export data to csv --------------------------------------------------------
  const handleExportData = () => {
    const csv = generateCsv(csvConfig)(data)
    download(csvConfig)(csv)
  }

  // eslint-disable-next-line camelcase
  const handleExportRows = (rows: MRT_Row<any>[]) => {
    const rowData = rows.map((row) => row.original)
    const csv = generateCsv(csvConfig)(rowData)
    download(csvConfig)(csv)
  }

  // CRUD ----------------------------------------------------------------------
  const [createModalOpen, setCreateModalOpen] = useState(false)
  const [tableData, setTableData] = useState<any[]>(() => data)
  const [validationErrors, setValidationErrors] = useState<{
    [cellId: string]: string
  }>({})

  const handleCreateNewRow = (values: any) => {
    tableData.push(values)
    setTableData([...tableData])
  }

  // @ts-ignore
  const handleSaveRowEdits: MaterialReactTableProps<any>['onEditingRowSave'] = // @ts-ignore
    async ({ exitEditingMode, row, values }) => {
      if (!Object.keys(validationErrors).length) {
        tableData[row.index] = values
        // send/receive api updates here, then refetch or update local table data for re-render
        setTableData([...tableData])
        exitEditingMode() // required to exit editing mode and close modal
      }
    }

  const handleCancelRowEdits = () => {
    setValidationErrors({})
  }

  const handleDeleteRow = useCallback(
    // eslint-disable-next-line camelcase
    (row: MRT_Row<any>) => {
      if (!confirm(`Are you sure you want to delete row ${row.index + 1}?`)) {
        return
      }
      // send api delete request here, then refetch or update local table data for re-render
      tableData.splice(row.index, 1)
      setTableData([...tableData])
    },
    [tableData],
  )

  // @ts-ignore
  const table = useMaterialReactTable({
    columns,
    data,
    initialState: { showColumnFilters: false, showGlobalFilter: true },
    enableColumnFilterModes: true,
    enableRowSelection: true,
    enableColumnOrdering: true,
    enableGlobalFilter: true,
    enableGrouping: true,
    enableColumnPinning: true,
    enableFacetedValues: true,
    enableRowActions: true,
    paginationDisplayMode: 'pages',

    // eslint-disable-next-line no-shadow
    renderTopToolbarCustomActions: ({ table }) => (
      <Box sx={{ display: 'flex', gap: '1rem', p: '4px' }}>
        <Button onClick={() => setCreateModalOpen(true)} variant="outlined">
          Create New Item
        </Button>
        <Button
          color="error"
          disabled={
            !table.getIsSomeRowsSelected() && !table.getIsAllRowsSelected()
          }
          onClick={() => {
            alert('Delete Selected Rows')
          }}
          variant="contained"
        >
          Delete Selected Rows
        </Button>
      </Box>
    ),

    // eslint-disable-next-line no-shadow
    renderToolbarInternalActions: ({ table }) => (
      <Box>
        <IconButton onClick={onRefresh}>
          <Update />
        </IconButton>
        {/* eslint-disable-next-line react/jsx-pascal-case, camelcase */}
        <MRT_ToggleFiltersButton table={table} />
        {/* eslint-disable-next-line react/jsx-pascal-case, camelcase */}
        <MRT_ShowHideColumnsButton table={table} />
        {/* eslint-disable-next-line react/jsx-pascal-case, camelcase */}
        <MRT_ToggleDensePaddingButton table={table} />
        {/* eslint-disable-next-line react/jsx-pascal-case, camelcase */}
        <MRT_ToggleFullScreenButton table={table} />
      </Box>
    ),

    // eslint-disable-next-line no-shadow
    renderRowActions: ({ row, table }) => (
      <Box sx={{ display: 'flex', gap: '1rem' }}>
        <Tooltip arrow placement="left" title="Edit">
          <IconButton onClick={() => table.setEditingRow(row)}>
            <Edit />
          </IconButton>
        </Tooltip>
        <Tooltip arrow placement="right" title="Delete">
          <IconButton color="error" onClick={() => handleDeleteRow(row)}>
            <Delete />
          </IconButton>
        </Tooltip>
      </Box>
    ),

    // eslint-disable-next-line no-shadow
    renderBottomToolbarCustomActions: ({ table }) => (
      <Box
        sx={{
          display: 'flex',
          gap: '16px',
          padding: '8px',
          flexWrap: 'wrap',
        }}
      >
        <Button
          // export all data that is currently in the table (ignore pagination, sorting, filtering, etc.)
          onClick={handleExportData}
          startIcon={<FileDownload />}
        >
          Export All Data
        </Button>
        <Button
          disabled={table.getRowModel().rows.length === 0}
          // export all rows as seen on the screen (respects pagination, sorting, filtering, etc.)
          onClick={() => handleExportRows(table.getRowModel().rows)}
          startIcon={<FileDownload />}
        >
          Export Page Rows
        </Button>
        <Button
          disabled={
            !table.getIsSomeRowsSelected() && !table.getIsAllRowsSelected()
          }
          // only export selected rows
          onClick={() => handleExportRows(table.getSelectedRowModel().rows)}
          startIcon={<FileDownload />}
        >
          Export Selected Rows
        </Button>
      </Box>
    ),
  })

  return (
    <>
      {/* @ts-ignore */}
      <MaterialReactTable
        table={table}
        onEditingRowSave={handleSaveRowEdits}
        onEditingRowCancel={handleCancelRowEdits}
      />
      <CreateNewItemModal
        columns={columns}
        open={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        onSubmit={handleCreateNewRow}
      />
    </>
  )
}

export default Table
