import { forwardRef } from 'react';
import MaterialTable from "material-table"
import AddBox from '@material-ui/icons/AddBox';
import ArrowDownward from '@material-ui/icons/ArrowDownward';
import Check from '@material-ui/icons/Check';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import Clear from '@material-ui/icons/Clear';
import DeleteOutline from '@material-ui/icons/DeleteOutline';
import Edit from '@material-ui/icons/Edit';
import FilterList from '@material-ui/icons/FilterList';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Remove from '@material-ui/icons/Remove';
import SaveAlt from '@material-ui/icons/SaveAlt';
import Save from '@material-ui/icons/Save';
import Search from '@material-ui/icons/Search';
import ViewColumn from '@material-ui/icons/ViewColumn';
import Layout from '../components/Layout.js';

const tableIcons = {
  Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
  Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
  Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
  DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
  Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
  Save: forwardRef((props, ref) => <Save {...props} ref={ref} />),
  Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
  Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
  FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
  LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
  NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
  PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
  ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
  SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
  ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
  ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />)
};

const linkListPageContent = (
  <MaterialTable
    icons={tableIcons}
    columns={[
      { title: "URL", field: "url" },
      { title: "hash", field: "hash" },
      { title: "Describe", field: "describe" },
      {
        title: "Created at",
        field: "created_at",
      },
      {
        title: "Updated at",
        field: "updated_at",
      }
    ]}
    data={[
      {
        url: 'http://google.com',
        hash: '4535345',
        describe: 'Test URL for table',
        created_at: 1243432434,
        updated_at: 1243432434,
      }
    ]}
    actions={[
      {
        icon: tableIcons.Add,
        tooltip: 'Add Link',
        isFreeAction: true,
        onClick: (event) => alert("You want to add a new row")
      },
      {
        icon: tableIcons.Save,
        tooltip: 'Save User',
        onClick: (event, rowData) => alert("You saved " + rowData.name)
      },
      {
        icon: tableIcons.Delete,
        tooltip: 'Delete Link',
        onClick: (event, rowData) => confirm("You want to delete " + rowData.name)
      }
    ]}
    options={{
      exportButton: true
    }}
    title="Link Table"
  />
);

export default function LinkTable() {
  return <Layout content={linkListPageContent} />;
}
