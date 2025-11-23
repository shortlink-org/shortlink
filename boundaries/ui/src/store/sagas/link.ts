import { put, takeLatest, call } from 'redux-saga/effects'

import * as t from '@/store/types'

import API from './api'

function* fetchLinkById(action: { payload: string }) {
  try {
    // @ts-ignore
    const response = yield call(API.links.getLink, action.payload)

    yield put({
      type: t.LINK_FETCH_SUCCEEDED,
      payload: response.data.link,
    })
  } catch (error: any) {
    yield put({
      type: t.LINK_FETCH_FAILED,
      payload: error.message || 'Failed to fetch link',
    })
  }
}

function* watchFetchLinkById() {
  yield takeLatest(t.LINK_FETCH_REQUESTED, fetchLinkById)
}

function* fetchLinkList() {
  try {
    // @ts-ignore
    const response = yield call(API.links.listLinks)

    yield put({
      type: t.LINK_FETCH_LIST_SUCCEEDED,
      payload: response.data,
    })
  } catch (error: any) {
    yield put({
      type: t.LINK_FETCH_LIST_FAILED,
      payload: error.message || 'Failed to fetch links',
    })
  }
}

function* watchFetchLinkList() {
  yield takeLatest(t.LINK_FETCH_LIST_REQUESTED, fetchLinkList)
}

function* addLink(action: { payload: any }) {
  try {
    // @ts-ignore
    const response = yield call(API.links.addLink, action.payload)

    yield put({
      type: t.LINK_ADD_SUCCEEDED,
      payload: response.data.link,
    })
  } catch (error: any) {
    yield put({
      type: t.LINK_ADD_FAILED,
      payload: error.message || 'Failed to add link',
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
    yield call(API.links.deleteLink, action.payload)

    yield put({
      type: t.LINK_DELETE_SUCCEEDED,
      payload: action.payload, // hash
    })
  } catch (error: any) {
    yield put({
      type: t.LINK_DELETE_FAILED,
      payload: error.message || 'Failed to delete link',
    })
  }
}

function* watchDeleteLink() {
  // @ts-ignore
  yield takeLatest(t.LINK_DELETE_REQUESTED, deleteLink)
}

function* updateLink(action: { payload: any }) {
  try {
    // @ts-ignore
    const response = yield call(API.links.updateLink, action.payload.hash, { link: action.payload })

    yield put({
      type: t.LINK_UPDATE_SUCCEEDED,
      payload: response.data.link,
    })
  } catch (error: any) {
    yield put({
      type: t.LINK_UPDATE_FAILED,
      payload: error.message || 'Failed to update link',
    })
  }
}

function* watchUpdateLink() {
  yield takeLatest(t.LINK_UPDATE_REQUESTED, updateLink)
}

export default [watchFetchLinkById(), watchFetchLinkList(), watchAddLink(), watchDeleteLink(), watchUpdateLink()]
