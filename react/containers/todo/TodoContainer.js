import React from 'react';
import { connect } from 'react-redux'
import TodoList from '../../components/todo/TodoList'
import TodoFormContainer from './TodoFormContainer'
import { fetchTodos } from '../../actions/TodoActions'

class TodoContainer extends React.Component{
  constructor(props){
    super(props);
  }
  
  componentDidMount(){
	  this.props.fetchTodos();
  }

  render(){
    return (
      <div>
        <TodoFormContainer />
        <TodoList todos={this.props.todos}/>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    todos: state.TodosReducer.todos
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
	  fetchTodos: () => {
		  dispatch(fetchTodos())
	  }
  }
}


export default connect(mapStateToProps, mapDispatchToProps)(TodoContainer);