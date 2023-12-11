import { createWrapper, Context } from 'next-redux-wrapper'
import { Store } from 'redux'
import createSagaMiddleware, { Task } from 'redux-saga'
import { configureStore } from '@reduxjs/toolkit'

import rootReducer from './reducers'
import rootSaga from './sagas'

export interface SagaStore extends Store {
  sagaTask?: Task
}

const makeStore = (context: Context) => {
  // Create the middleware
  const sagaMiddleware = createSagaMiddleware()

  const middleware = [sagaMiddleware]

  // Mount it on the Store
  const store = configureStore({
    reducer: rootReducer,
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware().concat(middleware),
  })

  // Run your sagas on server
  sagaMiddleware.run(rootSaga)

  return store
}

// export an assembled wrapper
export const wrapper = createWrapper<Store<any>>(makeStore, { debug: false })
