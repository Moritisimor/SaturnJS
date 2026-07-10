#!/usr/bin/env saturn
const res = net.http.get("https://itcorp.com");

io.println(`Status: ${res.status}`);
io.println('Headers:');
for (const [key, value] of res.headers) {
    io.println(`\t=> ${key}: ${value}`);
}

io.println(res["body"]);
