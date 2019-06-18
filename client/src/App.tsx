import React from 'react'
import Child from './Child'
import './antd.less'
import { inject, observer } from 'mobx-react'

const App: React.FC = (props: any) => {
  const { loadingWS } = props.stores

  return !loadingWS ? <Child /> : <div>Loading...</div>
}

export default inject('stores')(observer(App))
