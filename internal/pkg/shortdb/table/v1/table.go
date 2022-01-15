package v1

func (t *Table) AddPage() (int32, error) {
	if t.Stats.RowsCount%t.Option.PageSize == 0 {
		t.Pages = append(t.Pages, &Page{Rows: []*Row{}})
		t.Stats.PageCount += 1
	}

	return t.Stats.PageCount - 1, nil
}
