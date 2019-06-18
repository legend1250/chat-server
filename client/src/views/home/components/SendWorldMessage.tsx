import React, { Component } from 'react'
import { Form, Input, Row } from 'antd'
import { inject, observer } from 'mobx-react'
import { FormComponentProps } from 'antd/lib/form'
import { StoresTypes } from '../../../stores'
import { FORM_ITEM_LAYOUT } from './common'

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
      stores: { conn },
      props: { form }
    } = this
    const { message } = form.getFieldsValue()
    conn.send(JSON.stringify({ message }))
    form.setFieldsValue({ message: '' })
  }

  render() {
    const { form } = this.props

    return (
      <Row>
        <Form onSubmit={this.onSubmit}>
          <FormItem label='message' {...FORM_ITEM_LAYOUT}>
            {form.getFieldDecorator('message', {})(<Input />)}
          </FormItem>
        </Form>
      </Row>
    )
  }
}

const SendWorldMessageWrapper = inject('stores')(observer(Form.create()(SendWorldMessage)))

export default SendWorldMessageWrapper
