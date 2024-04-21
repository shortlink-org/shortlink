package file

import (
	"context"
	"fmt"
	"os"

	"github.com/sasha-s/go-deadlock"
	"github.com/sourcegraph/conc"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	database "github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/domain/database/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/domain/query/v1"
	table "github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/domain/table/v1"
	"github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/engine/options"
	"github.com/shortlink-org/shortlink/boundaries/shortdb/shortdb/io_uring"
)

type File struct {
	database *database.DataBase
	path     string
	mu       deadlock.RWMutex
}

func New(ctx context.Context, opts ...options.Option) (*File, error) {
	const SHORTDB_PAGE_SIZE = 100

	viper.AutomaticEnv()
	viper.SetDefault("SHORTDB_DEFAULT_DATABASE", "public")   // ShortDB default database
	viper.SetDefault("SHORTDB_PAGE_SIZE", SHORTDB_PAGE_SIZE) // ShortDB default page of size

	var err error
	f := &File{
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
	err = f.init(ctx)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) Exec(query *v1.Query) (any, error) {
	switch query.GetType() {
	case v1.Type_TYPE_UNSPECIFIED:
		return nil, ErrIncorrectType
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
		return nil, f.DropTable(query.GetTableName())
	case v1.Type_TYPE_CREATE_INDEX:
		return nil, f.CreateIndex(query)
	case v1.Type_TYPE_DELETE_INDEX:
		return nil, f.DropIndex(query.GetTableName())
	}

	return nil, nil
}

func (f *File) init(ctx context.Context) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// create directory if not exist
	err := os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return err
	}

	// create file if not exist
	fileOpenFile, err := f.createFile(fmt.Sprintf("%s.db", f.database.GetName()))
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
		for {
			select {
			case <-ctx.Done():
				return
			case errIOUring := <-io_uring.Err():
				fmt.Println(errIOUring)
			}
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

func (f *File) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// create a database if not exist
	databaseFile, err := f.createFile(fmt.Sprintf("%s.db", f.database.GetName()))
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

	var wg conc.WaitGroup

	// save last page
	for tableName := range f.database.GetTables() {
		err = f.savePage(tableName, f.database.GetTables()[tableName].GetStats().GetPageCount())
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
	err = io_uring.WriteFile(databaseFile.Name(), payload, 0o644, func(n int) { //nolint:mnd,revive // #nosec
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

func (f *File) createFile(name string) (*os.File, error) {
	return os.OpenFile(fmt.Sprintf("%s/%s", f.path, name), os.O_RDONLY|os.O_CREATE, os.ModePerm) // #nosec
}

func (f *File) writeFile(name string, payload []byte) error {
	err := os.WriteFile(name, payload, 0o600) //nolint:mnd,revive // #nosec
	if err != nil {
		return err
	}

	return nil
}
