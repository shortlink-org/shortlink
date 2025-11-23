import * as t from '@/store/types'
import { GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link } from '@/store/api/Api'

export interface LinkState {
  list: GithubComBatazorShortlinkInternalServicesLinkDomainLinkV1Link[]
  loading: boolean
  error: string | null
}

interface LinkAction {
  type: string
  payload?: any
}

const initialState: LinkState = {
  list: [],
  loading: false,
  error: null,
}

const mainReducer = (state: LinkState = initialState, action: LinkAction): LinkState => {
  switch (action.type) {
    case t.LINK_FETCH_REQUESTED:
    case t.LINK_FETCH_LIST_REQUESTED:
    case t.LINK_ADD_REQUESTED:
    case t.LINK_UPDATE_REQUESTED:
    case t.LINK_DELETE_REQUESTED:
      return {
        ...state,
        loading: true,
        error: null,
      }

    case t.LINK_FETCH_FAILED:
    case t.LINK_FETCH_LIST_FAILED:
    case t.LINK_ADD_FAILED:
    case t.LINK_UPDATE_FAILED:
    case t.LINK_DELETE_FAILED:
      return {
        ...state,
        loading: false,
        error: action.payload || 'An error occurred',
      }

    case t.LINK_FETCH_SUCCEEDED:
      return {
        ...state,
        loading: false,
        error: null,
        list: [...state.list, action.payload],
      }

    case t.LINK_FETCH_LIST_SUCCEEDED:
      // API возвращает { links: { link: [...] } } или { links: [...] }
      const linksData = action.payload?.links
      const linksList = Array.isArray(linksData) 
        ? linksData 
        : linksData?.link || []
      
      return {
        ...state,
        loading: false,
        error: null,
        list: linksList,
      }

    case t.LINK_ADD_SUCCEEDED:
      return {
        ...state,
        loading: false,
        error: null,
        list: [...state.list, action.payload],
      }

    case t.LINK_UPDATE_SUCCEEDED:
      return {
        ...state,
        loading: false,
        error: null,
        list: state.list.map((item) => (item.hash === action.payload?.hash ? action.payload : item)),
      }

    case t.LINK_DELETE_SUCCEEDED:
      return {
        ...state,
        loading: false,
        error: null,
        list: state.list.filter((item) => action.payload !== item.hash),
      }

    default:
      return state
  }
}

export default mainReducer
