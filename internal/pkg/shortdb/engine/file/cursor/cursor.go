package cursor

import (
	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

func New(table *table.Table, isEnd bool) (*Cursor, error) {
	cursor := &Cursor{
		Table:      table,
		RowId:      0,
		EndOfTable: isEnd,
	}

	if isEnd {
		cursor.RowId = table.Stats.RowsCount
	}

	return cursor, nil
}

func (c *Cursor) Advance() {
	c.wc.Lock()
	defer c.wc.Unlock()

	if c.Table.Stats.RowsCount == c.RowId {
		c.EndOfTable = true
	} else {
		c.RowId += 1
	}
}

func (c *Cursor) Value() (*table.Row, error) {
	page := c.Table.Pages[c.Table.Stats.PageCount-1]
	if len(page.Rows) == 0 {
		c.wc.Lock()
		page.Rows = make([]*table.Row, c.Table.Option.PageSize)
		c.wc.Unlock()
	}

	rowNum := len(page.Rows) - 1

	if page.Rows[rowNum] == nil {
		c.wc.Lock()
		page.Rows[rowNum] = &table.Row{}
		c.wc.Unlock()
	}

	return page.Rows[rowNum], nil
}
