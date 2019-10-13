var q = 'messages';

var url = process.env.CLOUDAMQP_URL || "amqp://localhost";
var open = require('amqplib').connect(url);

// Publisher
open.then(function(conn) {
  var ok = conn.createChannel();
  return ok.then(function(ch) {
    ch.assertQueue(q);
    ch.sendToQueue(q, new Buffer('something to do'));
  })
})