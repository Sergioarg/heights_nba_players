# Project Heights NBA Players

<p>
	<img align="right" src="https://play-lh.googleusercontent.com/I6Jz9nEZq-8jBPn2bmjywWLXLZ7GMn2WR64x9w1xQm8H5Isd1rhaQ4NDsp3jVUhLPFI" width="90" height="90">
	This project consists in the creation of a function which will find the pairs of the NBA players by adding the height of the players in inches with the information provided by the call of an api which contains a <code>.json</code> file with the information of the NBA players.
</p>

## Requirements

In order to run this project you need to have installed <strong>GO</strong> installed on our computer.

If you do not have it installed, you can go the following <a href="https://golang.org/doc/install">documentation</a> of Golang and install it according to your operating system.

## Compile and run

---

1. Clone the repositoriy.

```
$ git clone https://github.com/Sergioarg/heights_nba_players.git
```

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
---

- Run the program
- Introduces a numeric value

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
