import { put, takeLatest } from 'redux-saga/effects'
// @ts-ignore
import {LoginRequest, PublicApi} from "@ory/kratos-client"
import * as t from 'store/types'

// Init Kratos API
// @ts-ignore
const kratos = new PublicApi(process.env.KRATOS_API || 'http://127.0.0.1:4433/.ory/kratos/public')

function* fetchSession(token: string) {
  try {
    // @ts-ignore
    const response = yield kratos.toSession(token)

    // @ts-ignore
    const link = yield response.json()

    yield put({
      type: t.SESSION_FETCH_SUCCEEDED,
      payload: link.data,
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

function* loginAuth(action: { payload: any }) {
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
  yield takeLatest(t.AUTH_LOGIN_REQUESTED, loginAuth)
}

export default [
  watchFetchSession(),
  watchLoginAuth(),
]
