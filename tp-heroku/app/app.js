const express = require('express');
const cookieParser = require('cookie-parser');
const logger = require('morgan');
const mysql = require('promise-mysql');
const cors = require('cors');

const app = express();

app.use(cors());
app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.get('/', (req, res) => {
  return res.send({
    message: `Hello ${process.env.NAME || 'World'}`
  });
});

app.get('/messages', async (req, res) => {
  const query = 'SELECT * FROM messages';
  try {
    const connection = await mysql.createConnection(process.env.TO_DEFINE || '');
  
    const results = await connection.query(query);
    connection.end();
  
    return res.send(results);
  } catch(error) {
    console.log(error);
    return res.status(500).send({
      error: error.message || error,
      message: 'Work is not done yet...'
    })
  }
});

module.exports = app;
