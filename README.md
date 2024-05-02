## TODIZER

**todizer** is a very basic utility that lets you build an **interactive todo list** from a list of sentences
stored in a file (the default file stdin). The state of the list can optionally be stored in an output file or
printed in stdout by default.

### Installation

There is no prepackaged executable, you need to have [Go](https://go.dev/) on your machine

1. Clone this repository on your local machine 
2. Move into the root of the cloned directory
3. Run the following command

```bash
go mod tidy 
go run cmd/main.go
```
4. Enter your list of todos
5. Hit `Ctrl+d` whenever you finish editing your list

That is all !

### Keybindings

+ To move up, you can use either of these keys `Arrow up key`, `Ctrl+p`, `k`
+ To move down, you can use either of these keys `Arrow down key`, `Ctrl+n`, `j`
+ To mark a todo as done, you can you either of these keys `Enter`, `Ctrl+y`
+ To quit, hit `q`

### Démo

