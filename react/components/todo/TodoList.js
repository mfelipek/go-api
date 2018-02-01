import React from 'react'
import Todo from './Todo'

class TodoList extends React.Component {

  constructor(props) {
    super(props);
  }
   
  render() {
	if(this.props.todos){
		return (
		<ul>
        	{this.props.todos.map(function(todo) {
        		return <Todo todo={todo} key={todo.id} />
        	})}
        </ul>
		)
	}else{
		return (
	        <span>"Empty"</span>
	    )
	}
  }
}

export default TodoList