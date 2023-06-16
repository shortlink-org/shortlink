import { createWrapper, Context } from 'next-redux-wrapper'
import { createStore, applyMiddleware, Store } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import createSagaMiddleware, { Task } from 'redux-saga'

import rootReducer from './reducers'
import rootSaga from './sagas'

export interface SagaStore extends Store {
  sagaTask?: Task
}

const makeStore = (context: Context) => {
  // Create the middleware
  const sagaMiddleware = createSagaMiddleware()

  // Add an extra parameter for applying middleware:
  const store = createStore(
    rootReducer,
    composeWithDevTools(applyMiddleware(sagaMiddleware)),
  )

  // Run your sagas on server
  sagaMiddleware.run(rootSaga)

  return store
}

// export an assembled wrapper
export const wrapper = createWrapper<Store<any>>(makeStore, { debug: false })
