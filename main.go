package main
import(
    "fmt"
    "os"
)
// vars
var src []byte
var err error
func read() {
    src, err = os.ReadFile(os.Args[1])
    if err != nil { 
        panic(err)
    }
    fmt.Println(string(src))
}


func main(){
    read()
//    fmt.Print(src) 
    tokens := lex(string(src))
    parser := par(tokens)
//    parser.stmt()
//    fmt.Println(vars["x"])
//    result := parser.expr()
//    parser := npars(tokens)
//    name := parser.current().Value
//    parser.adv()
//    parser.adv() // took me ages to figure out how to skip "=" >:(
//    value := parser.expr()
//    vars[name] = value
//
//    fmt.Println(vars[name])
    
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
   for _, token := range tokens {
        fmt.Printf("\n%s: %q", token.Type, token.Value)
    }
    for name, value := range vars {
        fmt.Printf("\n%s = %v", name, value)
    }
}
