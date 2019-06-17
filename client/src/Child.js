import React, { useState } from 'react'
import { conn } from './App'

const Child = () => {
  const [state, setstate] = useState([])

  conn.onmessage = function (evt) {
    var { message } = JSON.parse(evt.data);
    console.log('messagees: ',message)
    // const msg = messages.pop()
    const update = [...state, message]
    setstate(update)
  };
  

  return(state.map((str, index) => <li key={`li-${index}`} >{str}</li>))
}

export default Child