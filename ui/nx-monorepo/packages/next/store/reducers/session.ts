import { HYDRATE } from 'next-redux-wrapper'

import * as t from 'store/types'

const initialState = {
  kratos: null,
}

// @ts-ignore
const mainReducer = (state = initialState, action) => {
  switch (action.type) {
    case HYDRATE:
      return { ...state, ...action.payload.session }
    case t.SESSION_FETCH_SUCCEEDED:
      return {
        ...state,
        kratos: action.payload,
      }
    case t.SESSION_CREATE_SUCCEEDED:
      return {
        ...state,
        kratos: action.payload,
      }
    default:
      return state
  }
}

export default mainReducer
