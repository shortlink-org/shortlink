'use client'

import { configureStore } from '@reduxjs/toolkit'

import rootReducer from './reducers'

export const store = configureStore({
  reducer: rootReducer,
  devTools: true,
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

import { createStore, applyMiddleware, Store } from 'redux'
import { createWrapper, Context } from 'next-redux-wrapper'
import { composeWithDevTools } from 'redux-devtools-extension'
import createSagaMiddleware, { Task } from 'redux-saga'

// import rootSaga from './sagas'
//
// export interface SagaStore extends Store {
//   sagaTask?: Task
// }
//
// const makeStore = (context: Context) => {
//   // Create the middleware
//   const sagaMiddleware = createSagaMiddleware()
//
//   // Add an extra parameter for applying middleware:
//   let store = createStore(
//     rootReducer,
//     composeWithDevTools(applyMiddleware(sagaMiddleware)),
//   )
//
//   // Run your sagas on server
//   sagaMiddleware.run(rootSaga)
//
//   return store
// }
//
// // export an assembled wrapper
// export const wrapper = createWrapper<Store<any>>(makeStore, { debug: false })
