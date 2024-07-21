package file

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"

	page "github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/domain/page/v1"
)

func (f *File) getPage(nameTable string, p int32) error { //nolint:unused // ignore
	t := f.database.GetTables()[nameTable]

	// read page
	pageFile, err := f.loadPage(f.pageName(nameTable))
	if err != nil {
		return err
	}

	t.Pages[p] = pageFile

	return nil
}

func (f *File) addPage(nameTable string) (int32, error) {
	t := f.database.GetTables()[nameTable]

	if t.GetStats().GetRowsCount()%t.GetOption().GetPageSize() == 0 { //nolint:nestif // ignore
		if t.GetPages() == nil {
			t.Pages = make(map[int32]*page.Page, 0)
		}

		t.Stats.PageCount += 1
		t.Pages[t.GetStats().GetPageCount()] = &page.Page{Rows: []*page.Row{}}

		// create a page file
		newPageFile, err := f.createFile(f.pageName(nameTable))
		if err != nil {
			return t.GetStats().GetPageCount(), err
		}

		err = newPageFile.Close()
		if err != nil {
			return t.GetStats().GetPageCount(), err
		}

		// if this not first page, save current date
		if t.GetStats().GetPageCount() > 0 && t.GetPages()[t.GetStats().GetPageCount()-1] != nil {
			// save data after clear memory page
			err = f.savePage(nameTable, t.GetStats().GetPageCount()-1)
			if err != nil {
				return t.GetStats().GetPageCount(), err
			}

			// clear old page
			err = f.clearPage(nameTable, t.GetStats().GetPageCount()-1)
			if err != nil {
				return t.GetStats().GetPageCount(), err
			}
		}
	}

	return t.GetStats().GetPageCount(), nil
}

func (f *File) savePage(nameTable string, pageCount int32) error {
	t := f.database.GetTables()[nameTable]

	if pageCount == -1 {
		return nil
	}

	// save date
	openFile, err := f.createFile(fmt.Sprintf("%s_%s_%d.page", f.database.GetName(), nameTable, pageCount))
	if err != nil {
		return err
	}

	defer func() {
		_ = openFile.Close() // #nosec
	}()

	payload, err := proto.Marshal(t.GetPages()[pageCount])
	if err != nil {
		return err
	}

	// Write something
	err = f.writeFile(openFile.Name(), payload)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) clearPage(nameTable string, pageCount int32) error { //nolint:unparam // ignore param
	f.database.Tables[nameTable].Pages[pageCount] = nil

	return nil
}

//nolint:unparam // ignore param
func (f *File) clearPages(nameTable string) error {
	f.database.Tables[nameTable].Pages = nil

	return nil
}

func (f *File) loadPage(path string) (*page.Page, error) {
	p := page.Page{}

	payload, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, &p)
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
}

func (f *File) pageName(nameTable string) string {
	return fmt.Sprintf("%s_%s_%d.page", f.database.GetName(), nameTable, f.database.GetTables()[nameTable].GetStats().GetPageCount())
}
