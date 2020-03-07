const mysql = require('promise-mysql');

let connection;

const getConnection = async (configuration) => {
  if (!connection) {
    connection = await mysql.createConnection({
      host: process.env.DB_HOST,
      user: process.env.DB_USER,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_NAME
    });
  }
  return connection;
}

const getMessages = async (conn) => {
  const query = 'SELECT * FROM messages';
  const results = await conn.query(query);

  return results;
}

module.exports = {
  getMessages,
  getConnection,
  connection,
}
