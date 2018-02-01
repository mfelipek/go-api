import React from 'react';
import { logout } from '../../actions/UserActions'
import { connect } from 'react-redux'
import { withRouter } from 'react-router'

class LogoutContainer extends React.Component {
  constructor(props) {
    super(props);
    
    this.props.logout();
    this.props.history.push({
    	  pathname: '/login'
    });
  }
  
  render() {	  
    return (<div></div>);
  }
}


const mapStateToProps = state => {
  return {}
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
	  logout: () => {
		  dispatch(logout())
	  }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(withRouter(LogoutContainer));  