import type { ColumnHelper } from '@tanstack/react-table'

declare module '@shortlink-org/ui-kit' {
  export const DataTable: React.ComponentType<any>
  export const DataTableWithSuspense: React.ComponentType<any>
  export const DataTableWithErrorBoundary: React.ComponentType<any>
  export const createDataTableColumnHelper: <TData extends Record<string, unknown>>() => ColumnHelper<TData>
  export const SearchForm: React.ComponentType<any>
  export const ScrollToTopButton: React.ComponentType<any>
  export const Sidebar: React.ComponentType<any>
  export const ToggleDarkMode: React.ComponentType<any>
}
