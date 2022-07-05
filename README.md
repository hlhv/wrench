# wrench

[![A+ Golang report card.](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/hlhv/wrench)

Swiss army knife for configuring HLHV.

## Usage

`wrench <Command> [-h|--help]`

### Commands

- `newkey`: Generate a new key
- `adduser`: Add a user for the specified cell
- `deluser`: Deletes the user for the specified cell
- `authuser`: Authorizes a user to access files for the specified cell

### Arguments

- `-h|--help`: Print help information`
- `newkey`
  - `-t|--text`: Contents of the key as text
  - `-c|--cost`: Cost of the key
- `adduser`
  - `-c|--cell`: Name of the cell to add a user for. Default: `queen`
- `deluser`
  - `-c|--cell`: Cell who's user will be deleted. Default: `queen`
- `authuser`
  - `-c|--cell`: Cell to grant access to. Default: `queen`
  - `-u|--user`: User to be given access
