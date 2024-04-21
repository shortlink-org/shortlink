## value-object plugin

> [!NOTE]
> This `golangci-lint` plugin ensures that setter methods in value objects are private, 
> supporting immutability and data integrity in Go applications.

### Getting Started

We use Makefile for build and deploy.

```bash
make help # show help message with all commands and targets
```

### Example Usage

#### Bad Code

```go
package example

type ValueObject struct {
    value int
}

// Public setter method - not recommended
func (v *ValueObject) SetValue(val int) {
    v.value = val
}
```

#### Good Code

```go
package example

type ValueObject struct {
    value int
}

// Private setter method - follows best practices
func (v *ValueObject) setValue(val int) {
    v.value = val
}
```
