import { combineReducers } from "redux";
import linkReducer from "./link";

const rootReducer = combineReducers({
	link: linkReducer,
});

export default rootReducer;
