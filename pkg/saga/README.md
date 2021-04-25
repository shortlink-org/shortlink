### Saga manager

#### Saga steps of state:

+ WAIT -> START -> RUN -> READY
+ REJECT -> FAIL

#### example

```go
import (
	"saga"
)

func (l *linkUseCase) addLinkSaga(ctx, link link.Link) error {
	// create a new saga for add link
  sagaAddLink := saga.
  	New("add link"). // name saga
  	WithContext(ctx) // ctx for tracing
  	Store(store)     // create store for save state saga
  
  // step: save to store
  saga.AddStep("Save link in store").
  	Do(func(context.Context) error {
    	err := l.Store.Add(link)
    	return err
    }).
  	Reject(func(context.Context) error {
      err := l.Store.Delete(link)
    	return err
    })

  // step: get metadata
  saga.AddStep("Get metadata by link").
  	Do(func(context.Context, link link.Link) error {
    	err := l.MetadataServer.Add(link)
    	return err
    }).
  	Reject(func(context.Context) error {
    	err := l.MetadataServer.Delete(link)
    	return err
    })

  // step: send notify
  saga.AddStep("send notify").
  	Needs([]string{"Save link in store", "Get metadata by link"}).
  	Do(func(context.Context, link link.Link) error {
    	err := l.NotifyServer.Add(link)
    	return err
    })
  
  // Run saga
  err := sagaAddLink.Play()
  return err
}
```

### Ref

- [DAG](https://github.com/goombaio/dag) - for build pipeline steps
- libs:
  - [go-saga](https://github.com/itimofeev/go-saga) - example go-library
  - https://github.com/danielgerlag/workflow-core
