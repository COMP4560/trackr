import TableRowsIcon from "@mui/icons-material/TableRows";
import TableEditor from "./TableEditor";
import TableView from "./TableView";

const Table = {
  name: "Table",
  icon: TableRowsIcon,
  editor: TableEditor,
  view: TableView,

  deserialize: (metadata) => {
    const sort = metadata?.sort || "asc";

    return { sort };
  },

  serialize: (sort) => {
    return JSON.stringify({
      name: Table.name,
      sort,
    });
  },
};

export default Table;
