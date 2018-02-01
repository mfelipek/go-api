import React from 'react'
import {
  BrowserRouter as Router,
  Route,
  Redirect
} from 'react-router-dom'
import { Provider } from 'react-redux'

import MainNavBar from './containers/site/MainNavBarContainer'
import LogoutContainer from './containers/authentication/LogoutContainer';
import AuthenticationApp from './components/authentication/AuthenticationApp';
import TodoApp from './components/todo/TodoApp';

const routes = [
	{ 
		path: '/login',
		component: AuthenticationApp,
		exact: true
	},  
	{
		path: '/erro',
		exact: true
	}
]

const privateRoutes = [
	{
		path: '/todo',
		component: TodoApp,
		exact: false
	},
	{
		path: '/logout',
		component: LogoutContainer,
		exact: false
	}
]

// wrap <Route> and use this everywhere instead, then when
// sub routes are added to any route it'll work
const RouteWithSubRoutes = (route) => (
  <Route path={route.path} exact={route.exact == true} render={props => (
    // pass the sub-routes down to keep nesting
    <route.component {...props} routes={route.routes}/>
  )}/>
)

const PrivateRouteWithSubRoutes = (route) => (
  <Route path={route.path} exact={route.exact == true} render={props => (
		  
	localStorage.getItem('jwt') != null ? (
	//pass the sub-routes down to keep nesting
	<route.component {...props} routes={route.routes}/>
   ) : (
    <Redirect to={{
       pathname: '/login',
       state: { from: props.location.pathname }
    }}/>
   )
    
  )}/>
)

const RouterConfig = ({ store }) => (
	<Provider store={store}>
		<Router>
			<div>
				<MainNavBar />
				<div className="container">
					{routes.map((route, i) => (
					  <RouteWithSubRoutes key={i} {...route}/>
					))}
		
					{privateRoutes.map((route, i) => (
					  <PrivateRouteWithSubRoutes key={i} {...route}/>
					))}
				</div>
			</div>
	  </Router>
	</Provider>
)

export default RouterConfig