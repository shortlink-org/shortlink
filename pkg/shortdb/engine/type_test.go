package engine

import (
	"fmt"
	"os"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/pkg/shortdb/engine/file"
	parser "github.com/batazor/shortlink/pkg/shortdb/parser/v1"
)

func BenchmarkEngine(b *testing.B) {
	// set engine
	id, err := uuid.NewV4()
	assert.Nil(b, err)
	path := fmt.Sprintf("/tmp/shortdb_test_%s", id.String())

	store, err := New("file", file.SetName("testDatabase"), file.SetPath(path))
	assert.Nil(b, err)

	b.Cleanup(func() {
		err = os.RemoveAll(path)
		assert.Nil(b, err)
	})

	b.Run("CREATE DATABASE", func(b *testing.B) {
		// create table
		qCreateTable, err := parser.New("create table users (id integer, name string, active bool)")
		assert.Nil(b, err)

		_, _ = (*store).Exec(qCreateTable.Query)
	})

	b.Run("INSERT INTO USERS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qInsertUsers, err := parser.New("insert into users ('id', 'name', 'active') VALUES ('1', 'Ivan', 'false')")
			assert.Nil(b, err)

			err = (*store).Insert(qInsertUsers.Query)
			assert.Nil(b, err)

			// save data
			err = (*store).Close()
			assert.Nil(b, err)
		}
	})

	b.Run("SELECT USERS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qInsertUsers, err := parser.New("select id, name, active from users limit 5")
			assert.Nil(b, err)

			resp, err := (*store).Select(qInsertUsers.Query)
			assert.Nil(b, err)
			assert.Equal(b, 5, len(resp))

			// save data
			err = (*store).Close()
			assert.Nil(b, err)
		}
	})
}
