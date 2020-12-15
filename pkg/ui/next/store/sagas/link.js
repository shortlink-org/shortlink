import { all, put, takeLatest } from "redux-saga/effects";
import * as t from "../types";

function* fetchLinkById(id) {
	try {
		const response = yield fetch(`/link/${id}`);

		const link = yield response.json();

		yield put({
			type: t.LINK_FETCH_SUCCEEDED,
			payload: link.data,
		});
	} catch (error) {
		yield put({
			type: t.LINK_FETCH_FAILED,
			payload: error.message,
		});
	}
}

function* watchFetchLinkById(id) {
	yield takeLatest(t.LINK_FETCH_REQUESTED, fetchLinkById);
}

function* fetchLinkList() {
	try {
		const response = yield fetch(`/link`);

		const links = yield response.json();

		yield put({
			type: t.LINK_FETCH_LIST_SUCCEEDED,
			payload: links.data,
		});
	} catch (error) {
		yield put({
			type: t.LINK_FETCH_LIST_FAILED,
			payload: error.message,
		});
	}
}

function* watchFetchLinkList() {
	yield takeLatest(t.LINK_FETCH_LIST_REQUESTED, fetchLinkList);
}

function* addLink(action) {
	try {
		const response = yield fetch("/link", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(action.payload),
		});

		const newLink = yield response.json();

		yield put({
			type: t.LINK_ADD_SUCCEEDED,
			payload: newLink.data,
		});
	} catch (error) {
		yield put({
			type: t.LINK_ADD_FAILED,
			payload: error.message,
		});
	}
}

function* watchAddLink() {
	yield takeLatest(t.LINK_ADD_REQUESTED, addLink);
}

function* deleteLink(action) {
	try {
		const response = yield fetch(`/link/${action.payload}`, {
			method: "DELETE",
		});

		const deletedLink = yield response.json();

		yield put({
			type: t.LINK_DELETE_SUCCEEDED,
			payload: deletedLink.data.id,
		});
	} catch (error) {
		yield put({
			type: t.LINK_DELETE_FAILED,
			payload: error.message,
		});
	}
}

function* watchDeleteLink() {
	yield takeLatest(t.LINK_DELETE_REQUESTED, deleteLink);
}

function* updateLink(action) {
	try {
		const response = yield fetch(`/link/${action.payload._id}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(action.payload),
		});

		const updatedLink = yield response.json();

		yield put({
			type: t.LINK_UPDATE_SUCCEEDED,
			payload: updatedLink.data,
		});
	} catch (error) {
		yield put({
			type: t.LINK_UPDATE_FAILED,
			payload: error.message,
		});
	}
}

function* watchUpdateLink() {
	yield takeLatest(t.LINK_UPDATE_REQUESTED, updateLink);
}

export default function* rootSaga() {
	yield all([
	  watchFetchLinkById(),
		watchFetchLinkList(),
		watchAddLink(),
		watchDeleteLink(),
		watchUpdateLink(),
	]);
}
