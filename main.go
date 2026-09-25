package main
import(
    "fmt"
    "os"
)
// vars
var src []byte
var err error
func read() {
    src, err = os.ReadFile("test.st")
    if err != nil { 
        panic(err)
    }
    fmt.Println(string(src))
}


func main(){
    read()
    fmt.Print(src) 
    tokens := lex(string(src))
    parser := npars(tokens)
    
    for _, token := range tokens {
        fmt.Printf("\n%s: %q", token.Type, token.Value)
    }
}
