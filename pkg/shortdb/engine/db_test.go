package engine

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/pkg/shortdb/engine/file"
	parser "github.com/batazor/shortlink/pkg/shortdb/parser/v1"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

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

		// save data
		err = store.Close()
		assert.Nil(t, err)
	})

	t.Run("INSERT INTO USERS SINGLE", func(t *testing.T) {
		qInsertUsers, errParser := parser.New(fmt.Sprintf("insert into users ('id', 'name', 'active') VALUES ('1', 'Ivan', 'false')"))
		assert.Nil(t, errParser)

		errParser = store.Insert(qInsertUsers.Query)
		assert.Nil(t, errParser)

		errParser = store.Insert(qInsertUsers.Query)
		assert.Nil(t, errParser)

		errParser = store.Insert(qInsertUsers.Query)
		assert.Nil(t, errParser)

		// save data
		errClose := store.Close()
		assert.Nil(t, errClose)
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

	t.Run("INSERT INTO USERS +173", func(t *testing.T) {
		for i := 0; i < 173; i++ {
			qInsertUsers, err := parser.New(fmt.Sprintf("insert into users ('id', 'name', 'active') VALUES ('%d', 'Ivan', 'false')", i))
			assert.Nil(t, err)

			err = store.Insert(qInsertUsers.Query)
			assert.Nil(t, err)
		}

		// save data
		err = store.Close()
		assert.Nil(t, err)
	})

	t.Run("INSERT INTO USERS +207", func(t *testing.T) {
		for i := 0; i < 207; i++ {
			qInsertUsers, err := parser.New(fmt.Sprintf("insert into users ('id', 'name', 'active') VALUES ('%d', 'Ivan', 'false')", i))
			assert.Nil(t, err)

			err = store.Insert(qInsertUsers.Query)
			assert.Nil(t, err)
		}

		// save data
		err = store.Close()
		assert.Nil(t, err)
	})

	t.Run("SELECT USERS WITH LIMIT 300", func(t *testing.T) {
		qSelectUsers, err := parser.New("select id, name, active from users limit 300")
		assert.Nil(t, err)

		resp, err := store.Select(qSelectUsers.Query)
		assert.Nil(t, err)
		assert.Equal(t, 300, len(resp))
	})

	t.Run("SELECT USERS WITH WHERE id=99 AND LIMIT 2", func(t *testing.T) {
		qSelectUsers, err := parser.New("select id, name, active from users where id='99' limit 2")
		assert.Nil(t, err)

		resp, err := store.Select(qSelectUsers.Query)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(resp))
	})

	t.Run("SELECT USERS FULL SCAN", func(t *testing.T) {
		qSelectUsers, err := parser.New("select id, name, active from users")
		assert.Nil(t, err)

		resp, err := store.Select(qSelectUsers.Query)
		assert.Nil(t, err)
		assert.Equal(t, 1383, len(resp))
	})

	t.Run("CREATE INDEX BINARY", func(t *testing.T) {
		qCreateIndex, err := parser.New("CREATE INDEX userId ON users USING BINARY (id);")
		assert.Nil(t, err)

		err = store.CreateIndex(qCreateIndex.Query)
		assert.Nil(t, err)
	})
}
