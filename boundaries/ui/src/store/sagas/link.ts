import { put, takeLatest } from 'redux-saga/effects'

import * as t from '@/store/types'

import API from './api'

// @ts-ignore
function* fetchLinkById(id) {
  try {
    // @ts-ignore
    const response = yield API.links.getLink(id)

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
    const response = yield API.links.listLinks()

    yield put({
      type: t.LINK_FETCH_LIST_SUCCEEDED,
      payload: response.data,
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
    yield API.links.addLink(action.payload)

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

function* deleteLink(action: { payload: string }) {
  try {
    // @ts-ignore
    const response = yield API.links.deleteLink(action.payload)

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
    const response = yield API.links.updateLink(action.payload._id, action.payload)

    yield put({
      type: t.LINK_UPDATE_SUCCEEDED,
      payload: response.data,
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

export default [watchFetchLinkById(), watchFetchLinkList(), watchAddLink(), watchDeleteLink(), watchUpdateLink()]
