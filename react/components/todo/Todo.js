import React from 'react'

class Todo extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <li>
        <h1>{this.props.todo.id}</h1>
        <h2>It is {this.props.todo.name}.</h2>
      </li>
    );
  }
}

export default Todo