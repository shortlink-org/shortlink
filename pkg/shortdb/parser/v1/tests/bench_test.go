package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	parser "github.com/shortlink-org/shortlink/pkg/shortdb/parser/v1"
)

func BenchmarkParser(b *testing.B) {
	b.Run("CREATE TABLE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := parser.New("CREATE TABLE users ( id integer, name text );")
			assert.Nil(b, err)
		}
	})

	b.Run("SELECT", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := parser.New("SELECT a, c, d FROM 'b' WHERE a != '1' LIMIT 5")
			assert.Nil(b, err)
		}
	})

	b.Run("INSERT INTO", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := parser.New("INSERT INTO 'a' (b,c,d) VALUES ('1','2','3'),('4','5','6');")
			assert.Nil(b, err)
		}
	})
}
