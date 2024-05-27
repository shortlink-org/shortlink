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
  + [In Search of an Understandable Consensus Algorithm (Extended Version)](https://raft.github.io/raft.pdf)
+ [Implementing Raft](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/)
+ [Raft: Understandable Distributed Consensus](https://thesecretlivesofdata.com/raft/)
+ [GopherCon 2023: Philip O'Toole - Build Your Own Distributed System Using Go](https://youtu.be/8XbxQ1Epi5w?si=BHtRY59yORQrGoyt)
