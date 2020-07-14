package rethinkdb

import (
	p "gopkg.in/rethinkdb/rethinkdb-go.v6/ql2"
)

// TableCreateOpts contains the optional arguments for the TableCreate term
type TableCreateOpts struct {
	PrimaryKey           interface{} `rethinkdb:"primary_key,omitempty"`
	Durability           interface{} `rethinkdb:"durability,omitempty"`
	Shards               interface{} `rethinkdb:"shards,omitempty"`
	Replicas             interface{} `rethinkdb:"replicas,omitempty"`
	PrimaryReplicaTag    interface{} `rethinkdb:"primary_replica_tag,omitempty"`
	NonVotingReplicaTags interface{} `rethinkdb:"nonvoting_replica_tags,omitempty"`
}

func (o TableCreateOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// TableCreate creates a table. A RethinkDB table is a collection of JSON
// documents.
//
// Note: Only alphanumeric characters and underscores are valid for the table name.
func TableCreate(name interface{}, optArgs ...TableCreateOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructRootTerm("TableCreate", p.Term_TABLE_CREATE, []interface{}{name}, opts)
}

// TableCreate creates a table. A RethinkDB table is a collection of JSON
// documents.
//
// Note: Only alphanumeric characters and underscores are valid for the table name.
func (t Term) TableCreate(name interface{}, optArgs ...TableCreateOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructMethodTerm(t, "TableCreate", p.Term_TABLE_CREATE, []interface{}{name}, opts)
}

// TableDrop deletes a table. The table and all its data will be deleted.
func TableDrop(args ...interface{}) Term {
	return constructRootTerm("TableDrop", p.Term_TABLE_DROP, args, map[string]interface{}{})
}

// TableDrop deletes a table. The table and all its data will be deleted.
func (t Term) TableDrop(args ...interface{}) Term {
	return constructMethodTerm(t, "TableDrop", p.Term_TABLE_DROP, args, map[string]interface{}{})
}

// TableList lists all table names in a database.
func TableList(args ...interface{}) Term {
	return constructRootTerm("TableList", p.Term_TABLE_LIST, args, map[string]interface{}{})
}

// TableList lists all table names in a database.
func (t Term) TableList(args ...interface{}) Term {
	return constructMethodTerm(t, "TableList", p.Term_TABLE_LIST, args, map[string]interface{}{})
}

// IndexCreateOpts contains the optional arguments for the IndexCreate term
type IndexCreateOpts struct {
	Multi interface{} `rethinkdb:"multi,omitempty"`
	Geo   interface{} `rethinkdb:"geo,omitempty"`
}

func (o IndexCreateOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// IndexCreate creates a new secondary index on a table. Secondary indexes
// improve the speed of many read queries at the slight cost of increased
// storage space and decreased write performance.
//
// IndexCreate supports the creation of the following types of indexes, to create
// indexes using arbitrary expressions use IndexCreateFunc.
//   - Simple indexes based on the value of a single field.
//   - Geospatial indexes based on indexes of geometry objects, created when the
//     geo optional argument is true.
func (t Term) IndexCreate(name interface{}, optArgs ...IndexCreateOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructMethodTerm(t, "IndexCreate", p.Term_INDEX_CREATE, []interface{}{name}, opts)
}

// IndexCreateFunc creates a new secondary index on a table. Secondary indexes
// improve the speed of many read queries at the slight cost of increased
// storage space and decreased write performance. The function takes a index
// name and RQL term as the index value , the term can be an anonymous function
// or a binary representation obtained from the function field of indexStatus.
//
// It supports the creation of the following types of indexes.
//   - Simple indexes based on the value of a single field where the index has a
//     different name to the field.
//   - Compound indexes based on multiple fields.
//   - Multi indexes based on arrays of values, created when the multi optional argument is true.
func (t Term) IndexCreateFunc(name, indexFunction interface{}, optArgs ...IndexCreateOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructMethodTerm(t, "IndexCreate", p.Term_INDEX_CREATE, []interface{}{name, funcWrap(indexFunction)}, opts)
}

// IndexDrop deletes a previously created secondary index of a table.
func (t Term) IndexDrop(args ...interface{}) Term {
	return constructMethodTerm(t, "IndexDrop", p.Term_INDEX_DROP, args, map[string]interface{}{})
}

// IndexList lists all the secondary indexes of a table.
func (t Term) IndexList(args ...interface{}) Term {
	return constructMethodTerm(t, "IndexList", p.Term_INDEX_LIST, args, map[string]interface{}{})
}

// IndexRenameOpts contains the optional arguments for the IndexRename term
type IndexRenameOpts struct {
	Overwrite interface{} `rethinkdb:"overwrite,omitempty"`
}

func (o IndexRenameOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// IndexRename renames an existing secondary index on a table.
func (t Term) IndexRename(oldName, newName interface{}, optArgs ...IndexRenameOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructMethodTerm(t, "IndexRename", p.Term_INDEX_RENAME, []interface{}{oldName, newName}, opts)
}

// IndexStatus gets the status of the specified indexes on this table, or the
// status of all indexes on this table if no indexes are specified.
func (t Term) IndexStatus(args ...interface{}) Term {
	return constructMethodTerm(t, "IndexStatus", p.Term_INDEX_STATUS, args, map[string]interface{}{})
}

// IndexWait waits for the specified indexes on this table to be ready, or for
// all indexes on this table to be ready if no indexes are specified.
func (t Term) IndexWait(args ...interface{}) Term {
	return constructMethodTerm(t, "IndexWait", p.Term_INDEX_WAIT, args, map[string]interface{}{})
}

// ChangesOpts contains the optional arguments for the Changes term
type ChangesOpts struct {
	Squash              interface{} `rethinkdb:"squash,omitempty"`
	IncludeInitial      interface{} `rethinkdb:"include_initial,omitempty"`
	IncludeStates       interface{} `rethinkdb:"include_states,omitempty"`
	IncludeOffsets      interface{} `rethinkdb:"include_offsets,omitempty"`
	IncludeTypes        interface{} `rethinkdb:"include_types,omitempty"`
	ChangefeedQueueSize interface{} `rethinkdb:"changefeed_queue_size,omitempty"`
}

// ChangesOpts contains the optional arguments for the Changes term
func (o ChangesOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// Changes returns an infinite stream of objects representing changes to a query.
func (t Term) Changes(optArgs ...ChangesOpts) Term {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return constructMethodTerm(t, "Changes", p.Term_CHANGES, []interface{}{}, opts)
}
