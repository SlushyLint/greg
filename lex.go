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

    indents := []int{0} 
    lnstart := true

    for i := 0; i < len(src); i++ { // we have to check indentation
        if lnstart {
            indent := 0 // NOT 'indents' there's no 's'
            for i < len(src) && src[i] == ' ' {
                indent++
                i++
            }

            if i >= len(src) {
                break
            }

            if src[i] == '\n' {
                tokens = append(tokens, Token{
                    Type: "nl", // newline
                    Value: "\n",
                })
                lnstart = true
                continue
            }

            curindent := indents[len(indents)-1]
            if indent > curindent { // check if the current indent is more than the total indentation?
                indents = append(indents, indent)
                tokens = append(tokens, Token{
                    Type:   "indent",
                })
            } else {
                for indent <  curindent {

                    indents = indents[:len(indents)-1]
                    curindent = indents[len(indents)-1]
                    tokens = append(tokens, Token{
                        Type:   "dedent",
                    })
                }
                if indent != curindent {
                    panic("[!] E: Inconsistent indentation")
                }
            }
            lnstart = false
            i-- // this pissed me off; the loop incriments i and we have to go back to the first non-space char
            continue
        }
                

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
            for i < len(src) && isdigit(src[i]) || src[i] == '.' {
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
            lnstart = true
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
        case '^':
            tokens = append(tokens, Token{
                Type: "dc", // declaration
                Value: "^",
            })
        case '>':
            tokens = append(tokens, Token{
                Type:   "op",     //opperator
                Value:  ">",
            })
        case  '<':
            tokens = append(tokens, Token{
                Type:   "op",   //operator
                Value:  "<",
            })
        case '{','}':
            tokens = append(tokens, Token{
                Type:   "brace",
                Value:  string(src[i]),
            })



        
        }
    }
    for len(indents) > 1{
        indents = indents[:len(indents)-1]
        tokens = append(tokens, Token{
            Type: "dedent",
        })
    }

    return tokens
}
