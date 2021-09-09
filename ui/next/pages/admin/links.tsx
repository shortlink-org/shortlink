// @ts-nocheck
import React, { forwardRef, useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import MaterialTable from 'material-table'
import Tooltip from '@material-ui/core/Tooltip'
import { formatRelative } from 'date-fns'
import AddBox from '@material-ui/icons/AddBox'
import Update from '@material-ui/icons/Update'
import ArrowDownward from '@material-ui/icons/ArrowDownward'
import Check from '@material-ui/icons/Check'
import ChevronLeft from '@material-ui/icons/ChevronLeft'
import ChevronRight from '@material-ui/icons/ChevronRight'
import Clear from '@material-ui/icons/Clear'
import DeleteOutline from '@material-ui/icons/DeleteOutline'
import Edit from '@material-ui/icons/Edit'
import FilterList from '@material-ui/icons/FilterList'
import FirstPage from '@material-ui/icons/FirstPage'
import LastPage from '@material-ui/icons/LastPage'
import Remove from '@material-ui/icons/Remove'
import SaveAlt from '@material-ui/icons/SaveAlt'
import Save from '@material-ui/icons/Save'
import Search from '@material-ui/icons/Search'
import ViewColumn from '@material-ui/icons/ViewColumn'
import Link from '@material-ui/core/Link'
import { Layout } from 'components'
import Statistic from 'components/Dashboard/stats'
import { fetchLinkList } from 'store'
import withAuthSync from 'components/Private'

// @ts-ignore
const tableIcons = {
  Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
  Update: forwardRef((props, ref) => <Update {...props} ref={ref} />),
  Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
  Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
  DetailPanel: forwardRef((props, ref) => (
    <ChevronRight {...props} ref={ref} />
  )),
  Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
  Save: forwardRef((props, ref) => <Save {...props} ref={ref} />),
  Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
  Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
  FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
  LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
  NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
  PreviousPage: forwardRef((props, ref) => (
    <ChevronLeft {...props} ref={ref} />
  )),
  ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
  SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
  ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
  ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />),
}

const columns = [
  {
    title: 'URL',
    field: 'url',
    render: (rowData) => (
      <Link href={rowData.url} target="_blank" rel="noopener" variant="p">
        {rowData.url}
      </Link>
    ),
  },
  { title: 'hash', field: 'hash' },
  { title: 'Describe', field: 'describe' },
  {
    title: 'Created at',
    field: 'createdAt',
    render: (rowData) => (
      <Tooltip arrow title={rowData.createdAt} interactive>
        <span>
          {formatRelative(new Date(rowData.createdAt), new Date(), {
            addSuffix: true,
          })}
        </span>
      </Tooltip>
    ),
  },
  {
    title: 'Updated at',
    field: 'updatedAt',
    render: (rowData) => (
      <Tooltip arrow title={rowData.updatedAt} interactive>
        <span>
          {formatRelative(new Date(rowData.updatedAt), new Date(), {
            addSuffix: true,
          })}
        </span>
      </Tooltip>
    ),
  },
]

export function LinkTable() {
  const state = useSelector((rootState) => rootState.link)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(fetchLinkList())
  }, [dispatch])

  return (
    <Layout>
      <Statistic />

      <MaterialTable
        icons={tableIcons}
        columns={columns}
        data={state.list || []}
        actions={[
          {
            icon: tableIcons.Add,
            tooltip: 'Add link',
            isFreeAction: true,
            onClick: () => alert('You want to add a new row'),
          },
          {
            icon: tableIcons.Update,
            tooltip: 'Update link',
            isFreeAction: true,
            onClick: () => dispatch(fetchLinkList()),
          },
          {
            icon: tableIcons.Save,
            tooltip: 'Save link',
            onClick: (event, rowData) => alert(`You saved ${rowData.name}`),
          },
          {
            icon: tableIcons.Delete,
            tooltip: 'Delete link',
            onClick: (event, rowData) =>
              confirm(`You want to delete ${rowData.name}`), // eslint-disable-line
          },
        ]}
        options={{
          exportButton: true,
        }}
        title="Link Table"
      />
    </Layout>
  )
}

export default withAuthSync(LinkTable)
