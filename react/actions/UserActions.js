import fetch from 'cross-fetch'
import { UserConstants } from '../constants/UserConstants';
import { RequestConstants } from '../constants/RequestConstants';

export const requestAuthentication = username => {
  return {
    type: UserConstants.REQUEST_AUTHENTICATION,
    username
  }
}

export const receiveAuthentication = (username, json) => {
	return {
	    type: UserConstants.RECEIVE_AUTHENTICATION,
	    username,
	    receivedAt: Date.now()
	}
}

export function fetchAuthentication(username, password) {

	return function (dispatch) {
		// inform we are starting to fetch todos
		dispatch(requestAuthentication(username))

		return fetch('http://192.168.99.100:6060/api/sessions', {
				method: "POST",
				body: JSON.stringify({ "Username": username, "Password": password}),
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
		   		localStorage.setItem('jwt', response.token);
		   		
		   		dispatch(receiveAuthentication(username, response))
		   })
	}
}

export const logout = () => {
	localStorage.removeItem('jwt');
    return { type: UserConstants.LOGOUT };
}