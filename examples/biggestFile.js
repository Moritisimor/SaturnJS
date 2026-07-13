#!/usr/bin/env saturn
const args = os.getArgs();
if (args.length < 2) {
    io.println(color.redify("Invalid Arguments!"));
    io.println(color.greenify(`Usage: ${args[0]} <directory>`));
    os.exit(1);
}

let files;
try {
    files = fs.readDir(args[1])
        .filter(f => f.isFile())
        .sort((x, y) => y.size - x.size)
        .slice(0, 10);

} catch (_) {
    io.println(color.redify(`Could not access '${args[1]}'`));
    os.exit(1);
}

if (files.length === 0) {
    io.println(color.yellowify("No files in that directory!"))
    os.exit(0);
}

let idx = 1;
for (const file of files) {
    io.println(`${color.magentaify(idx)}, ${color.blueify(file.name)}: ${color.greenify(file.size)} ${color.magentaify("Bytes")}`);
    idx++;
}
