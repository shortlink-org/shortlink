## Specification pattern

This package provides a simple implementation of the specification pattern. 

**The specification pattern** is a software design pattern that allows us to define a business rule that 
can be applied to a set of objects.

The pattern is useful when we need to filter a collection of objects based on a set of rules.

### How to use

You can see an example by [link](./example/example.go).

```shell
# Example print:

$> User Alice satisfies the specification
$> User Bob does not satisfy the specification: Specification failed: User Bob's name does not start with 'A'
$> User Charlie does not satisfy the specification: Specification failed: User Charlie's name does not start with 'A'
$> Filtered users: [0xc0001080c0]
```

### References

> [!TIP]
> 
> - [Start Using The Specification Pattern In Golang Now!](https://thegodev.com/specification-pattern/) - Tutorial on how to use the specification pattern in Go.
> - [Wikipedia](https://en.wikipedia.org/wiki/Specification_pattern) - Specification pattern on Wikipedia.
