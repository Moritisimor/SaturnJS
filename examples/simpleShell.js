const prompt = 'simpleShell >> ';
while (true) {
    const userInput = io.input(prompt);
    if (userInput === "exit") {
        break;
    }

    if (userInput === "") {
        continue;
    }

    const parts = userInput.split(" ");
    const programName = parts[0];
    const programArgs = parts.slice(1, parts.length);
    io.println(programName);
    io.println(programArgs);

    try {
        exec.execute(programName, programArgs);
    } catch (err) {
        io.println(`Error: ${err}`);
    }
}
