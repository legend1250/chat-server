import React from 'react'
// import Child from './Child'
import { Home } from './views'
import { inject, observer } from 'mobx-react'
import './antd.less'

const App: React.FC = (props: any) => {
  const { loadingWS } = props.stores

  return !loadingWS ? <Home /> : <div>Loading...</div>
}

export default inject('stores')(observer(App))
