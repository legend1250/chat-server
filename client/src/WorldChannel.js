import React, { Component } from 'react'
import { Row, Form, Col } from 'antd'
import styles from './Div.module.scss'
import { inject } from 'mobx-react'
import { observer } from 'mobx-react-lite'

const WorldChannel = () => {
  return (
    <Row>
      <Col span={16}>
        <div className={styles['left-side']}>
          <Col span={4}>
            <OnlineUsers />
          </Col>
          <Col span={20}>
            <WorldMessages />
          </Col>
        </div>
        <div>
          <SendWorldMessageWrapper />
        </div>
      </Col>
      <Col span={8}>Span8</Col>
    </Row>
  )
}

export default WorldChannel

const OnlineUsers = () => {
  return <div>OnlineUsers</div>
}

const WorldMessages = inject('stores')(
  observer(({ stores }) => {
    // const { conn } = stores

    return <div>World Messages</div>
  })
)

class SendWorldMessage extends Component {
  render() {
    return (
      <Form>
        <div>SendWorldMessage</div>
      </Form>
    )
  }
}

const SendWorldMessageWrapper = Form.create()(SendWorldMessage)
