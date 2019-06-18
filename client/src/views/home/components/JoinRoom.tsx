import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'
import { Button } from 'antd'
import { StoresTypes } from '../../../stores'

class JoinRoom extends Component<{ stores?: any }> {
  get stores(): StoresTypes {
    return this.props.stores
  }

  onClickJoin = () => {
    const { conn, roomInfo } = this.stores
    // join room
    if (!roomInfo) {
      conn.send(JSON.stringify({ type: 1 }))
    } else {
      // leave room
      conn.send(JSON.stringify({ type: 2 }))
    }
  }

  render() {
    const {
      onClickJoin,
      stores: { roomInfo }
    } = this

    return (
      <div>
        <Button onClick={onClickJoin}>{!roomInfo ? 'Join room' : 'Leave room'}</Button>
        {roomInfo && <span style={{ paddingLeft: 14 }}>roomId: {roomInfo.roomId}</span>}
      </div>
    )
  }
}

export default inject('stores')(observer(JoinRoom))
