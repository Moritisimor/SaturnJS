#!/usr/bin/env saturn
while (true) {
    const lower = 0;
    const upper = 100;

    io.println(`I'm thinking of a number between ${lower} and ${upper}.`);
    const num = rand.randInt(lower, upper);
    while (true) {
        const guess = Number.parseInt(io.input('Enter your guess: '));
        if (Number.isNaN(guess)) {
            io.println('Only enter whole numbers.');
            continue;
        }

        if (guess === num) {
            io.println("That's right!");
            break;
        }

        io.println("That's wrong!");
        if (guess > num) {
            io.println("Go lower!");
        } else {
            io.println("Go higher!")
        }
    }

    if (io.input("Play again? [Y/n]: ").toLowerCase() === "y") {
        continue;
    }

    io.println("Goodbye!");
    break;
}
