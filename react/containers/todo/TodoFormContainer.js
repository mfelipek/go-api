import React from 'react';
import { fetchAddTodo } from '../../actions/TodoActions'
import { connect } from 'react-redux'

class TodoFormContainer extends React.Component {
  constructor(props) {
    super(props);
    
    this.state = {value: ''};
    
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }
  
  handleSubmit(event) {
    event.preventDefault();
    this.props.fetchAddTodo(this.state.value)
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
          <input type="text" value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  return {}
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
	  fetchAddTodo: text => {
		  dispatch(fetchAddTodo(text))
	  }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(TodoFormContainer);  