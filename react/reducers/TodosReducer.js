import { combineReducers } from 'redux'
import { TodoConstants } from '../constants/TodoConstants';

function todosApi(
  state = {
    isFetching: false,
    didInvalidate: false,
    todos: []
  },
  action
) {
  switch (action.type) {
    case TodoConstants.REQUEST_LIST_TODO:
      return Object.assign({}, state, {
        isFetching: true,
        didInvalidate: false,
        todos: []
      })
    case TodoConstants.RECEIVE_LIST_TODO:
      return Object.assign({}, state, {
        isFetching: false,
        didInvalidate: false,
        todos: action.todos,
        lastUpdated: action.receivedAt
      })
    case TodoConstants.REQUEST_ADD_TODO:
        return Object.assign({}, state, {
          isFetching: true,
          didInvalidate: false
    })
    case TodoConstants.RECEIVE_ADD_TODO:
        return Object.assign({}, state, {
          isFetching: false,
          didInvalidate: false,
          todos: [
	        ...state.todos,
	        action.todo
	      ],	
          lastUpdated: action.receivedAt
    })
    default:
      return state
  }
}

function TodosReducer(state = {}, action) {
  switch (action.type) {
  	case TodoConstants.REQUEST_LIST_TODO:
  		return todosApi(state, action)
  	case TodoConstants.RECEIVE_LIST_TODO:
    	return todosApi(state, action)
    case TodoConstants.REQUEST_ADD_TODO:
    	return todosApi(state, action)
    case TodoConstants.RECEIVE_ADD_TODO:
    	return todosApi(state, action)
    default:
      return state
  }
}

export default TodosReducer