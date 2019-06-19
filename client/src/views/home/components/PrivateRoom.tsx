import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'
import { Button, Card, Form, Input, Row } from 'antd'
import SendRoomMessage from './SendRoomMessage'
import styles from './PrivateRoom.module.scss'
import { withModal } from '@components'
import { toJS } from 'mobx'
import { FORM_ITEM_LAYOUT } from './common'
import { FormComponentProps } from 'antd/lib/form'

const FormItem = Form.Item

const PrivateRoom = () => (
  <div>
    <Card title='Private Room'>
      <ButtonArea />
      <MessageArea />
      <SendRoomMessage />
    </Card>
  </div>
)

export default PrivateRoom

const ButtonArea = inject('stores')(
  observer(({ stores }) => {
    const { conn, roomInfo } = stores

    const onClickJoin = () => {
      // join room
      if (!roomInfo) {
        conn.send(JSON.stringify({ type: 1 }))
      } else {
        // leave room
        conn.send(JSON.stringify({ type: 3 }))
      }
    }

    return (
      <div className={styles['button-join-area']}>
        <Button onClick={onClickJoin} type='primary'>
          {!roomInfo ? 'Quick join' : 'Leave room'}
        </Button>
        <RoomJoinBtn />
        <div>{roomInfo && <span style={{ paddingLeft: 14 }}>roomId: {roomInfo.roomId}</span>}</div>
      </div>
    )
  })
)

const MessageArea = inject('stores')(
  observer(({ stores }) => {
    const { roomMessages } = stores

    return (
      <div className={styles['chat-body']}>
        {toJS(roomMessages).map((m: any, index: number) => (
          <li key={`mess-${index}`}>{m.message}</li>
        ))}
      </div>
    )
  })
)

const RoomJoinBtn = withModal(
  inject('stores')(
    observer(({ stores, modal }: any) => {
      const { conn, roomInfo } = stores
      console.log('modal: ', modal)

      const onClickJoin = () => {
        modal.show({
          title: 'Input Room Code',
          body: <FormInputRoomCodeWrapper />
        })
      }
      if (!roomInfo) {
        return (
          <Button onClick={onClickJoin} type='primary' style={{ marginLeft: 24 }}>
            Join room
          </Button>
        )
      } else {
        return <></>
      }
    })
  )
)

const FormInputRoomCodeWrapper = Form.create()(
  class FormInputRoomCode extends Component<FormComponentProps> {
    onSubmitForm = (e: React.FormEvent) => {
      e.preventDefault()
      const { form } = this.props
      form.validateFields(['code'], (err, values) => {
        if (!err) {
          console.log('values: ', values)
        }
      })
    }

    render() {
      const { form } = this.props
      return (
        <Form onSubmit={this.onSubmitForm}>
          <FormItem label='code' {...FORM_ITEM_LAYOUT}>
            {form.getFieldDecorator('code', {
              rules: [
                {
                  type: 'string',
                  required: true,
                  whitespace: true,
                  message: 'Code could not be empty'
                }
              ]
            })(<Input />)}
          </FormItem>
          <FormItem>
            <Row type='flex' justify='center'>
              <Button htmlType='submit' type='primary'>
                JOIN
              </Button>
            </Row>
          </FormItem>
        </Form>
      )
    }
  }
)
