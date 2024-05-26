## Raft

### Getting Started

```go
package main

import (
	"github.com/shortlink-org/shortlink/pkg/raft"
)

func main() {
	r, err := raft.New(raft.Config{
		Name:     "node1",
		BindAddr: "localhost:8001",
	})
	if err != nil {
		panic(err)
	}

	r.Go(func() {
		// Do something
	})
}
```

### Raft server states

```mermaid
stateDiagram-v2
  UNSPECIFIED --> Follower: times out, start election
  Follower    --> Candidate: times out, start election
  Candidate   --> Candidate: timeout, new election
  Candidate   --> Leader: win election
  Candidate   --> Follower: receive vote from majority
```

#### Docs

+ [The Raft Consensus Algorithm](https://raft.github.io/)
+ [Implementing Raft](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/)
