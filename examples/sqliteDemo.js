#!/usr/bin/env saturn
const db = sqlite.connect(":memory:");
try {
    db.execute("CREATE TABLE people (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)");
    db.execute(`
        INSERT INTO people (name, age) VALUES 
        ('John Doe', 21), 
        ('Jane Doe', 22), 
        ('Max Mustermann', 20),
        ('Erika Mustermann', 23)
    `);

    // Here, id would be the 0th element of the array, name the 1st and age the 2nd.
    const johnDoe = db.querySingle("SELECT id, name, age FROM people WHERE name = ?", "John Doe");
    io.println(`Id: ${johnDoe[0]}, Name: ${johnDoe[1]}, Age: ${johnDoe[2]}`);

    const people = db.query("SELECT id, name, age FROM people");
    io.println("People:");
    // JS Destructuring makes this pleasant to read and write
    for (const [id, name, age] of people) {
        io.println(`\t- Id: ${id}, Name: ${name}, Age: ${age}`);
    }

    fs.writeJson("./people.json", people);
} catch (err) {
    io.println(`Error: ${err}`);
} finally {
    db.close();
}
