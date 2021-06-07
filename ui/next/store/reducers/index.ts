import { combineReducers } from 'redux'
import linkReducer from 'store/reducers/link'
import authReducer from 'store/reducers/auth'

const rootReducer = combineReducers({
  link: linkReducer,
  auth: authReducer,
})

export default rootReducer
