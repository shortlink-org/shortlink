import { HYDRATE } from "next-redux-wrapper";
import * as t from "../types";

const initialState = {
	list: [],
};

const mainReducer = (state = initialState, action) => {
  switch (action.type) {
    case HYDRATE:
      return {...state, ...action.payload.link}
    case t.LINK_FETCH_SUCCEEDED:
      return {
        ...state,
        list: [
          ...state.list,
          action.payload,
        ]
      }
    case t.LINK_FETCH_LIST_SUCCEEDED:
      return {
        ...state,
        list: action.payload.link
      }
    case t.LINK_ADD_SUCCEEDED:
      return {
        ...state,
        list: [
          ...state.list,
          action.payload,
        ]
      }
    case t.LINK_UPDATE_SUCCEEDED:
      return {
        ...state,
        list: [
          ...state.list,
          action.payload,
        ]
      }
    case t.LINK_DELETE_SUCCEEDED:
      return {
        ...state,
        list: state.list.filter(item => action.payload.hash !== item.hash),
      }
    default:
      return state
  }
}

export default mainReducer;
