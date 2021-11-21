# Project Heights NBA Players

<p>
	<img align="right" src="https://play-lh.googleusercontent.com/I6Jz9nEZq-8jBPn2bmjywWLXLZ7GMn2WR64x9w1xQm8H5Isd1rhaQ4NDsp3jVUhLPFI" width="120" height="120">
	This project consists in the creation of a function which will find the pairs of the NBA players by adding the height of the players in inches with the information provided by the call of an api which contains a <code>.json</code> file with the information of the NBA players.
</p>

## Table of Contents

- [Requirements](#requirements)
- [Compile & run](#compile--run)
- [Usage](#usage)
- [Testing](#testing)
- [Files](#files)

## Requirements

In order to run this project you need to have installed **GO** installed on our computer.

If you do not have it installed, you can go the following [documentation](https://golang.org/doc/install) of Golang and install it according to your operating system.

## Compile & run

Clone the repositoriy.

```
$ git clone https://github.com/Sergioarg/heights_nba_players.git
```
Inside the folder runs the following commands in your terminal.

You can run the program in different ways.

1. The `go build` command will create an executable file according to the operating system, it will be executed differently.

On Linux

```
$ go build main.go
$ ./main
```

On Windows

```
> go build main.go
> main.exe
```

2. Run the program.

```
$ go run main.go
```

## Usage

- Run the program
- Introduces a numeric value after the message `Introduce Inch: number`.

### Exmaple:

```
$ go run main.go
Introduce Inch: 139
Brevin Knight   Nate Robinson
Mike Wilks      Nate Robinson
```

## Testing

You can do a quick test of the function with the command With predetermined values.

### Example:

```
$ go test
Input test: 139
Nate Robinson   Brevin Knight
Nate Robinson   Mike Wilks

Input test: 1
No matches found
```

### Files

| File                           | Description                       |
| ------------------------------ | --------------------------------- |
| [main.go](./main.go)           | Entry point to call the functions |
| [main_test.go](./main_test.go) | Tests made to test the function   |
| [go.mod](./go.mod)             | Describes the module’s properties |
