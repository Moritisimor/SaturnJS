#!/usr/bin/env saturn
const rnDate = time.date();
const rnTime = time.now();

io.println(`Day: ${rnDate.day}, Month: ${rnDate.month}, Year: ${rnDate.year}`);
io.println(`Second: ${rnTime.second}, Minute: ${rnTime.minute}, Hour: ${rnTime.hour}`)
