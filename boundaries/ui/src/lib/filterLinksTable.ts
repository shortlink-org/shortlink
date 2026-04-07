import type { LinkTableItem } from '@/components/Page/user/linksTable'

export function filterLinksBySearch(rows: LinkTableItem[], query: string): LinkTableItem[] {
  const needle = query.trim().toLowerCase()
  if (!needle) return rows
  return rows.filter((r) =>
    [r.url, r.hash, r.describe ?? '', r.created_at, r.updated_at]
      .join(' ')
      .toLowerCase()
      .includes(needle),
  )
}
