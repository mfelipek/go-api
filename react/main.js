import 'babel-polyfill'
import React from 'react'
import ReactDOM from 'react-dom';
import { createStore, applyMiddleware } from 'redux'
import thunkMiddleware from 'redux-thunk'
import { createLogger } from 'redux-logger'
import GoApiApp from './reducers'
import RouterConfig from './router';

const loggerMiddleware = createLogger()

let store = createStore(
  GoApiApp,
  applyMiddleware(
    thunkMiddleware,
    loggerMiddleware
  )
)

ReactDOM.render((
  <RouterConfig store={store} />
), document.getElementById('root'));

