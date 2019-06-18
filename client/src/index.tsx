import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import * as serviceWorker from './serviceWorker'
import { rootStore } from './stores'
import { Provider as MobxProvider } from 'mobx-react'

const stores = new rootStore()

ReactDOM.render(
  <MobxProvider stores={stores}>
    <App />
  </MobxProvider>,
  document.getElementById('root')
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
