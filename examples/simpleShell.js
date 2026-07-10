#!/usr/bin/env saturn
const prompt = 'simpleShell >> ';
while (true) {
    io.println(os.cwd());
    const userInput = io.input(prompt).trim();
    if (userInput === "exit") {
        break;
    }
 
    if (userInput === "") {
        continue;
    }

    const parts = userInput.split(" ");
    const programName = parts[0];
    const programArgs = parts.slice(1, parts.length);

    if (programName === "cd") {
        if (programArgs.length === 0) {
            io.println("Usage: cd <target directory>");
            continue;
        }       

        const targetDirectory = programArgs[0];
        try {
            os.chdir(targetDirectory);
        } catch (err) {
            io.println(`Error while changing directory: ${err}`);
        }

        continue;
    }

    try {
        exec.execute(programName, programArgs);
    } catch (err) {
        io.println(`Error: ${err}`);
    }
}
