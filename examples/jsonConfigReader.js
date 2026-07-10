#!/usr/bin/env saturn
const config = fs.readJsonObject('config.json');
io.println(`Port: ${config.port}, Host: ${config.host}`);
io.println(`DB port: ${config.db_port}, DB host: ${config.db_host}`);
io.println(`DB username: ${config.db_user}, DB password: ${config.db_password}`);
io.println(`Database: ${config.database}`);
io.println('Blacklisted IPs:');
for (const ip of config.blacklisted_ips) {
    io.println(`\t=> ${ip}`);
}
