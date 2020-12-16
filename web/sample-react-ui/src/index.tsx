import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import store from './redux/store'
import App from './App'
import 'semantic-ui-less/semantic.less'
import 'normalize.css'
render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.querySelector('#root') as HTMLElement
)