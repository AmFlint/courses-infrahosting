import React, { useState } from 'react';
import logo from './logo.svg';
import * as axios from 'axios';
import './App.css';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL || 'http://localhost:3000';

const fetchMessages = async () => {
  const response = await axios.get(`${BACKEND_URL}/messages`);
  return response.data;
}

function App() {
  const [messages, setMessages] = useState([]);
  fetchMessages()
    .then(messages => setMessages(messages))
  return (
    <div className="App">
      <header className="App-header">
        <h1>Messages:</h1>
        <ul>
          {
            messages.map(message => (
              <li>{message.content || ''}</li>
            ))
          }
        </ul>
      </header>
    </div>
  );
}

export default App;
