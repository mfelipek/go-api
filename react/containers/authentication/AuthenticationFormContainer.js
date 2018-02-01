import React from 'react';
import { fetchAuthentication } from '../../actions/UserActions'
import { connect } from 'react-redux'
import { Redirect } from 'react-router-dom'
import { withRouter } from 'react-router'

class AuthenticationFormContainer extends React.Component {
  constructor(props) {
    super(props);
    
    this.state = {
    	username: '',
    	password: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
	const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  handleSubmit(event) {
    event.preventDefault();
    this.props.fetchAuthentication(this.state.username, this.state.password);
  } 
  
  componentWillReceiveProps(nextProps){
	  const { from } = this.props.location.state || { from: { pathname: '/' } }
		
	  if(nextProps.loggedIn == true){
		  this.props.history.push(from) 
	  }  
  }

  render() {
	  
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Username:
          <input type="text" name="username" value={this.state.username} onChange={this.handleChange} />
        </label>
        <label>
	        Password:
	        <input type="password" name="password" value={this.state.password} onChange={this.handleChange} />
	      </label>
        <input type="submit" value="Submit" />
      </form>
    );
  }
}


const mapStateToProps = state => {
  return {
	  loggedIn: state.AuthenticationReducer.loggedIn
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
	  fetchAuthentication: (username, password) => {
		  dispatch(fetchAuthentication(username, password))
	  }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(withRouter(AuthenticationFormContainer));  