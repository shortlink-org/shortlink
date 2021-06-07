import {all} from "redux-saga/effects";

import link from 'store/sagas/link'
import auth from 'store/sagas/auth'

export default function* rootSaga() {
  yield all([
      ...link,
    ...auth
  ])
}
