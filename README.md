# Greg

Greg is a small interpreter written in Go and is under active, rapid
development. Its syntax and behavior may change without notice; the features
below describe the current implementation, not a stable language
specification.

## Requirements

- Go 1.27.1 or later

## Run

Pass a Greg source file to the interpreter:

```sh
go run . path/to/program.greg
```

Use `-v` to print the source, tokens, and variable values before execution:

```sh
go run . -v path/to/program.greg
```

You can also build an executable with `go build -o greg .` and run it as
`./greg path/to/program.greg`.

## Language basics

Declare variables with a type marker (`i` for integer, `f` for float, and `s`
for string), followed by `^`:

```text
i^count = 2
f^ratio = 3.5
s^message = "hello"
```

Numeric variables can be reassigned and used in expressions with `+`, `-`,
`*`, and `/`. Print a literal or variable with `print`:

```text
count = count + 1
print count
print "done"
```

Conditional blocks use braces and support `<` and `>` comparisons, as well as
`elif` and `else`:

```text
if count > 2 {
    print "large"
} else {
    print "small"
}
```