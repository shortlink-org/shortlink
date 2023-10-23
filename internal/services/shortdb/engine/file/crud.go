package file

import (
	"fmt"

	page "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/page/v1"
	v1 "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/engine/file/cursor"
)

func (f *file) Select(query *v1.Query) ([]*page.Row, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// check table
	t := f.database.GetTables()[query.GetTableName()]
	if t == nil {
		return nil, fmt.Errorf("at SELECT: not exist table")
	}

	if len(query.GetFields()) == 0 {
		return nil, fmt.Errorf("at SELECT: expected field to SELECT")
	}

	// response
	response := make([]*page.Row, 0)

	currentRow, err := cursor.New(t, false)
	if err != nil {
		return nil, fmt.Errorf("at SELECT: error create a new cursor")
	}

	for !currentRow.EndOfTable {
		// load data
		if t.GetPages()[currentRow.PageId] == nil {
			pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.GetName(), t.GetName(), currentRow.PageId)
			payload, errLoadPage := f.loadPage(pagePath)
			if errLoadPage != nil {
				return nil, errLoadPage
			}

			if t.GetPages() == nil {
				t.Pages = make(map[int32]*page.Page, 0)
			}

			t.Pages[currentRow.PageId] = payload
		}

		// get value
		record, errGetValue := currentRow.Value()
		if errGetValue != nil {
			return nil, errGetValue
		}

		for _, field := range query.GetFields() {
			if record.GetValue()[field] == nil {
				return nil, fmt.Errorf("at SELECT: incorrect name fields %s in table %s", field, query.GetTableName())
			}
		}
		if query.IsFilter(record,t.GetFields())) {
			response = append(response, record)

			if query.IsLimit() {
				query.Limit--
			}
		}

		if !query.IsLimit() {
			break
		}

		currentRow.Advance()
	}

	return response, nil
}

func (f *file) Update(query *v1.Query) error {
	// TODO implement me
	return nil
}

func (f *file) Insert(query *v1.Query) error {
	err := f.insertToTable(query)
	if err != nil {
		return err
	}

	err = f.insertToIndex(query)
	if err != nil {
		return err
	}

	return nil
}

func (f *file) insertToTable(query *v1.Query) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// check the table's existence
	t := f.database.GetTables()[query.GetTableName()]
	if t == nil {
		return fmt.Errorf("at INSERT INTO: not exist table")
	}

	// check if a new page needs to be created
	_, err := f.addPage(query.GetTableName())
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new page")
	}

	if t.GetStats().GetPageCount() > -1 && t.GetPages()[t.GetStats().GetPageCount()] == nil {
		// load page
		pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.GetName(), t.GetName(), t.GetStats().GetPageCount())
		payload, errLoadPage := f.loadPage(pagePath)
		if errLoadPage != nil {
			return errLoadPage
		}

		if t.GetPages() == nil {
			t.Pages = make(map[int32]*page.Page, 0)
		}

		t.Pages[t.GetStats().GetPageCount()] = payload
	}

	// insert to last page
	currentRow, err := cursor.New(t, true)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new cursor")
	}

	row, err := currentRow.Value()
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error get value from cursor")
	}

	// check values and create row record
	record := page.Row{
		Value: make(map[string][]byte),
	}
	for index, field := range query.GetFields() {
		if t.GetFields()[field].String() == "" {
			return fmt.Errorf("at INSERT INTO: incorrect type fields %s in table %s", field, query.GetTableName())
		}

		record.Value[field] = []byte(query.GetInserts()[0].GetItems()[index])
	}
	row.Value = record.GetValue()

	// update stats
	t.Stats.RowsCount += 1

	// iterator to next value
	currentRow.Advance()

	return nil
}

func (f *file) insertToIndex(query *v1.Query) error {
	// TODO implement me
	return nil
}

func (f *file) Delete(query *v1.Query) error {
	// TODO implement me
	return nil
}
