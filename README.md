# File renaming tool
- simple CLI tool to rename all files in current directory containing specified string

## Build

1. `go build`

## Usage

### Add prefix
- `./file-renamer ${stringContainedInfilesToRename} prefix ${prefixToAdd}`

Example:

`./file-renamer test prefix 02/03/2020` → `test.js, testing.exe, testttttttt` => `02/03/2020-test.js, 02/03/2020-testing.exe, 02/03/2020-testttttttt`


### Add suffix 
- `./file-renamer ${stringContainedInfilesToRename} suffix ${suffixToAdd}`

Example:

`./file-renamer test suffix new`　→ `test.js, testing.exe, testttttttt` => `test-new.js, testing-new.exe, testttttttt-new`

### Rename
- `./file-renamer ${stringContainedInfilesToRename} ${newName}` 

Example:
`./file-renamer test coolCatPics` → `test.js, testing.exe, testttttttt` => `coolCatPics1.js, coolCatPics2.exe, coolCatPics3`