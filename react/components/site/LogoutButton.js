import React from 'react'
import { NavLink } from 'react-router-dom'

class LogoutButton extends React.Component {
  render() {
    return (
    	<li className="nav-item">
    		<NavLink exact to="/logout" className="nav-link" activeClassName="active">
    			Logout
    		</NavLink>
    	</li>
	)
  }
}

export default LogoutButton