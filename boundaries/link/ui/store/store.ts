import { Store } from 'redux'
import createSagaMiddleware, { Task } from 'redux-saga'
import { configureStore } from '@reduxjs/toolkit'

import rootReducer from './reducers'
import rootSaga from './sagas'

export interface SagaStore extends Store {
  sagaTask?: Task
}

export const makeStore = () => {
  // Create the middleware
  const sagaMiddleware = createSagaMiddleware()

  const middleware = [sagaMiddleware]

  // Mount it on the Store
  const store = configureStore({
    reducer: rootReducer,
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(middleware),
  })

  // Run your sagas on server
  sagaMiddleware.run(rootSaga)

  return store
}

// Infer the type of makeStore
export type AppStore = ReturnType<typeof makeStore>
// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<AppStore['getState']>
export type AppDispatch = AppStore['dispatch']
