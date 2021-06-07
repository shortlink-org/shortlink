import * as t from '../types'

export const loginAuth = (login: any) => ({
  type: t.AUTH_LOGIN_REQUESTED,
  payload: login,
})
