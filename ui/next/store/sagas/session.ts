import { put, takeLatest } from 'redux-saga/effects'
// @ts-ignore
import { LoginRequest, PublicApi } from '@ory/kratos-client' // eslint-disable-line
import * as t from 'store/types'
import { SESSION_FETCH_REQUESTED } from 'store/types' // eslint-disable-line

// Init Kratos API
const KRATOS_PUBLIC_API = process.env.KRATOS_API || 'http://127.0.0.1:4433'

function* fetchSession() {
  try {
    // @ts-ignore
    const response = yield fetch(`${KRATOS_PUBLIC_API}/sessions/whoami`)

    // @ts-ignore
    const session = yield response.json()

    yield put({
      type: t.SESSION_FETCH_SUCCEEDED,
      payload: session,
    })
  } catch (error) {
    yield put({
      type: t.SESSION_FETCH_FAILED,
      payload: error.message,
    })
  }
}

function* watchFetchSession() {
  // @ts-ignore
  yield takeLatest(t.SESSION_FETCH_REQUESTED, fetchSession)
}

function* loginAuth(action: { payload: any }) { // eslint-disable-line
  try {
    // const request = initialiseRequest({ type: "login" }) as Promise<LoginRequest>
    // // @ts-ignore
    // const response = yield request
    // console.warn('TEST', response)
    // // @ts-ignore
    // const response = yield fetch('/api/link', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json',
    //   },
    //   body: JSON.stringify(action.payload),
    // })
    //
    // // @ts-ignore
    // const newSession = yield response.json()
    //
    // yield put({
    //   type: t.AUTH_LOGIN_SUCCEEDED,
    //   payload: newSession.data,
    // })
  } catch (error) {
    yield put({
      type: t.LINK_ADD_FAILED,
      payload: error.message,
    })
  }
}

function* watchLoginAuth() {
  // @ts-ignore
  yield takeLatest(t.SESSION_FETCH_REQUESTED, loginAuth)
}

export default [watchFetchSession(), watchLoginAuth()]
