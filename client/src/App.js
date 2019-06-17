import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import Child from './Child'

export var conn

function App() {

  const [initConn, setInitConn] = useState(false)

  useEffect(() => {
    if (window["WebSocket"]) {
      conn = new WebSocket("ws://localhost:8080/ws");
      setInitConn(true)
    }
  }, [])

  return (
    initConn ? <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <Child />
    </div>
    :
    <div>Loading...</div>
  );
}

export default App;
