package engine

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/shortdb/engine/file"
	parser "github.com/shortlink-org/shortlink/pkg/shortdb/parser/v1"
)

func BenchmarkEngine(b *testing.B) {
	// set engine
	path := "/tmp/shortdb_test_unit"

	store, err := New("file", file.SetName("testDatabase"), file.SetPath(path))
	require.NoError(b, err)

	b.Cleanup(func() {
		err = os.RemoveAll(path)
		require.NoError(b, err)
	})

	b.Run("CREATE TABLE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qCreateTable, errParserNew := parser.New("create table users (id integer, name string, active bool)")
			require.NoError(b, errParserNew)

			_, err = store.Exec(qCreateTable.Query)
			require.NoError(b, err)
		}
	})

	b.Run("INSERT INTO USERS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qInsertUsers, errParserNew := parser.New(fmt.Sprintf("insert into users ('id', 'name', 'active') VALUES ('%d', 'Ivan', 'false')", i))
			require.NoError(b, errParserNew)

			errInsert := store.Insert(qInsertUsers.Query)
			require.NoError(b, errInsert)
		}

		// save data
		err = store.Close()
		require.NoError(b, err)
	})

	b.Run("SELECT USERS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qInsertUsers, err := parser.New("select id, name, active from users limit 5")
			require.NoError(b, err)

			resp, err := store.Select(qInsertUsers.Query)
			require.NoError(b, err)
			assert.Equal(b, 5, len(resp))
		}
	})

	b.Run("SELECT USERS WITH WHERE id=99 AND LIMIT 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qSelectUsers, err := parser.New("select id, name, active from users where id='99' limit 2")
			require.NoError(b, err)

			_, err = store.Select(qSelectUsers.Query)
			require.NoError(b, err)
		}
	})

	b.Run("SELECT USERS FULL SCAN", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qSelectUsers, err := parser.New("select id, name, active from users")
			require.NoError(b, err)

			_, err = store.Select(qSelectUsers.Query)
			require.NoError(b, err)
		}
	})

	b.Run("CREATE INDEX BTREE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			qCreateIndex, err := parser.New("CREATE INDEX userId ON users USING BTREE (id);")
			require.NoError(b, err)

			err = store.CreateIndex(qCreateIndex.Query)
			require.NoError(b, err)
		}
	})
}
