const db = sqlite.connect(':memory:');
const tx = db.startTransaction();
try {
    tx.execute("CREATE TABLE people (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)");
    tx.execute(
        "INSERT INTO people (name, age) VALUES (?, ?), (?, ?)",
        "John Doe", 20,
        "Jane Doe", 21
    );

    tx.commit();
    for (const [name, age] of db.query("SELECT name, age FROM people")) {
        io.println(`Name: ${name}, Age: ${age}`);
    }
} catch (err) {
    io.println(`Error while running transaction: ${err}`);
    tx.rollback();
} finally {
    db.close();
}
