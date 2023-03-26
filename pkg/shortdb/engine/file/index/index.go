package index

import (
	"reflect"
	"strconv"

	index "github.com/shortlink-org/shortlink/pkg/shortdb/domain/index/v1"
	v2 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/index/v1"
	page "github.com/shortlink-org/shortlink/pkg/shortdb/domain/page/v1"
	binary_tree "github.com/shortlink-org/shortlink/pkg/shortdb/engine/file/index/binary-tree"
)

func New(in *index.Index, rows []*page.Row) (Index[any], error) {
	var tree Index[any]

	switch in.Type {
	case v2.Type_TYPE_BINARY_SEARCH:
		tree = binary_tree.New(func(a, b any) int {
			switch x, y := reflect.TypeOf(a), reflect.TypeOf(b); true {
			case x.String() == "int" && y.String() == "int":
				return a.(int) - b.(int)
			default:
				return 0
			}
		})

		for i := range rows {
			v, err := strconv.Atoi(string(rows[i].GetValue()["id"]))
			if err != nil {
				return nil, err
			}

			err = tree.Insert(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return tree, nil
}
