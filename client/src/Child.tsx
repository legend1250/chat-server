import React, { useState } from 'react'
import { inject, observer } from 'mobx-react'

interface Props {
  stores?: any
}

const Child: React.FC<Props> = (props: Props) => {
  const { conn } = props.stores
  const [state, setstate] = useState<Array<string>>([])

  conn.onmessage = function(evt: any) {
    var { message } = JSON.parse(evt.data)
    console.log('messagees: ', message)
    // const msg = messages.pop()
    const update = [...state, message]
    setstate(update)
  }

  return (
    <div>
      {/* {state.map((str, index) => (
        <li key={`li-${index}`}>{str}</li>
      ))} */}
    </div>
  )
}

export default inject('stores')(observer(Child))
