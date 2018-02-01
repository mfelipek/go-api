import React from 'react'
import { NavLink } from 'react-router-dom'
import { connect } from 'react-redux'
import LogoutButton from '../../components/site/LogoutButton'
import LoginButton from '../../components/site/LoginButton'

class MainNavBarContainer extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
	  
	  var loginButton;
		
	  if(this.props.loggedIn == true){
		  loginButton = <LogoutButton />;
	  }else{
		  loginButton = <LoginButton />;
	  }
	  
    return (
      <nav id="mainNavBar" className="navbar navbar-toggleable-md navbar-inverse bg-inverse fixed-top">
        <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault"
          aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <a className="navbar-brand" href="#">Navbar</a>
        <div className="collapse navbar-collapse">
          <ul className="navbar-nav mr-auto">
            <li className="nav-item" >
              <NavLink exact to="/" className="nav-link" activeClassName="active">
                Main
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink exact to="/todo" className="nav-link" activeClassName="active">
                My Todos
              </NavLink>
            </li>
            {loginButton}
          </ul>
        </div>
      </nav>
    );
  }
}

const mapStateToProps = state => {
  return {
	  loggedIn: state.AuthenticationReducer.loggedIn
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {}
}

export default connect(mapStateToProps, mapDispatchToProps)(MainNavBarContainer);  