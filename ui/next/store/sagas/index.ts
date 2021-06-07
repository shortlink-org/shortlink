import { all } from 'redux-saga/effects'

import link from 'store/sagas/link'
import session from 'store/sagas/session'

export default function* rootSaga() {
  yield all([...session, ...link])
}
