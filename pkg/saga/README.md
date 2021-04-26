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
	const SAGA_NAME = "Add link"
	const SAGA_STEP_SAVE_LINK = "Save link in store"
	const SAGA_STEP_GET_METADATA = "Get metadata by link"

  // create a new saga for add link
  sagaAddLink, errs := saga.
    New(SAGA_NAME).   // name saga
    WithContext(ctx). // ctx for tracing
    SetStore(store).  // create store for save state saga
    Build()
  
  if len(errs) > 0 {
    // check err...
  }
  
  // step: save to store
  _, err = saga.AddStep(SAGA_STEP_SAVE_LINK).
    Then(func(context.Context) error {
      err := l.Store.Add(link)
      return err
    }).
    Reject(func(context.Context) error {
      err := l.Store.Delete(link)
      return err
    }).
  	Build()
  
  if len(errs) > 0 {
    // check err...
  }

  // step: get metadata
  saga.AddStep(SAGA_STEP_GET_METADATA).
    Then(addFunc).
    Reject(cancelAddFunc)

  // step: send notify
  saga.AddStep("send notify").
    Needs([]string{SAGA_STEP_SAVE_LINK, SAGA_STEP_GET_METADATA}).
    Then(youNotifyFunc)
  
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

### Alternative solve

#### Use `Step.degree` increment

Run jobs by `degree`: `0 -> 1 -> 2 -> ... -> N`  
Reject jobs by `degree`: `Current -> Current - 1 -> Current - 2 -> ... -> 0`
