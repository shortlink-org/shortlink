import { HYDRATE } from 'next-redux-wrapper'
import * as t from 'store/types'

const initialState = {
  session: {},
}

// @ts-ignore
const mainReducer = (state = initialState, action) => {
  switch (action.type) {
    case HYDRATE:
      return { ...state, ...action.payload.session }
    case t.AUTH_LOGIN_SUCCEEDED:
      return {
        ...state,
        // @ts-ignore
        list: [...state.session, action.payload],
      }
    default:
      return state
  }
}

export default mainReducer
