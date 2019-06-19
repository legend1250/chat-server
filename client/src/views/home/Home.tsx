import React, { Component } from 'react'
import { Row, Form, Col, Card } from 'antd'
import styles from './Home.module.scss'
import { inject, observer } from 'mobx-react'
import { SendWorldMessage, JoinRoom } from './components'
import { toJS } from 'mobx'

const Home = () => {
  return (
    <Row>
      <Col span={16}>
        <Row>
          <Col span={4} className={styles['left-side']}>
            <OnlineUsers />
          </Col>
          <Col span={20}>
            <WorldMessages />
          </Col>
        </Row>
      </Col>
      <Col span={8}>
        <JoinRoom />
      </Col>
    </Row>
  )
}

export default Home

const OnlineUsers = () => {
  return (
    <div>
      <Card title='OnlineUsers'>none</Card>
    </div>
  )
}

const WorldMessages = inject('stores')(
  observer(({ stores }: any) => {
    const { messages } = stores

    return (
      <div className={styles['left-side']}>
        <Card title='Messages'>
          <div className={styles['chat-body']}>
            {toJS(messages).map((m: any, index: number) => (
              <li key={`mess-${index}`}>{m.message}</li>
            ))}
          </div>
          <SendWorldMessage />
        </Card>
      </div>
    )
  })
)
