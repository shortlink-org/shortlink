import { HYDRATE } from "next-redux-wrapper";
import * as t from "../types";

const DEMO_SET = {
  url: 'http://google.com',
  hash: '4535345',
  describe: 'Test URL for table',
  created_at: 1243432434,
  updated_at: 1243432434,
}

const initialState = {
	list: {
	  [DEMO_SET.hash]: DEMO_SET,
  },
};

const mainReducer = (state = initialState, action) => {
  switch (action.type) {
    case HYDRATE:
      return {...state, ...action.payload}
    case t.LINK_FETCH_SUCCEEDED:
      return {
        ...state,
        list: {
          ...state.list,
          [action.payload.hash]: action.payload,
        }
      }
    case t.LINK_FETCH_LIST_SUCCEEDED:
      return {
        ...state,
        list: action.payload
      }
    case t.LINK_ADD_SUCCEEDED:
      return {
        ...state,
        list: {
          ...state.list,
          [action.payload.hash]: action.payload,
        }
      }
    case t.LINK_UPDATE_SUCCEEDED:
      return {
        ...state,
        list: {
          ...state.list,
          [action.payload.hash]: action.payload,
        }
      }
    case t.LINK_DELETE_SUCCEEDED:
      return {
        ...state,
        list: {
          ...state.list,
          [action.payload.hash]: undefined,
        }
      }
    default:
      return state
  }
}

export default mainReducer;
