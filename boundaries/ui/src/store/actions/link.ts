import * as t from '../types'
import { V1AddRequest } from '../api/Api'

export const fetchLinkById = (id: string) => ({
  type: t.LINK_FETCH_REQUESTED,
  payload: id,
})

export const fetchLinkList = () => ({
  type: t.LINK_FETCH_LIST_REQUESTED,
})

export const addLink = (link: V1AddRequest) => ({
  type: t.LINK_ADD_REQUESTED,
  payload: link,
})

export const updateLinkById = (link: { hash: string; [key: string]: any }) => ({
  type: t.LINK_UPDATE_REQUESTED,
  payload: link,
})

export const deleteLinkById = (hash: string) => ({
  type: t.LINK_DELETE_REQUESTED,
  payload: hash,
})
