import React from 'react'
import { NavLink } from 'react-router-dom'

class LoginButton extends React.Component {
  render() {
    return (
    	<li className="nav-item">
    		<NavLink exact to="/login" className="nav-link" activeClassName="active">
    			Login
    		</NavLink>
    	</li>
	)
  }
}

export default LoginButton