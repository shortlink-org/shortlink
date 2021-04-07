import { forwardRef, useEffect } from 'react';
import { useSelector, useDispatch } from "react-redux";
import MaterialTable from "material-table"
import AddBox from '@material-ui/icons/AddBox';
import Update from '@material-ui/icons/Update';
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
import { Layout } from '../components';
import { fetchLinkList } from "../store";

const tableIcons = {
  Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
  Update: forwardRef((props, ref) => <Update {...props} ref={ref} />),
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

export function LinkTableContent() {
  const state = useSelector((state) => state.link);
  const dispatch = useDispatch();

  useEffect(() => {
		dispatch(fetchLinkList());
	}, [dispatch]);

  return (
    <MaterialTable
      icons={tableIcons}
      columns={[
        { title: "URL", field: "Url" },
        { title: "hash", field: "Hash" },
        { title: "Describe", field: "Describe" },
        {
          title: "Created at",
          field: "CreatedAt",
        },
        {
          title: "Updated at",
          field: "UpdatedAt",
        }
      ]}
      data={state.list || []}
      actions={[
        {
          icon: tableIcons.Add,
          tooltip: 'Add link',
          isFreeAction: true,
          onClick: (event) => alert("You want to add a new row")
        },
        {
          icon: tableIcons.Update,
          tooltip: 'Update link',
          isFreeAction: true,
          onClick: (event) => dispatch(fetchLinkList())
        },
        {
          icon: tableIcons.Save,
          tooltip: 'Save link',
          onClick: (event, rowData) => alert("You saved " + rowData.name)
        },
        {
          icon: tableIcons.Delete,
          tooltip: 'Delete link',
          onClick: (event, rowData) => confirm("You want to delete " + rowData.name)
        }
      ]}
      options={{
        exportButton: true
      }}
      title="Link Table"
    />
  );
}

export default function LinkTable() {
  return <Layout content={LinkTableContent()} />;
}
