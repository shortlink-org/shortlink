## go_orm: A Minimalistic Go ORM Prototype

`go_orm` is an illustrative Object-Relational Mapping (ORM) package designed for Go. 
This ORM is generated using the Protocol Buffers (Protobuf) language, 
providing a mechanism to automatically create code from structured data definitions.

### How It Works

1. **Protobuf Definitions:** The base of this ORM is the `link.proto` file, which contains data structure definitions. 
The use of Protobuf ensures consistency and structured data definitions that can be easily serialized and 
transmitted over the wire.

2. **Code Generation:** Using the `protoc` tool along with the `protoc-gen-gotemplate` plugin, 
we can generate Go code from the Protobuf definitions. The line `//go:generate` at the start 
of our `go_orm` package signifies this process.

3. **Data Access Object (DAO):** The `Store` structure in our Go package serves as the DAO, 
allowing us to connect to a SQL database and execute queries.

4. **Dynamic Query Builder:** The generated code provides a `buildFilter` function, 
which can construct SQL queries dynamically based on provided filters. This leverages the `squirrel` package, 
a fluent SQL query builder for Go.

### Conclusion

This is a simple example of how we can leverage Protobuf to generate an ORM for Go.
