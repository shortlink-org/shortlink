import * as t from "../types";

export const fetchLinkById = (id) => {
  return {
    type: t.LINK_FETCH_REQUESTED,
    payload: id,
  }
}

export const fetchLinkList = () => {
  return {
    type: t.LINK_FETCH_LIST_REQUESTED,
  }
}

export const addLink = (link) => {
  return {
    type: t.LINK_ADD_REQUESTED,
    payload: link,
  }
}

export const updateLinkById = (link) => {
  return {
    type: t.LINK_UPDATE_REQUESTED,
    payload: link,
  }
}

export const deleteLinkById = (link) => {
  return {
    type: t.LINK_DELETE_REQUESTED,
    payload: link,
  }
}
