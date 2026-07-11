# SaturnJS Standard Library 
This Markdown document serves as documentation for the standard library.

The standard library is always subject to changes, and although breaking changes probably won't happen, the API will be extended upon.

## io
Collection of functions for Input and Output via stdin and stdout.

### print
Prints to stdout without newline.

#### Signature
`io.print(a: variadic any): undefined`

### println
Prints to stdout with newline.

#### Signature
`io.println(a: variadic any): undefined`

### printf
Prints to stdout without newline according to format string and format arguments.

#### Signature
`io.printf(formatString: string, args: variadic any): undefined`

### input
Prints prompt to stdout without newline and reads a line from stdin, returning it as a string.

#### Signature
`io.input(prompt: string): string`

## fs
Collection of functions for interacting with the OS's filesystem.

### readDir
Reads a directory to an object array.

Throws if directory does not exist.

#### Signature
`fs.readDir(path: string): object array`

##### Object Fields
- name: string
- size: int 
- isDir: () => bool
- isFile: () => bool
- modificationUnix: int

### readFile
Reads file to a string.

Throws if file does not exist.

#### Signature
`fs.readFile(path: string): string`

### readFileLines
Reads file to an array of newline-seperated strings.

Throws if file does not exist.

#### Signature
`fs.readFileLines(path: string): string array`

### readFileBytes
Reads file to an array of integers that represent the bytes of the file.

Throws if file does not exist.

#### Signature
`fs.readFileBytes(path: string): integer array`

### readJsonObject
Reads file to a JSON object.

Throws if file does not exist or cannot be parsed to a JSON object.

#### Signature
`fs.readJsonObject(path: string): object`

### readJsonArray
Reads file to a JSON array.

Throws if file does not exist or cannot be parsed to a JSON array.

#### Signature
`fs.readJsonArray(path: string): object array`

### createFile
Creates a file.

Throws if path before file name does not exist or the user does not have the necessary permissions.

#### Signature
`fs.createFile(path: string): undefined`

### createDir
Creates a directory.

Throws if path before directory name does not exist or the user does not have the necessary permissions.

#### Signature
`fs.createDir(path: string): undefined`

### deleteFile
Deletes a file.

Throws if file does not exist or the user does not have the necessary permissions.

#### Signature
`fs.deleteFile(path: string): undefined`

### deleteDir
Recursively deletes a directory and all its children.

Throws if directory does not exist.

#### Signature
`fs.deleteDir(path: string): undefined`

### move
Moves a file to another path.

Throws if the file does not exist or the user does not have the necessary permissions.

#### Signature
`fs.move(oldPath: string, newPath: string): undefined`

### copy
Copies a file to another path.

Throws if `oldPath` does not exist, `newPath` is in a non-existant directory, or the user does not have the necessary permissions.

#### Signature
`fs.copy(oldPath: string, newPath: string): undefined`

### writeFile
Writes string to a file.

Throws if you are trying to write to a file that's in a non-existant directory or you don't have the necessary permissions.

#### Signature
`fs.writeFile(path: string, content: string)`

### writeJson
Writes JSON object or array to a file.

Throws if you are trying to write to a file that's in a non-existant directory or you don't have the necessary permissions.

#### Signature
`fs.writeJson(path: string, content: object)`

## os
Collection of functions for interacting with the Operating System.

### cwd
Returns the Current Working Directory as a string.

May throw.

#### Signature
`os.cwd(): string`

### chdir
Changes the Working Directory of the process.

Throws if the target directory does not exist

#### Signature
`os.chdir(path: string): undefined`

### getEnv
Gets environment variable.

Does not throw. If the key does not exist, an empty string is returned

#### Signature
`os.getEnv(key: string): string`

### setEnv
Sets environment variable.

May throw.

#### Signature
`os.setEnv(key: string, value: string): undefined`

### getArgs
Gets CLI Args, where the 0th element is the script name.

Does not throw.

#### Signature
`os.getArgs(): string array`

### getUser
Gets the name of the currently logged-in user.

Does not throw.

#### Signature
`os.getUser(): string`

### getHome
Gets the home directory of the currently logged-in user.

May throw.

#### Signature
`os.getHome(): string`

### exit
Exits the process with exit code.

#### Signature
`os.exit(code: int): undefined`

## exec
Collection of functions for spawning processes

### execute
Spawns a process, waiting for it to finish. Returns the process exit code.

#### signature
exec.execute(programName: string, programArgs: string array): integer

## color
Collection of functions for coloring strings via ANSI Escape Codes.

None of the functions throw.

### redify
Turns a string red.

#### Signature
`color.redify(x: string): string`

### blueify
Turns a string blue.

#### Signature
`color.blueify(x: string): string`

### greenify
Turns a string green.

#### Signature
`color.greenify(x: string): string`

### yellowify
Turns a string yellow.

#### Signature
`color.yelllowify(x: string): string`

### blackify
Turns a string black.

#### Signature
`color.blackify(x: string): string`

### magentaify
Turns a string magenta.

#### Signature
`color.magentaify(x: string): string`

### cyanify
Turns a string cyan.

#### Signature
`color.cyanify(x: string): string`

## net.http
Functions for networking via the HyperText Transfer Protocol.

### request
Sends a HTTP Request.

May throw for several reasons, like no server running on that address.

#### Signature
`net.http.request(url: string, method: string, body: string, headers: object): object`

Note that the `headers` object must only have keys where the value is a string.

##### Object Fields
- body: string
- bodyBytes: int array
- status: int
- headers: string array array
- getHeader: (key: string) => string
- ip: string

### get
Comfortable convenience function wrapping `request`. 

Primarily meant for simple, header- and bodyless GET requests, like downloads.

#### Signature
`net.http.get(url: string): object`

##### Object Fields
- body: string
- bodyBytes: int array
- status: int
- headers: string array array
- getHeader: (key: string) => string
- ip: string

### post
Convenience function wrapping `request`.

Primarily meant for simple POST requests.

#### Signature
`net.http.post(url: string, body: string, headers: object, contentType: string): object`

Note that the `headers` object must only have keys where the value is a string.

##### Object Fields
- body: string
- bodyBytes: int array
- status: int
- headers: string array array
- getHeader: (key: string) => string
- ip: string

## rand
A collection of functions for randomly generating numbers

### randInt
Inclusively generates a random integer between two numbers.

#### Signature
`rand.randInt(lower: int, upper: int): int`

### randFloat
Inclusively generates a random float between two numbers.

#### Signature
`rand.randFloat(lower: float, upper: float): float`

## sqlite
Functions for interacting with sqlite databases.

### connect
Connects to a sqlite database, creating a new one if it doesn't already exist.

May throw for several reasons, like the user not having the necessary permissions.

SQLite syntax for parameters applies here. A small example:
```js
const db = sqlite.connect(":memory");
db.execute("INSERT INTO myTable (val1, val2) VALUES (?, ?), (?, ?)", 1, 2, 3, 4)
```

The `?`'s are placeholders for arguments, which are later given to the function as parameters.

#### Signature
`sqlite.connect(path: string): object`

##### Object Methods
###### execute: (cmd: string, params: variadic any) => object
Executes a SQL statement, returning an object containing `insertRowId: int` and `rowsAffected: int`.

###### querySingle: (cmd: string, params: variadic any) => object array
Executes a SQL statement, returning an array containing the rows.

This SQL statement MUST return a single row, otherwise an exception will be thrown.

Example:
```js
const db = sqlite.connect(":memory:");
db.execute("CREATE TABLE people (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)");
db.execute("INSERT INTO people (name, age) VALUES ('John Doe', 21), ('Jane Doe', 22)");

// Here, id would be the 0th element of the array, name the 1st and age the 2nd.
const johnDoe = db.querySingle("SELECT id, name, age FROM people WHERE name = ?", "John Doe");
io.println(`Id: ${johnDoe[0]}, Name: ${johnDoe[1]}, Age: ${johnDoe[2]}`);
```

###### query: (cmd: string, params: variadic any) => object array array
Executes a SQL statement, returning the records as an array of arrays.

Example:
```js
const db = sqlite.connect(":memory:");
db.execute("CREATE TABLE people (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)");
db.execute(`
    INSERT INTO people (name, age) VALUES 
    ('John Doe', 21), 
    ('Jane Doe', 22), 
    ('Max Mustermann', 20),
    ('Erika Mustermann', 23)
`);

const people = db.query("SELECT id, name, age FROM people");
io.println("People:");
// JS Destructuring makes this pleasant to read and write
for (const [id, name, age] of people) {
    io.println(`\t- Id: ${id}, Name: ${name}, Age: ${age}`);
}
```

###### startTransaction: () => object
Starts a SQL transaction, returning the transaction as an object.

The transaction has the same methods as a regular database object, except that it also has:
- Commit: () => undefined
- Rollback: () => undefined

It is also missing the `close` method, as `close` is for the database only.

Example:
```js
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
```

###### close: () => undefined
Closes the SQLite connection.
