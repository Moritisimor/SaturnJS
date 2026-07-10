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
name: string
size: int 
isDir: () => bool
isFile: () => bool
modificationUnix: int

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
