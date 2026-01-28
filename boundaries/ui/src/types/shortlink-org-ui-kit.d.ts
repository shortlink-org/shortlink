declare module '@shortlink-org/ui-kit' {
  type ColumnHelper<TData extends object> = import('@tanstack/react-table').ColumnHelper<TData>

  export const DataTable: React.ComponentType<any>
  export const DataTableWithSuspense: React.ComponentType<any>
  export const DataTableWithErrorBoundary: React.ComponentType<any>
  export const createDataTableColumnHelper: <TData extends object>() => ColumnHelper<TData>
  export const SearchForm: React.ComponentType<any>
  export const ScrollToTopButton: React.ComponentType<any>
  export const Sidebar: React.ComponentType<any>
  export const ToggleDarkMode: React.ComponentType<any>
}
