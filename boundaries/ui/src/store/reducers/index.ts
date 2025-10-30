import linkReducer from '@/store/reducers/link'
import sessionReducer from '@/store/reducers/session'

const rootReducer = {
  link: linkReducer,
  session: sessionReducer,
}

export default rootReducer
