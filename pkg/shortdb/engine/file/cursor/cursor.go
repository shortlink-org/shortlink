package cursor

import (
	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

func New(table *table.Table, isEnd bool) (*Cursor, error) {
	cursor := &Cursor{
		Table:      table,
		RowId:      0,
		PageId:     0,
		EndOfTable: isEnd,
	}

	if isEnd {
		cursor.RowId = table.Stats.RowsCount
		cursor.PageId = table.Stats.PageCount
	}

	return cursor, nil
}

func (c *Cursor) Advance() {
	c.wc.RLock()
	defer c.wc.RUnlock()

	if c.RowId > 0 && c.RowId%c.Table.Option.PageSize == 0 {
		c.PageId = int32(c.RowId / c.Table.Option.PageSize)
	}

	if (c.Table.Stats.RowsCount - 1) == c.RowId {
		c.EndOfTable = true
	} else {
		c.RowId += 1
	}
}

func (c *Cursor) Value() (*table.Row, error) {
	c.wc.RLock()
	defer c.wc.RUnlock()

	if c.Table.Pages == nil {
		return nil, &ErrorGetPage{}
	}

	page := c.Table.Pages[c.PageId]
	if len(page.Rows) == 0 {
		page.Rows = make([]*table.Row, c.Table.Option.PageSize)
	}

	rowNum := int(c.RowId) % len(page.Rows)

	if page.Rows[rowNum] == nil {
		page.Rows[rowNum] = &table.Row{}
	}

	return page.Rows[rowNum], nil
}
