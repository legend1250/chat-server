import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'
import { Button, Card } from 'antd'
import { StoresTypes } from '../../../stores'
import SendRoomMessage from './SendRoomMessage'
import styles from './JoinRoom.module.scss'
import { toJS } from 'mobx'

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
      conn.send(JSON.stringify({ type: 3 }))
    }
  }

  render() {
    const {
      onClickJoin,
      stores: { roomInfo, roomMessages }
    } = this

    return (
      <div>
        <Card title='Private Room'>
          <div>
            <Button onClick={onClickJoin} type='primary'>
              {!roomInfo ? 'Join room' : 'Leave room'}
            </Button>
            {roomInfo && <span style={{ paddingLeft: 14 }}>roomId: {roomInfo.roomId}</span>}
          </div>
          <div className={styles['chat-body']}>
            {toJS(roomMessages).map((m: any, index: number) => (
              <li key={`mess-${index}`}>{m.message}</li>
            ))}
          </div>
          <SendRoomMessage />
        </Card>
      </div>
    )
  }
}

export default inject('stores')(observer(JoinRoom))
