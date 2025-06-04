import * as t from 'store/types'

const initialState = {
  list: [],
}

// @ts-ignore
const mainReducer = (state = initialState, action) => {
  switch (action.type) {
    case t.LINK_FETCH_SUCCEEDED:
      return {
        ...state,
        list: [...state.list, action.payload],
      }
    case t.LINK_FETCH_LIST_SUCCEEDED:
      return {
        ...state,
        list: action.payload.links,
      }
    case t.LINK_ADD_SUCCEEDED:
      return {
        ...state,
        list: [...state.list, action.payload],
      }
    case t.LINK_UPDATE_SUCCEEDED:
      return {
        ...state,
        list: [...state.list, action.payload],
      }
    case t.LINK_DELETE_SUCCEEDED:
      return {
        ...state,
        // @ts-ignore
        list: state.list.filter((item) => action.payload.hash !== item.hash),
      }
    default:
      return state
  }
}

export default mainReducer
