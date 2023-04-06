package file

import (
	"fmt"
	"os"
	"sync"

	"github.com/sourcegraph/conc"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	database "github.com/shortlink-org/shortlink/pkg/shortdb/domain/database/v1"
	table "github.com/shortlink-org/shortlink/pkg/shortdb/domain/table/v1"

	v1 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/query/v1"
	"github.com/shortlink-org/shortlink/pkg/shortdb/engine/options"
	"github.com/shortlink-org/shortlink/pkg/shortdb/io_uring"
)

type file struct {
	database *database.DataBase
	path     string
	mu       sync.RWMutex
}

func New(opts ...options.Option) (*file, error) {
	const SHORTDB_PAGE_SIZE = 100

	viper.AutomaticEnv()
	viper.SetDefault("SHORTDB_DEFAULT_DATABASE", "public")   // ShortDB default database
	viper.SetDefault("SHORTDB_PAGE_SIZE", SHORTDB_PAGE_SIZE) // ShortDB default page of size

	var err error
	f := &file{
		database: &database.DataBase{
			Name:   viper.GetString("SHORTDB_DEFAULT_DATABASE"),
			Tables: make(map[string]*table.Table),
		},
	}

	for _, opt := range opts {
		if errApplyOptions := opt(f); errApplyOptions != nil {
			panic(errApplyOptions)
		}
	}

	// if not set path, set temp directory
	if f.path == "" {
		f.path = os.TempDir()
	}

	// init db
	err = f.init()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *file) Exec(query *v1.Query) (interface{}, error) {
	switch query.Type {
	case v1.Type_TYPE_UNSPECIFIED:
		return nil, fmt.Errorf("exec: incorret type")
	case v1.Type_TYPE_SELECT:
		return f.Select(query)
	case v1.Type_TYPE_UPDATE:
		return nil, f.Update(query)
	case v1.Type_TYPE_INSERT:
		return nil, f.Insert(query)
	case v1.Type_TYPE_DELETE:
		return nil, f.Delete(query)
	case v1.Type_TYPE_CREATE_TABLE:
		return nil, f.CreateTable(query)
	case v1.Type_TYPE_DROP_TABLE:
		return nil, f.DropTable(query.TableName)
	case v1.Type_TYPE_CREATE_INDEX:
		return nil, f.CreateIndex(query)
	case v1.Type_TYPE_DELETE_INDEX:
		return nil, f.DropIndex(query.TableName)
	}

	return nil, nil
}

func (f *file) init() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// create directory if not exist
	err := os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return err
	}

	// create file if not exist
	fileOpenFile, err := f.createFile(fmt.Sprintf("%s.db", f.database.Name))
	if err != nil {
		return err
	}
	defer func() {
		_ = fileOpenFile.Close() // #nosec
	}()

	// init io_uring
	err = io_uring.Init()
	if err != nil {
		return err
	}
	defer io_uring.Cleanup()

	go func() {
		for errIOUring := range io_uring.Err() {
			fmt.Println(errIOUring)
		}
	}()

	payload := []byte{}
	var wg conc.WaitGroup

	// Read a file.
	err = io_uring.ReadFile(fileOpenFile.Name(), func(buf []byte) {
		wg.Go(func() {
			payload = buf
		})
	})
	if err != nil {
		return err
	}

	io_uring.Poll()
	wg.Wait()

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, f.database)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *file) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// create database if not exist
	databaseFile, err := f.createFile(fmt.Sprintf("%s.db", f.database.Name))
	if err != nil {
		return err
	}
	defer func() {
		_ = databaseFile.Close() // #nosec
	}()

	// init io_uring
	err = io_uring.Init()
	if err != nil {
		return err
	}
	defer io_uring.Cleanup()

	go func() {
		for err := range io_uring.Err() {
			fmt.Println(err)
		}
	}()

	var wg conc.WaitGroup

	// save last page
	for tableName := range f.database.Tables {
		err = f.savePage(tableName, f.database.Tables[tableName].Stats.PageCount)
		if err != nil {
			return err
		}

		// clear cache
		err = f.clearPages(tableName)
		if err != nil {
			return err
		}
	}

	payload, err := proto.Marshal(f.database)
	if err != nil {
		return err
	}

	// save database
	err = io_uring.WriteFile(databaseFile.Name(), payload, 0o644, func(n int) { // nolint:gomnd
		wg.Go(func() {})
		// handle n
	})
	if err != nil {
		return err
	}

	// Call Poll to let the kernel know to read the entries.
	io_uring.Poll()
	// Wait till all callbacks are done.
	wg.Wait()

	return nil
}

func (f *file) createFile(name string) (*os.File, error) {
	return os.OpenFile(fmt.Sprintf("%s/%s", f.path, name), os.O_RDONLY|os.O_CREATE, os.ModePerm) // #nosec
}

func (f *file) writeFile(name string, payload []byte) error {
	err := os.WriteFile(name, payload, 0o600) // nolint:gomnd
	if err != nil {
		return err
	}

	return nil
}
