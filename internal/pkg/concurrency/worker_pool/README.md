## Worker pool

```mermaid
graph LR

subgraph Worker Pool
W1-->Task1
W2-->Task2
W3-->Task3
W4-->Task4
W5-->Task5
end
 
Task1-->Done
Task2-->Done
Task3-->Done
Task4-->Done
Task5-->Done
Done-->Result
```

#### TODO:

- [ ] Use `sync.Pool` or `sync.Cond` to manage workers
- [ ] Add Graceful Shutdown

#### References

**Looking for Alternatives? Peek Here:**

- https://github.com/panjf2000/ants
