import { put, takeLatest } from 'redux-saga/effects'
import * as t from 'store/types'

// @ts-ignore
function* fetchLinkById(id) {
  try {
    // @ts-ignore
    const response = yield fetch(`/api/link/${id}`)

    // @ts-ignore
    const link = yield response.json()

    yield put({
      type: t.LINK_FETCH_SUCCEEDED,
      payload: link.data,
    })
  } catch (error) {
    yield put({
      type: t.LINK_FETCH_FAILED,
      payload: error.message,
    })
  }
}

function* watchFetchLinkById() {
  yield takeLatest(t.LINK_FETCH_REQUESTED, fetchLinkById)
}

function* fetchLinkList() {
  try {
    // @ts-ignore
    const response = yield fetch(`/api/links`)

    // @ts-ignore
    const links = yield response.json()

    yield put({
      type: t.LINK_FETCH_LIST_SUCCEEDED,
      payload: links,
    })
  } catch (error) {
    yield put({
      type: t.LINK_FETCH_LIST_FAILED,
      payload: error.message,
    })
  }
}

function* watchFetchLinkList() {
  yield takeLatest(t.LINK_FETCH_LIST_REQUESTED, fetchLinkList)
}

function* addLink(action: { payload: any }) {
  try {
    // @ts-ignore
    const response = yield fetch('/api/link', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(action.payload),
    })

    // @ts-ignore
    const newLink = yield response.json()

    yield put({
      type: t.LINK_ADD_SUCCEEDED,
      payload: newLink.data,
    })
  } catch (error) {
    yield put({
      type: t.LINK_ADD_FAILED,
      payload: error.message,
    })
  }
}

function* watchAddLink() {
  // @ts-ignore
  yield takeLatest(t.LINK_ADD_REQUESTED, addLink)
}

function* deleteLink(action: { payload: any }) {
  try {
    // @ts-ignore
    const response = yield fetch(`/api/link/${action.payload}`, {
      method: 'DELETE',
    })

    // @ts-ignore
    const deletedLink = yield response.json()

    yield put({
      type: t.LINK_DELETE_SUCCEEDED,
      payload: deletedLink.data.id,
    })
  } catch (error) {
    yield put({
      type: t.LINK_DELETE_FAILED,
      payload: error.message,
    })
  }
}

function* watchDeleteLink() {
  // @ts-ignore
  yield takeLatest(t.LINK_DELETE_REQUESTED, deleteLink)
}

// @ts-ignore
function* updateLink(action) {
  try {
    // @ts-ignore
    const response = yield fetch(`/api/link/${action.payload._id}`, {
      // eslint-disable-line
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(action.payload),
    })

    // @ts-ignore
    const updatedLink = yield response.json()

    yield put({
      type: t.LINK_UPDATE_SUCCEEDED,
      payload: updatedLink.data,
    })
  } catch (error) {
    yield put({
      type: t.LINK_UPDATE_FAILED,
      payload: error.message,
    })
  }
}

function* watchUpdateLink() {
  yield takeLatest(t.LINK_UPDATE_REQUESTED, updateLink)
}

export default [
  watchFetchLinkById(),
  watchFetchLinkList(),
  watchAddLink(),
  watchDeleteLink(),
  watchUpdateLink(),
]
