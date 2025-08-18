import { all } from 'redux-saga/effects'

import link from '@/store/sagas/link'

export default function* rootSaga() {
  yield all([...link])
}
