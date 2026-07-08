const args = os.getArgs();
if (args.length < 2) {
    io.println('Usage: logAnalyzer.js <log file>');
    os.exit(1);
}

const logFilePath = args[1];
let errorLines = [];
let warningLines = [];

const lines = fs.readFileLines(logFilePath);
let currentLine = 1;
for (const line of lines) {
    if (line.toLowerCase().includes('error')) {
        errorLines.push(currentLine);
    }

    if (line.toLowerCase().includes('warning')) {
        warningLines.push(currentLine);
    }

    currentLine++;
}

if (errorLines.length === 0) {
    io.println(color.greenify(`No errors!`));
} else {
    io.println(color.redify(`${errorLines.length} errors!`));
    for (const line of errorLines) {
        io.println(color.blueify(`- ${line}`));
    }
}

if (warningLines.length === 0) {
    io.println(color.greenify('No warnings!'));
} else {
    io.println(color.yellowify(`${warningLines.length} warnings!`));
    for (const line of warningLines) {
        io.println(color.blueify(`- ${line}`));
    }
}
