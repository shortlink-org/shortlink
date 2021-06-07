import { combineReducers } from 'redux'
import linkReducer from 'store/reducers/link'
import sessionReducer from 'store/reducers/session'

const rootReducer = combineReducers({
  link: linkReducer,
  auth: sessionReducer,
})

export default rootReducer
