package file

import (
	"fmt"

	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
	v12 "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

func (f *file) Select(query *v1.Query) ([]*v12.Row, error) {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check table
	if f.database.Tables[query.TableName] == nil {
		return nil, fmt.Errorf("at SELECT: not exist table")
	}

	// response
	response := []*v12.Row{}

	for _, page := range f.database.Tables[query.TableName].Pages {
		for _, row := range page.Rows {
			record := &v12.Row{
				Value: map[string][]byte{},
			}

			for _, field := range query.Fields {
				if row.Value[field] == nil {
					return nil, fmt.Errorf("at SELECT: incorrect name fields %s in table %s", field, query.TableName)
				}

				record.Value[field] = row.Value[field]
			}
			response = append(response, record)
		}
	}

	return response, nil
}

func (f *file) Update(query *v1.Query) error {
	//TODO implement me
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
	record := &v12.Row{
		Value: map[string][]byte{},
	}
	for index, field := range query.Fields {
		if f.database.Tables[query.TableName].Fields[field].String() == "" {
			return fmt.Errorf("at INSERT INTO: incorrect type fields %s in table %s", field, query.TableName)
		}

		record.Value[field] = []byte(query.Inserts[0].Items[index])
	}

	// insert
	lastPageIndex, err := f.createPage(query)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new page")
	}

	// insert to last page
	f.database.Tables[query.TableName].Pages[lastPageIndex].Rows = append(f.database.Tables[query.TableName].Pages[lastPageIndex].Rows, record)

	// update stats
	f.database.Tables[query.TableName].Stats.RowsCount += 1

	return nil
}

func (f *file) Delete(query *v1.Query) error {
	//TODO implement me
	return nil
}

func (f *file) createPage(query *v1.Query) (int32, error) {
	if f.database.Tables[query.TableName].Stats.RowsCount%f.pageSize == 0 {
		f.database.Tables[query.TableName].Pages = append(f.database.Tables[query.TableName].Pages, &v12.Page{Rows: []*v12.Row{}})
		f.database.Tables[query.TableName].Stats.PageCount += 1
	}

	return f.database.Tables[query.TableName].Stats.PageCount - 1, nil
}
