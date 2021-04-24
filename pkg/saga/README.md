### Saga manager

#### Saga steps of state:

+ START (REJECT)
+ SUCCESS (FAIL)
+ DONE

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
  saga.AddStep(&saga.Step{
    Name: "Save link in store",
    Needs: []string{},
    Func: func(context.Context) error {
    	err := l.Store.Add(link)
    	return err
    },
    CompenstateFunc: func(context.Context) error {
      err := l.Store.Delete(link)
    	return err
    }
  })

  // step: get metadata
  saga.AddStep(&saga.Step{
    Name: "Get metadata by link",
    Needs: []string{},
    Func: func(context.Context, link link.Link) error {
    	err := l.MetadataServer.Add(link)
    	return err
    },
    CompenstateFunc: func(context.Context) error {
    	err := l.MetadataServer.Delete(link)
    	return err
    }
  })
  
  // step: send notify
  saga.AddStep(&saga.Step{
    Name: "send notify",
    Needs: []string{"Save link in store", "Get metadata by link"},
    Func: func(context.Context, link link.Link) error {
    	err := l.NotifyServer.Add(link)
    	return err
    },
    CompenstateFunc: func(context.Context) error {
    	return nil
    }
  })
  
  // Run saga
  err := sagaAddLink.Play()
  return err
}
```

### Ref

- [go-saga](https://github.com/itimofeev/go-saga) - example go-library
- [DAG][https://github.com/goombaio/dag] - for build pipeline steps
