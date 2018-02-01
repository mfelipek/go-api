import { combineReducers } from 'redux'
import TodosReducer from './TodosReducer'
import AuthenticationReducer from './AuthenticationReducer'

const GoApiApp = combineReducers({
	TodosReducer,
	AuthenticationReducer
})

export default GoApiApp