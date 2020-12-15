import { createStore, applyMiddleware } from "redux";
import { createWrapper } from "next-redux-wrapper";
import { composeWithDevTools } from "redux-devtools-extension";
import createMiddleware from "redux-saga";

import rootReducer from "./reducers";
import rootSaga from "./sagas";

const sagaMiddleware = createMiddleware();

const store = createStore(
	rootReducer,
	composeWithDevTools(applyMiddleware(sagaMiddleware))
);

sagaMiddleware.run(rootSaga);

const makeStore = () => store;

export const wrapper = createWrapper(makeStore);
