import { put, takeLatest } from 'redux-saga/effects'
// @ts-ignore
import { LoginRequest } from "@ory/kratos-client"
import { initialiseRequest } from "./kratos/kratos"
import * as t from 'store/types'

function* loginAuth(action: { payload: any }) {
  try {
    const request = initialiseRequest({ type: "login" }) as Promise<LoginRequest>
    // @ts-ignore
    const response = yield request
    console.warn('TEST', response)

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
  watchLoginAuth(),
]
