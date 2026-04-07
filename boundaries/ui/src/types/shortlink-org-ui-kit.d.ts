declare module '@shortlink-org/ui-kit' {
  import type { ComponentType } from 'react'
  type ColumnHelper<TData extends object> = import('@tanstack/react-table').ColumnHelper<TData>

  /* Page / marketing */
  export const AppHeader: ComponentType<any>
  export const Footer: ComponentType<any>
  export const Header: ComponentType<any>
  export const Newsletter: ComponentType<any>
  export const MultiColumnLayout: ComponentType<any>
  export const PriceTable: ComponentType<any>

  /* Navigation / shell */
  export const Sidebar: ComponentType<any>
  export const SidebarShell: ComponentType<any>
  export const SecondaryMenu: ComponentType<any>
  export const ScrollToTopButton: ComponentType<any>
  export const Breadcrumbs: ComponentType<any>
  export const FlyoutMenu: ComponentType<any>
  export const Drawer: ComponentType<any>

  /* Data */
  export const DataTable: ComponentType<any>
  export const DataTableWithSuspense: ComponentType<any>
  export const DataTableWithErrorBoundary: ComponentType<any>
  export const Pagination: ComponentType<any>
  export const SearchForm: ComponentType<any>
  export function createDataTableColumnHelper<TData extends object>(): ColumnHelper<TData>
  export function useDataTable(...args: any[]): any

  /* UI primitives */
  export const Button: ComponentType<any>
  export const ToggleDarkMode: ComponentType<any>
  export const Skeleton: ComponentType<any>
  export const CardSkeleton: ComponentType<any>
  export const StatCard: ComponentType<any>
  export const ProfileIdentity: ComponentType<any>
  export const SuspenseFallback: ComponentType<any>
  export const Timeline: ComponentType<any>

  /* Cards / blocks */
  export const FeatureCard: ComponentType<any>
  export const GithubRepository: ComponentType<any>
  export const FeedbackPanel: ComponentType<any>
  export const FamilyDialog: ComponentType<any>

  /* Commerce / marketplace (vitrine) */
  export const AddToCartButton: ComponentType<any>
  export const Basket: ComponentType<any>
  export const BasketItem: ComponentType<any>
  export const BasketSummary: ComponentType<any>
  export const ProductColorSelector: ComponentType<any>
  export const ProductDescription: ComponentType<any>
  export const ProductGrid: ComponentType<any>
  export const ProductImageGallery: ComponentType<any>
  export const ProductPage: ComponentType<any>
  export const ProductQuickView: ComponentType<any>
  export const ProductReviews: ComponentType<any>
  export const ProductSizeSelector: ComponentType<any>
  export const MarketplaceLeaderboard: ComponentType<any>
  export const LeaderboardFilters: ComponentType<any>
  export const LeaderboardList: ComponentType<any>
  export const LeaderboardPodium: ComponentType<any>

  /* Errors */
  export const ErrorBoundary: ComponentType<any>

  /* Theme (MUI) */
  export const theme: any

  /* Spring / motion helpers */
  export function createSpring(...args: any[]): any
  export const springBouncy: any
  export const springCSSVariables: any
  export const springDefault: any
  export const springGentle: any
  export const springSlow: any
  export const springSnappy: any
  export const springStiff: any
  export const springWobbly: any

  /* Basket math */
  export function formatMoneyFromCents(...args: any[]): any
  export function getLineTotalCents(...args: any[]): any
  export function getUnitPriceCents(...args: any[]): any
  export function parsePriceStringToCents(...args: any[]): any
  export function resolveBasketSubtotalDisplay(...args: any[]): any
  export function sumBasketSubtotalCents(...args: any[]): any

  /* React utils */
  export function useControllableState(...args: any[]): any
  export function useAsyncData(...args: any[]): any
  export function clearPromiseCache(...args: any[]): any
}
