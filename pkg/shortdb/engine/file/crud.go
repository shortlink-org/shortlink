package file

import (
	"fmt"

	page "github.com/shortlink-org/shortlink/pkg/shortdb/domain/page/v1"
	v1 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/query/v1"

	"github.com/shortlink-org/shortlink/pkg/shortdb/engine/file/cursor"
)

func (f *file) Select(query *v1.Query) ([]*page.Row, error) {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check table
	t := f.database.Tables[query.TableName]
	if t == nil {
		return nil, fmt.Errorf("at SELECT: not exist table")
	}

	if len(query.Fields) == 0 {
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
		if t.Pages[currentRow.PageId] == nil {
			pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, t.Name, currentRow.PageId)
			payload, errLoadPage := f.loadPage(pagePath)
			if errLoadPage != nil {
				return nil, errLoadPage
			}

			if t.Pages == nil {
				t.Pages = make(map[int32]*page.Page, 0)
			}

			t.Pages[currentRow.PageId] = payload
		}

		// get value
		record, errGetValue := currentRow.Value()
		if errGetValue != nil {
			return nil, errGetValue
		}

		for _, field := range query.Fields {
			if record.Value[field] == nil {
				return nil, fmt.Errorf("at SELECT: incorrect name fields %s in table %s", field, query.TableName)
			}
		}
		if query.IsFilter(record, t.Fields) {
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
	f.mc.Lock()
	defer f.mc.Unlock()

	// check the table's existence
	t := f.database.Tables[query.TableName]
	if t == nil {
		return fmt.Errorf("at INSERT INTO: not exist table")
	}

	// check if a new page needs to be created
	_, err := f.addPage(query.TableName)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new page")
	}

	if t.Stats.PageCount > -1 && t.Pages[t.Stats.PageCount] == nil {
		// load page
		pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, t.Name, t.Stats.PageCount)
		payload, errLoadPage := f.loadPage(pagePath)
		if errLoadPage != nil {
			return errLoadPage
		}

		if t.Pages == nil {
			t.Pages = make(map[int32]*page.Page, 0)
		}

		t.Pages[t.Stats.PageCount] = payload
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
	for index, field := range query.Fields {
		if t.Fields[field].String() == "" {
			return fmt.Errorf("at INSERT INTO: incorrect type fields %s in table %s", field, query.TableName)
		}

		record.Value[field] = []byte(query.Inserts[0].Items[index])
	}
	row.Value = record.Value

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
