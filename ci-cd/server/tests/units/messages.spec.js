const messagesService = require('../../services/messages');

const connectionMock = {}

// -- Test Data -- //

const fixtureMessages = [
  { id: 1, content: 'first message' },
  { id: 2, content: 'test' },
]

test('get messages', async () => {
  connectionMock.query = (queryString) => fixtureMessages;

  const messages = await messagesService.getMessages(connectionMock);

  expect(messages.length).toBe(2);
  expect(JSON.stringify(messages)).toBe(JSON.stringify(fixtureMessages));
});

test('throws error', async () => {
  const errorMessage = 'database error';
  connectionMock.query = (queryString) => {
    throw new Error(errorMessage);
  }

  try {
    await messagesService.getMessages(connectionMock);
  } catch (err) {
    expect(err.message).toBe(errorMessage);
  }
});
