import * as t from '../types'

export const getSession = (token: string) => ({
  type: t.SESSION_FETCH_REQUESTED,
  payload: token,
})

export const loginAuth = (login: any) => ({
  type: t.SESSION_CREATE_REQUESTED,
  payload: login,
})
