package file

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

func (f *file) AddPage(nameTable string) (int32, error) {
	t := f.database.Tables[nameTable]

	if t.Stats.RowsCount%t.Option.PageSize == 0 {
		// create a page file
		newPageFile, err := f.createFile(fmt.Sprintf("%s_%s_%d.page", f.database.Name, nameTable, t.Stats.PageCount))
		if err != nil {
			return t.Stats.PageCount, err
		}

		err = newPageFile.Close()
		if err != nil {
			return t.Stats.PageCount, err
		}

		// if this not first page, save current date
		if t.Stats.PageCount != 0 {
			oldPageCount := t.Stats.PageCount - 1
			err = f.savePage(nameTable, oldPageCount)
			if err != nil {
				return t.Stats.PageCount, err
			}

			// clear old page
			err = f.clearPage(nameTable, oldPageCount)
			if err != nil {
				return t.Stats.PageCount, err
			}
		}

		t.Pages = append(t.Pages, &table.Page{Rows: []*table.Row{}})
		t.Stats.PageCount += 1
	}

	return t.Stats.PageCount, nil
}

func (f *file) savePage(nameTable string, pageCount int32) error {
	t := f.database.Tables[nameTable]

	// save date
	oldPageFile, err := f.createFile(fmt.Sprintf("%s_%s_%d.page", f.database.Name, nameTable, pageCount))
	if err != nil {
		return err
	}

	defer func() {
		_ = oldPageFile.Close() // #nosec
	}()

	payload, err := proto.Marshal(t.Pages[pageCount])
	if err != nil {
		return err
	}

	// Write something
	err = f.writeFile(oldPageFile.Name(), payload)
	if err != nil {
		return err
	}

	return nil
}

func (f *file) clearPage(nameTable string, pageCount int32) error {
	f.database.Tables[nameTable].Pages[pageCount] = nil
	return nil
}

func (f *file) clearPages(nameTable string) error {
	f.database.Tables[nameTable].Pages = nil
	return nil
}
