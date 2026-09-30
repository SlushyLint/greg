package main
import(
    "fmt"
    "os"
)

// src and err hold the source file contents and the most recent read error.
var src []byte
var err error

// read loads the selected source file into src.
func read(file  string) {
    src, err = os.ReadFile(file)
    if err != nil { 
        panic(err)
    }
}

var file = ""
var verbose = false

// main loads, tokenizes, and executes the source file passed on the command line.
func main(){
    defer func(){
        // Convert panics from file loading, lexing, or parsing into a CLI error.
        if err := recover(); err !=  nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    }()
    // -v enables diagnostics; any other argument selects the input file.
    for _, arg := range os.Args[1:] {
        if arg == "-v" {
            verbose = true
        } else {
            file = arg
        }
    }
    read(file)
    tokens := lex(string(src))
    parser := par(tokens)

     if verbose {
        // Show the source and token stream before executing any statements.
        fmt.Println(string(src))

        for _, token := range tokens {
            fmt.Printf("\n%s: %q", token.Type, token.Value)
        }
        for name, value := range vars {
            fmt.Printf("\n%s = %v", name, value)
        }
        fmt.Printf("\n\nOUT:\n\n")
    }
 
    // Execute statements until the parser reaches the end of the token stream.
    for parser.pos < len(parser.tokens) {
        if parser.current().Type == "nl" {
            parser.adv()
            continue
        }
    
        oldPos := parser.pos
        parser.stmt()
    
        // A statement must consume input; otherwise malformed input could loop forever.
        if parser.pos == oldPos {
            panic("parser got stuck")
        }
    }
}
