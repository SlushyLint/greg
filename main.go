package main
import(
    "fmt"
    "os"
)
// vars
var src []byte
var err error
func read(file  string) {
    src, err = os.ReadFile(file)
    if err != nil { 
        panic(err)
    }
}

var file = ""
var verbose = false

func main(){
    defer func(){
        if err := recover(); err !=  nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    }()
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
//    parser.stmt()
//    result := parser.expr()
//    parser := npars(tokens)
//    name := parser.current().Value
//    parser.adv()
//    parser.adv() // took me ages to figure out how to skip "=" >:(
//    value := parser.expr()
//    vars[name] = value
//
     if verbose {
        fmt.Println(string(src))

        for _, token := range tokens {
            fmt.Printf("\n%s: %q", token.Type, token.Value)
        }
        for name, value := range vars {
            fmt.Printf("\n%s = %v", name, value)
        }
        fmt.Printf("\n\nOUT:\n\n")
    }
 
    for parser.pos < len(parser.tokens) {
        if parser.current().Type == "nl" {
            parser.adv()
            continue
        }
    
        oldPos := parser.pos
        parser.stmt()
    
        if parser.pos == oldPos {
            panic("parser got stuck")
        }
    }
}
