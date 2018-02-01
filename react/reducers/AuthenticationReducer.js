import { combineReducers } from 'redux'
import { UserConstants } from '../constants/UserConstants';

function authenticationApi(
  state = {
    isFetching: false,
    didInvalidate: false,
    todos: []
  },
  action
) {
  switch (action.type) {
    case UserConstants.REQUEST_AUTHENTICATION:
      return Object.assign({}, state, {
        isFetching: true,
        loggedIn: false,
        didInvalidate: false
      })
    case UserConstants.RECEIVE_AUTHENTICATION:
      return Object.assign({}, state, {
        isFetching: false,
        didInvalidate: false,
        loggedIn: true,
        user: action.user,
        lastUpdated: action.receivedAt
      })
    case UserConstants.LOGOUT:
        return Object.assign({}, state, {
          isFetching: true,
          loggedIn: false,
          didInvalidate: false
    })
    default:
      return state
  }
}

let jwtToken = localStorage.getItem('jwt');
const initialState = jwtToken ? { loggedIn: true } : {};

function AuthenticationReducer(state = initialState, action) {
  switch (action.type) {
    case UserConstants.RECEIVE_AUTHENTICATION:
    	return authenticationApi(state, action)
    case UserConstants.REQUEST_AUTHENTICATION:
    	return authenticationApi(state, action)
    case UserConstants.LOGOUT:
    	return authenticationApi(state, action)
    default:
      return state
  }
}

export default AuthenticationReducer