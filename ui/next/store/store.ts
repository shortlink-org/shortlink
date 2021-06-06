import { createStore, applyMiddleware } from 'redux';
import { createWrapper } from 'next-redux-wrapper';
import { composeWithDevTools } from 'redux-devtools-extension';
import createMiddleware from 'redux-saga';

import rootReducer from './reducers';
import rootSaga from './sagas';

// Create the middleware
const sagaMiddleware = createMiddleware();

// Add an extra parameter for applying middleware:
const store = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(sagaMiddleware)),
);

// Run your sagas on server
sagaMiddleware.run(rootSaga);

const makeStore = () => store;

export const wrapper = createWrapper(makeStore);
