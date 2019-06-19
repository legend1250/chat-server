import React, { Component } from 'react'
import { Form, Input, Row, message as AntMessage } from 'antd'
import { inject, observer } from 'mobx-react'
import { FormComponentProps } from 'antd/lib/form'
import { StoresTypes } from '../../../stores'
import { FORM_ITEM_LAYOUT_SMALL } from './common'

const FormItem = Form.Item

interface Props {
  stores?: any
}

class SendWorldMessage extends Component<Props & FormComponentProps> {
  get stores(): StoresTypes {
    return this.props.stores
  }

  onSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    const {
      stores: { conn, roomInfo },
      props: { form }
    } = this
    if (!roomInfo) {
      AntMessage.error(`You haven't joined a room`)
    }
    const { message } = form.getFieldsValue()
    conn.send(JSON.stringify({ type: 5, message }))
    form.setFieldsValue({ message: '' })
  }

  render() {
    const { form } = this.props

    return (
      <Row>
        <Form onSubmit={this.onSubmit}>
          <FormItem label='message' {...FORM_ITEM_LAYOUT_SMALL}>
            {form.getFieldDecorator('message', {})(<Input />)}
          </FormItem>
        </Form>
      </Row>
    )
  }
}

const SendWorldMessageWrapper = inject('stores')(observer(Form.create()(SendWorldMessage)))

export default SendWorldMessageWrapper
