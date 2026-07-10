# SaturnJS Standard Library 
This Markdown document serves as documentation for the standard library.

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
- status: int
- headers: string array array
- getHeader: (key: string) => string
- ip: string
