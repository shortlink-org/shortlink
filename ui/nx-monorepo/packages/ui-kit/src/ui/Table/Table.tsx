import {
  MaterialReactTable,
  useMaterialReactTable,
  MRT_ToggleDensePaddingButton,
  MRT_ToggleFullScreenButton,
  MRT_ShowHideColumnsButton,
  MRT_ToggleFiltersButton,
  type MaterialReactTable as MaterialReactTableProps,
  type MRT_Cell,
  type MRT_ColumnDef,
  type MRT_Row,
} from 'material-react-table'
import React, { useState, useCallback } from 'react'
import { mkConfig, generateCsv, download } from 'export-to-csv'
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
  MenuItem,
  Stack,
  TextField,
  Tooltip,
} from '@mui/material'
import { Delete, Edit, FileDownload, Update } from '@mui/icons-material'
import CreateNewItemModal from './CreateNewItemModal/CreateNewItemModal'

const csvConfig = mkConfig({
  fieldSeparator: ',',
  decimalSeparator: '.',
  useKeysAsHeaders: true,
})

// @ts-ignore
export const Table = ({ columns, data }) => {
  // export data to csv --------------------------------------------------------
  const handleExportData = () => {
    const csv = generateCsv(csvConfig)(data)
    download(csvConfig)(csv)
  }

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

  const handleSaveRowEdits: MaterialReactTableProps<any>['onEditingRowSave'] = // @ts-ignore
    async ({ exitEditingMode, row, values }) => {
      if (!Object.keys(validationErrors).length) {
        tableData[row.index] = values
        //send/receive api updates here, then refetch or update local table data for re-render
        setTableData([...tableData])
        exitEditingMode() //required to exit editing mode and close modal
      }
    }

  const handleCancelRowEdits = () => {
    setValidationErrors({})
  }

  const handleDeleteRow = useCallback(
    (row: MRT_Row<any>) => {
      if (!confirm(`Are you sure you want to delete row ${row.index + 1}?`)) {
        return
      }
      //send api delete request here, then refetch or update local table data for re-render
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
    renderTopToolbarCustomActions: ({ table }) => (
      <Box sx={{ display: 'flex', gap: '1rem', p: '4px' }}>
        <Button
          color="primary"
          onClick={() => setCreateModalOpen(true)}
          variant="contained"
        >
          Create New Item
        </Button>
        <Button
          color="error"
          disabled={!table.getIsSomeRowsSelected() && !table.getIsAllRowsSelected()}
          onClick={() => {
            alert('Delete Selected Rows');
          }}
          variant="contained"
        >
          Delete Selected Rows
        </Button>
      </Box>
    ),
    renderToolbarInternalActions: ({ table }) => (
      <Box>
        <IconButton
          onClick={() => alert('Update')}
        >
          <Update />
        </IconButton>
        <MRT_ToggleFiltersButton table={table} />
        <MRT_ShowHideColumnsButton table={table} />
        <MRT_ToggleDensePaddingButton table={table} />
        <MRT_ToggleFullScreenButton table={table} />
      </Box>
    ),
    // @ts-ignore
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
          //export all data that is currently in the table (ignore pagination, sorting, filtering, etc.)
          onClick={handleExportData}
          startIcon={<FileDownload />}
        >
          Export All Data
        </Button>
        <Button
          disabled={table.getRowModel().rows.length === 0}
          //export all rows as seen on the screen (respects pagination, sorting, filtering, etc.)
          onClick={() => handleExportRows(table.getRowModel().rows)}
          startIcon={<FileDownload />}
        >
          Export Page Rows
        </Button>
        <Button
          disabled={
            !table.getIsSomeRowsSelected() && !table.getIsAllRowsSelected()
          }
          //only export selected rows
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
