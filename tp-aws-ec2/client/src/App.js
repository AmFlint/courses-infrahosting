import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import * as axios from 'axios';
import './App.css';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL || 'http://localhost:3000';

const fetchMessages = async () => {
  const response = await axios.get(`${BACKEND_URL}/messages`);
  return response.data;
}

class App extends React.Component {
  state = {
    messages: []
  };


  componentDidMount() {
    fetchMessages()
      .then(messages => this.setState({ messages }))
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1>Messages:</h1>
          <ul>
            {
              this.state.messages.map(message => (
                <li>{message.content || ''}</li>
              ))
            }
          </ul>
        </header>
      </div>
    );
  }
}

export default App;
