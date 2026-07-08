const args = os.getArgs();
if (args.length < 2) {
    io.println(color.redify('usage: listFiles <directory path>'));
    os.exit(1);
}

for (const fsNode of fs.readDir(args[1]).filter(f => f.isFile())) {
    io.println(`${color.greenify('=>')} ${color.blueify(fsNode.name)}`);
}
