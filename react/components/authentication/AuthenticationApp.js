import React from 'react'
import AuthenticationFormContainer from '../../containers/authentication/AuthenticationFormContainer'
import LoginHeader from './LoginHeader'

const AuthenticationApp = () => (
  <div>
  	<LoginHeader />
  	<AuthenticationFormContainer />
  </div>
)

export default AuthenticationApp