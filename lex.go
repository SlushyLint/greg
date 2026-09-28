package main

type Token struct {
    Type    string
    Value   string
}

// stupid peice of f*cking sh*t

func isdigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func isquote(c byte) bool {
    return c == '"'
}

func isletter(c byte) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
           c == '_'
}

func lex(src string) []Token {
    var tokens []Token

    for i := 0; i < len(src); i++ {
        if isquote(src[i]) {
            start := i + 1 
            i++
            for i < len(src) && !isquote(src[i]) {
                i++
            }
            if i >= len(src) {
                break
            }

            tokens = append(tokens, Token{
                Type: "s",
                Value: src[start:i],
            })
            continue
        }
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


        //  dont trust the below if much
        //if src[i] == ' '  || src[i] == '\n' || src[i] == '\t'{
        //    continue
        //} //  im fucking tired
        // i was correct, this is a peice of shit
            

        if src[i] == '\n' {
            tokens = append(tokens, Token{
                Type: "nl",
                Value: "\n",
            })
            continue
        }
        if src[i] == ' ' || src[i] == '\t' {
            continue
        }

        // still dont trust me though

        
        switch src[i] {
        case '+':
            tokens = append(tokens, Token{
                Type:   "op",
                Value:  "+",
            })
        case '-':
            tokens = append(tokens, Token{
                Type: "op",
                Value: "-",
            })
        case '*':
            tokens = append(tokens, Token{
                Type: "op",
                Value: "*",
            })
        case '/':
            tokens = append(tokens, Token{
                Type: "op",
                Value: "/",
            })
        case '=':
            tokens = append(tokens, Token{
                Type: "op",
                Value: "=",
            })
        case ':':
            tokens = append(tokens, Token{
                Type: "op",
                Value: ":",
            })
        case';':
        tokens  = append(tokens, Token{
            Type: "op",
            Value: ";",
        })
        
        }
    }

    return tokens
}
