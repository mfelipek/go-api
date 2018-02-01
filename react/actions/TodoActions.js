import fetch from 'cross-fetch'
import { TodoConstants } from '../constants/TodoConstants';
import { RequestConstants } from '../constants/RequestConstants';

export const requestAddTodo = text => {
  return {
    type: TodoConstants.REQUEST_ADD_TODO,
    text
  }
}

export const receiveAddTodo = (text, json) => {
	return {
	    type: TodoConstants.RECEIVE_ADD_TODO,
	    text,
	    todo: json.todo,
	    receivedAt: Date.now()
	}
}

export const requestTodos = () => {
  return {
    type: TodoConstants.REQUEST_LIST_TODO
  }
}

export const receiveTodos = json => {
	return {
	    type: TodoConstants.RECEIVE_LIST_TODO,
	    todos: json.todoList,
	    receivedAt: Date.now()
	}
}

export function fetchTodos() {

	return function (dispatch) {
		// inform we are starting to fetch todos
		dispatch(requestTodos())
		
		return fetch(
			'http://192.168.99.100:6060/api/todo', {
				method: "GET",
				headers: {
					"Authorization": "bearer " + localStorage.getItem('jwt')
				},
			})
			.then(
				response => {
					if (!response.ok) { 
		                return Promise.reject(response.statusText);
					}
					
					return response.json()
				},
				// Do not use catch erros
				error => console.log('An error occurred.', error)
		   )
		   .then(json =>
		   		// Here, we update the app state with the results of the API call.
		   		dispatch(receiveTodos(json))
		   )
	}
}

export function fetchAddTodo(text) {

	return function (dispatch) {
		// inform we are starting to fetch todos
		dispatch(requestAddTodo(text))

		return fetch('http://192.168.99.100:6060/api/todo', {
				method: "POST",
				body: JSON.stringify({ "Todo" : {"name": text} }),
				headers: {
					"Authorization": "bearer " + localStorage.getItem('jwt')
				},
			})
			.then(
				response => {
					if (!response.ok) { 
		                return Promise.reject(response.statusText);
					}
					
					return response.json()
				},
				// Do not use catch erros
				error => console.log('An error occurred.', error)
		   )
		   .then(response => {
		   		// Here, we update the app state with the results of the API call.
		   		if(response.status == RequestConstants.REQUEST_UNAUTHORIZED){
		   			dispatch(receiveAddTodo(text, response))
		   		}else{
		   			dispatch(receiveAddTodo(text, response))
		   		}
		   })
	}
}

