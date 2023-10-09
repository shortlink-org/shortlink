// import MaterialTable, { Icons } from '@material-table/core'
// import AddBox from '@mui/icons-material/AddBox'
// import ArrowDownward from '@mui/icons-material/ArrowDownward'
// import Check from '@mui/icons-material/Check'
// import ChevronLeft from '@mui/icons-material/ChevronLeft'
// import ChevronRight from '@mui/icons-material/ChevronRight'
// import Clear from '@mui/icons-material/Clear'
// import DeleteOutline from '@mui/icons-material/DeleteOutline'
// import Edit from '@mui/icons-material/Edit'
// import FilterList from '@mui/icons-material/FilterList'
// import FirstPage from '@mui/icons-material/FirstPage'
// import LastPage from '@mui/icons-material/LastPage'
// import Remove from '@mui/icons-material/Remove'
// import Save from '@mui/icons-material/Save'
// import SaveAlt from '@mui/icons-material/SaveAlt'
// import Search from '@mui/icons-material/Search'
// import Update from '@mui/icons-material/Update'
// import ViewColumn from '@mui/icons-material/ViewColumn'
// import Link from '@mui/material/Link'
// import Tooltip from '@mui/material/Tooltip'
// import { formatRelative } from 'date-fns'
// import React, { forwardRef } from 'react'
//
// // specified type as Icons
// const tableIcons: Icons<any> = {
//   Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
//   Update: forwardRef((props, ref) => <Update {...props} ref={ref} />),
//   Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
//   Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
//   Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
//   DetailPanel: forwardRef((props, ref) => (
//     <ChevronRight {...props} ref={ref} />
//   )),
//   Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
//   Save: forwardRef((props, ref) => <Save {...props} ref={ref} />),
//   Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
//   Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
//   FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
//   LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
//   NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
//   PreviousPage: forwardRef((props, ref) => (
//     <ChevronLeft {...props} ref={ref} />
//   )),
//   ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
//   Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
//   SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
//   ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
//   ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />),
// }
//
// function Table(props) {
//   return (
//     <MaterialTable
//       icons={tableIcons}
//       columns={[
//         {
//           title: 'URL',
//           field: 'url',
//           render: (rowData) => (
//             <Link
//               href={rowData.url}
//               target="_blank"
//               rel="noopener"
//               variant="p"
//               underline="hover"
//             >
//               {rowData.url}
//             </Link>
//           ),
//         },
//         { title: 'hash', field: 'hash' },
//         { title: 'Describe', field: 'describe' },
//         {
//           title: 'Created at',
//           field: 'createdAt',
//           render: (rowData) => (
//             <Tooltip arrow title={rowData.createdAt} interactive>
//                 <span>
//                   {formatRelative(new Date(rowData.createdAt), new Date(), {
//                     addSuffix: true,
//                   })}
//                 </span>
//             </Tooltip>
//           ),
//         },
//         {
//           title: 'Updated at',
//           field: 'updatedAt',
//           render: (rowData) => (
//             <Tooltip arrow title={rowData.updatedAt} interactive>
//                 <span>
//                   {formatRelative(new Date(rowData.updatedAt), new Date(), {
//                     addSuffix: true,
//                   })}
//                 </span>
//             </Tooltip>
//           ),
//         },
//       ]}
//       data={props.list || []}
//       actions={[
//         {
//           icon: tableIcons.Add,
//           tooltip: 'Add link',
//           isFreeAction: true,
//           onClick: () => alert('You want to add a new row'),
//         },
//         {
//           icon: tableIcons.Update,
//           tooltip: 'Update link',
//           isFreeAction: true,
//           // onClick: () => dispatch(fetchLinkList()),
//         },
//         {
//           icon: tableIcons.Save,
//           tooltip: 'Save link',
//           onClick: (event, rowData) => alert(`You saved ${rowData.name}`),
//         },
//         {
//           icon: tableIcons.Delete,
//           tooltip: 'Delete link',
//           onClick: (event, rowData) =>
//             confirm(`You want to delete ${rowData.name}`),
//         },
//       ]}
//       options={{
//         exportButton: true,
//       }}
//       title="Link Table"
//     />
//   )
// }
//
// export default Table
