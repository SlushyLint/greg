# Greg

Greg is a tiny Go-based interpreter for a very small, experimental language.
This project is still evolving, so the syntax below reflects the currently
working implementation rather than a formal language specification.

## Requirements

- Go 1.27.1 or later

## Running

You can execute a source file directly:

```sh
go run . path/to/program.hs
```

Use `-v` to print the source, tokens, and current variable values before the
program runs:

```sh
go run . -v path/to/program.hs
```

You can also build a standalone binary:

```sh
go build -o greg .
./greg path/to/program.hs
```

## Current language features

The parser currently supports:

- typed declarations: `i^`, `f^`, and `s^`
- numeric expressions with `+`, `-`, `*`, and `/`
- comparisons with `>` and `<`
- `++` and `--` increment/decrement operators
- `if` / `else` blocks with braces
- `for (init; condition; update) { ... }` loops
- `printf("format", args...)` output

## Examples

Declare and print variables:

```text
i^count = 2
f^ratio = 3.5
s^message = "hello"

printf("%v %v %v\n", count, ratio, message)
```

This prints:

```text
2 3.5 hello
```

Conditional logic:

```text
i^count = 2

if count > 1 {
    printf("large\n")
} else {
    printf("small\n")
}
```

This prints:

```text
large
```

For loops:

```text
for (i^i = 0; i < 10; i++) {
    printf("%v ", i)
}
```

This prints:

```text
0 1 2 3 4 5 6 7 8 9 
```

## Important caveat

This interpreter is intentionally minimal and not fully general-purpose. In the
current implementation, simple reassignment such as `count = count + 1` is not
handled by the parser; use declaration syntax plus `++` / `--` for updates.

The example programs in the repository are under `test/` and are the best
reference for the current syntax and behavior.


*(notice, this README was written with AI, as I'm a lazy motherfucker with a life)*