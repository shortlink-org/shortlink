// Copyright (C) 2017 ScyllaDB
// Use of this source code is governed by a ALv2-style
// license that can be found in the LICENSE file.

package qb

import (
	"bytes"
	"fmt"
	"time"
)

// BATCH reference:
// https://cassandra.apache.org/doc/latest/cql/dml.html#batch

// BatchBuilder builds CQL BATCH statements.
type BatchBuilder struct {
	unlogged bool
	counter  bool
	using    using
	stmts    []string
	names    []string
}

// Batch returns a new BatchBuilder.
func Batch() *BatchBuilder {
	return &BatchBuilder{}
}

// ToCql builds the query into a CQL string and named args.
func (b *BatchBuilder) ToCql() (stmt string, names []string) {
	cql := bytes.Buffer{}

	cql.WriteString("BEGIN ")
	if b.unlogged {
		cql.WriteString("UNLOGGED ")
	}
	if b.counter {
		cql.WriteString("COUNTER ")
	}
	cql.WriteString("BATCH ")

	names = append(names, b.using.writeCql(&cql)...)

	for _, stmt := range b.stmts {
		cql.WriteString(stmt)
		cql.WriteByte(';')
		cql.WriteByte(' ')
	}
	names = append(names, b.names...)

	cql.WriteString("APPLY BATCH ")

	stmt = cql.String()
	return
}

// Add builds the builder and adds the statement to the batch.
func (b *BatchBuilder) Add(builder Builder) *BatchBuilder {
	return b.AddStmt(builder.ToCql())
}

// AddStmt adds statement to the batch.
func (b *BatchBuilder) AddStmt(stmt string, names []string) *BatchBuilder {
	b.stmts = append(b.stmts, stmt)
	b.names = append(b.names, names...)
	return b
}

// AddWithPrefix builds the builder and adds the statement to the batch. Names
// are prefixed with the prefix + ".".
func (b *BatchBuilder) AddWithPrefix(prefix string, builder Builder) *BatchBuilder {
	stmt, names := builder.ToCql()
	return b.AddStmtWithPrefix(prefix, stmt, names)
}

// AddStmtWithPrefix adds statement to the batch. Names are prefixed with
// the prefix + ".".
func (b *BatchBuilder) AddStmtWithPrefix(prefix, stmt string, names []string) *BatchBuilder {
	b.stmts = append(b.stmts, stmt)
	for _, name := range names {
		if prefix != "" {
			name = fmt.Sprint(prefix, ".", name)
		}
		b.names = append(b.names, name)
	}
	return b
}

// UnLogged sets a UNLOGGED BATCH clause on the query.
func (b *BatchBuilder) UnLogged() *BatchBuilder {
	b.unlogged = true
	return b
}

// Counter sets a COUNTER BATCH clause on the query.
func (b *BatchBuilder) Counter() *BatchBuilder {
	b.counter = true
	return b
}

// TTL adds USING TTL clause to the query.
func (b *BatchBuilder) TTL(d time.Duration) *BatchBuilder {
	b.using.TTL(d)
	return b
}

// TTLNamed adds USING TTL clause to the query with a custom parameter name.
func (b *BatchBuilder) TTLNamed(name string) *BatchBuilder {
	b.using.TTLNamed(name)
	return b
}

// Timestamp adds USING TIMESTAMP clause to the query.
func (b *BatchBuilder) Timestamp(t time.Time) *BatchBuilder {
	b.using.Timestamp(t)
	return b
}

// TimestampNamed adds a USING TIMESTAMP clause to the query with a custom
// parameter name.
func (b *BatchBuilder) TimestampNamed(name string) *BatchBuilder {
	b.using.TimestampNamed(name)
	return b
}
