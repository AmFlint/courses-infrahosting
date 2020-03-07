module.exports = {
    "up": async function(connection, cb) {
        await connection.query("CREATE TABLE IF NOT EXISTS `messages` (`id` int(11) NOT NULL AUTO_INCREMENT, `content` varchar(155) NOT NULL, PRIMARY KEY(id));");
        await connection.query("INSERT INTO `messages` (`id`, `content`) VALUES(1, 'hello'), (2, 'world'), (3, 'docker'), (4, 'compose');")


        cb();
    },
    "down": "DROP TABLE `messages`;"
}