const express = require('express');
const cookieParser = require('cookie-parser');
const logger = require('morgan');
const mysql = require('promise-mysql');
const cors = require('cors')

const app = express();

app.use(cors())
app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.get('/', (req, res) => {
  return res.send({
    message: 'hello'
  });
});

app.get('/messages', async (req, res) => {
  const query = 'SELECT * FROM messages';
  try {
    const connection = await mysql.createConnection({
      host: process.env.DB_HOST,
      user: process.env.DB_USER,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_NAME
    });
  
    const results = await connection.query(query);
    connection.end();
  
    return res.send(results);
  } catch(error) {
    return res.status(500).send({
      error,
      message: 'Work is not done yet...'
    })
  }
});

module.exports = app;
