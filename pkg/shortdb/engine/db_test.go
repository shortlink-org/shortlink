package engine

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/pkg/shortdb/engine/file"
	parser "github.com/batazor/shortlink/pkg/shortdb/parser/v1"
)

func TestDatabase(t *testing.T) {
	// set engine
	path := fmt.Sprintf("/tmp/shortdb_test_unit")

	store, err := New("file", file.SetName("testDatabase"), file.SetPath(path))
	assert.Nil(t, err)

	t.Cleanup(func() {
		err = os.RemoveAll(path)
		assert.Nil(t, err)
	})

	t.Run("CREATE TABLE", func(t *testing.T) {
		// create table
		qCreateTable, err := parser.New("create table users (id integer, name string, active bool)")
		assert.Nil(t, err)

		_, _ = store.Exec(qCreateTable.Query)
	})

	t.Run("INSERT INTO USERS", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			qInsertUsers, err := parser.New(fmt.Sprintf("insert into users ('id', 'name', 'active') VALUES ('%d', 'Ivan', 'false')", i))
			assert.Nil(t, err)

			err = store.Insert(qInsertUsers.Query)
			assert.Nil(t, err)
		}

		// save data
		err = store.Close()
		assert.Nil(t, err)
	})

	//t.Run("SELECT USERS", func(t *testing.T) {
	//	qInsertUsers, err := parser.New("select id, name, active from users limit 300")
	//	assert.Nil(t, err)
	//
	//	resp, err := (*store).Select(qInsertUsers.Query)
	//	assert.Nil(t, err)
	//	assert.Equal(t, 300, len(resp))
	//})
}
