const amqpLib = require('amqplib');
const mysql = require('promise-mysql');

const url = process.env.AMQP_URL || "amqp://localhost";
const q = 'messages';

async function main() {
  const connection = await amqpLib.connect(url);
  const channel = await connection.createChannel();
  // Create message from Rabbit MQ
  channel.assertQueue(q);
  channel.consume(q, async (amqpMessage) => {
    if (amqpMessage === null) {
      console.log('message is empty');
      return;
    }
    const message = amqpMessage.content.toString();
    const query = 'INSERT INTO messages (content) VALUES(?)';
    const params = [message];
    try {
      const formattedQuery = mysql.format(query, params);
      const connection = await mysql.createConnection({
        host: process.env.DB_HOST,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME
      });
    
      await connection.query(formattedQuery);
      connection.end();
      console.log(`Message: ${message} created properly`);
      channel.ack(amqpMessage);
    } catch(error) {
      console.log(error);
    }
  })
}

main();
