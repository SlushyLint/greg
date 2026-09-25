package main

type Token struct {
    Type    string
    Value   string
}


func isdigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func isletter(c byte) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
           c == '_'
}

func lex(src string) []Token {
    var tokens []Token

    for i := 0; i < len(src); i++ {
        if isletter(src[i]) {
            start := i
            for i < len(src) && isletter(src[i]) {
                i++
            }
            tokens = append(tokens, Token{
                Type:   "l",
                Value:  src[start:i],
            })
            i--
            continue
        }
        if isdigit(src[i]) {
            start := i
            for i < len(src) && isdigit(src[i]) {
                i++
            }
            
            tokens = append(tokens, Token{
                Type: "n",
                Value: src[start:i],
            })
            i--
            continue
        }
            
        switch src[i] {
        case '+':
            tokens = append(tokens, Token{
                Type:   "p",
                Value:  "+",
            })
        }
    }

    return tokens
}
