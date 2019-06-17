import React, { useState } from 'react'
import { conn } from './App'

const Child = () => {
  const [state, setstate] = useState([])

  conn.onmessage = function (evt) {
    var messages = evt.data.split('\n');
    const msg = messages.pop()
    const update = [...state, msg]
    setstate(update)
  };
  

  return(state.map((str, index) => <li key={`li-${index}`} >{str}</li>))
}

export default Child