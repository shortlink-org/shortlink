package cursor

import (
	page "github.com/shortlink-org/shortlink/internal/boundaries/shortdb/shortdb/domain/page/v1"
	table "github.com/shortlink-org/shortlink/internal/boundaries/shortdb/shortdb/domain/table/v1"
)

func New(t *table.Table, isEnd bool) (*Cursor, error) {
	cursor := &Cursor{
		Table:      t,
		RowId:      0,
		PageId:     0,
		EndOfTable: isEnd,
	}

	if isEnd {
		cursor.RowId = t.GetStats().GetRowsCount()
		cursor.PageId = t.GetStats().GetPageCount()
	}

	return cursor, nil
}

func (c *Cursor) Advance() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.RowId > 0 && c.RowId%c.Table.GetOption().GetPageSize() == 0 {
		c.PageId = int32(c.RowId / c.Table.GetOption().GetPageSize())
	}

	if (c.Table.GetStats().GetRowsCount() - 1) == c.RowId {
		c.EndOfTable = true
	} else {
		c.RowId += 1
	}
}

func (c *Cursor) Value() (*page.Row, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Table.GetPages() == nil {
		return nil, ErrGetPage
	}

	p := c.Table.GetPages()[c.PageId]
	if len(p.GetRows()) == 0 {
		p.Rows = make([]*page.Row, c.Table.GetOption().GetPageSize())
	}

	rowNum := int(c.RowId) % len(p.GetRows())

	if p.GetRows()[rowNum] == nil {
		p.Rows[rowNum] = &page.Row{}
	}

	return p.GetRows()[rowNum], nil
}
