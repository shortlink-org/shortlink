import { MaterialReactTable } from 'material-react-table'
import React from 'react'

// @ts-ignore
export const Table = ({ columns, data }) => (
  <MaterialReactTable columns={columns} data={data} />
)

export default Table
