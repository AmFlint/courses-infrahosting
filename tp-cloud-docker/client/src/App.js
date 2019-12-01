import React from 'react';
import * as axios from 'axios';
import './App.css';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL || 'http://localhost:3000';

const fetchMessages = async () => {
  const response = await axios.get(`${BACKEND_URL}/messages`);
  return response.data;
}

class App extends React.Component {
  state = {
    messages: [],
    message: '',
  };

  componentDidMount() {
    this.fetchMessages();
  }

  fetchMessages = async () => {
    try {
      const messages = await fetchMessages();
      this.setState({ messages });
    } catch (err) {
      console.log(err);
    }
  }

  handleSubmit = async (event) => {
    event.preventDefault();
    await axios.post(`${BACKEND_URL}/messages`, {message: this.state.message});
    this.setState({ message: '' });
    setTimeout(() => this.fetchMessages(), 2000);
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1>Create a message:</h1>
          <form onSubmit={this.handleSubmit}>
            <input
              type="text"
              value={this.state.message}
              onChange={event => this.setState({ message: event.target.value })}
            />
            <button type="submit">Send message</button>
          </form>
          <h1>Messages:</h1>
          <ul>
            {
              this.state.messages.map(message => (
                <li key={message.id}>{message.content || ''}</li>
              ))
            }
          </ul>
        </header>
      </div>
    );
  }
}

export default App;
