package file

import (
	"fmt"

	"github.com/batazor/shortlink/pkg/shortdb/engine/file/cursor"
	v1 "github.com/batazor/shortlink/pkg/shortdb/query/v1"
	v12 "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

func (f *file) Select(query *v1.Query) ([]*v12.Row, error) {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check table
	table := f.database.Tables[query.TableName]
	if table == nil {
		return nil, fmt.Errorf("at SELECT: not exist table")
	}

	// set flags
	isLimit := query.Limit >= 0

	// response
	response := make([]*v12.Row, 0)

retrySelect:
	currentRow, err := cursor.New(table, false)
	if err != nil {
		return nil, fmt.Errorf("at SELECT: error create a new cursor")
	}

	for !currentRow.EndOfTable && isLimit {
		record, errGetValue := currentRow.Value()
		if errGetValue != nil {
			switch errGetValue.(type) {
			case *cursor.ErrorGetPage:
				pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, query.TableName, currentRow.PageId)
				payload, errLoadPage := f.loadPage(pagePath)
				if errLoadPage != nil {
					return nil, errLoadPage
				}

				// load data
				table.Pages = []*v12.Page{payload}
				goto retrySelect
			default:
				return nil, errGetValue
			}
		}

		for _, field := range query.Fields {
			if record.Value[field] == nil {
				return nil, fmt.Errorf("at SELECT: incorrect name fields %s in table %s", field, query.TableName)
			}
		}
		response = append(response, record)

		if isLimit {
			query.Limit -= 1

			if query.Limit < 1 {
				isLimit = false
			}
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
	f.mc.Lock()
	defer f.mc.Unlock()

	// check table
	if f.database.Tables[query.TableName] == nil {
		return fmt.Errorf("at INSERT INTO: not exist table")
	}

	// check values and create row record
	record := v12.Row{
		Value: make(map[string][]byte),
	}
	for index, field := range query.Fields {
		if f.database.Tables[query.TableName].Fields[field].String() == "" {
			return fmt.Errorf("at INSERT INTO: incorrect type fields %s in table %s", field, query.TableName)
		}

		record.Value[field] = []byte(query.Inserts[0].Items[index])
	}

	// insert
	_, err := f.addPage(query.TableName)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new page")
	}

	// insert to last page
	currentRow, err := cursor.New(f.database.Tables[query.TableName], true)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new cursor")
	}

	// iterator to next value
	currentRow.Advance()

	row, err := currentRow.Value()
	if err != nil {
		switch err.(type) {
		case *cursor.ErrorGetPage:
			pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, query.TableName, currentRow.PageId)
			_, errLoadPage := f.loadPage(pagePath)
			if errLoadPage != nil {
				return errLoadPage
			}
		default:
			return fmt.Errorf("at INSERT INTO: error get value from cursor")
		}
	}
	row.Value = record.Value

	// update stats
	f.database.Tables[query.TableName].Stats.RowsCount += 1

	return nil
}

func (f *file) Delete(query *v1.Query) error {
	// TODO implement me
	return nil
}
