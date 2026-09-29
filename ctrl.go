package main

func (p *Pars) condition() bool {
    left := p.expr()
    op := p.current()

    if op.Type != "op" || (op.Value != ">" && op.Value != "<") {
        panic("expected > or < in condition")
    }
    p.adv()  // so we just ignore the error?
            // or does `panic` end the program???

    right := p.expr()
    if op.Value == ">"{
        return left > right
    }
    return  left < right    //return  ends the program, so by default, this 
                            // is essentially "else { return left < right }"
}
func (p *Pars) ifstmt() {

    p.adv() // consume if
    result := p.condition()
    p.block(result)

    // the above is new code

//    if p.current().Type != "brace" || p.current().Value != "{" {
//        panic("[!] E: Expected { after 'if'")
//    }
    // allow 'else' on the same line or the next line.
    // the above code does not do this, so we comment it out.
    // the above code was from a beta version, which did not support 'else'



//        for p.pos < len(p.tokens) && 
//        !(p.current().Type == "brace" && p.current().Value == "}"){
//            if p.current().Type == "nl" ||      // 
//            p.current().Type == "indent" ||     // this should let the parser ignore indentation artifacts
//            p.current().Type == "dedent" {      //
//                p.adv()
//                continue
//            }
//            p.stmt()
//        }
//        if p.pos >= len(p.tokens) {
//            panic("[!] E: Expected closing }")
//        }
//        p.adv() //consume }
//        return
//    }
//
//    //skip the body
//    depth := 1
//    for p.pos < len(p.tokens) && depth > 0 {
//        token := p.adv()
//        if token.Type == "brace" {
//            if token.Value == "{" {
//                depth++
//            } else if token.Value == "}" {
//                depth --
//            }
//        }
//    }
//    if depth != 0 {
//        panic("[!] expected } to close 'if'")
//    }
//}


    // the above code, again, is redacted. we do not need this anymore, as block() does most of this for us

    for p.current().Type == "nl" ||
    p.current().Type == "indent" ||
    p.current().Type == "dedent" {
        p.adv() // if its not obvious already, p.adv() is basically just "continue" but for
                // the parser to continue doing its shit, not go to say "continue"
    }

    if p.current().Type == "l" && p.current().Value == "else" {
        p.adv() // read else
        p.block(!result)
    }
}

func (p *Pars) block(run bool) {
    if p.current().Type != "brace" || p.current().Value != "{" {
        panic("[!] E: Expected {")
    }
    p.adv() //consume {
    if !run {
        //skip the lbock,
        //skip nested braces too
        depth :=1
        for p.pos < len(p.tokens) && depth > 0 {
            token := p.adv()
            if token.Type == "brace" {
                if token.Value == "{" {
                    depth++
                } else if token.Value == "}" {
                    depth--
                }
            }
        }
        if depth !=0{
            // error handling for unfinished brackets
            panic("[!] E: Expected }")
        }
        return
    }
    for p.pos < len(p.tokens) {
        token := p.current()
        if token.Type == "brace" && token.Value == "}" {
            p.adv() // read }
            return
        }

        //the lexer emits these indentations, we will ignore them here however

        if token.Type == "nl" || token.Type == "indent" ||  token.Type == "dedent" {
            p.adv()
            continue
        }
        p.stmt()
    }
    panic("[!] E: Expected }")
}

//    if token.Type == "brace"

//    p.adv() //advance past 'if'
//    result := p.condition() //check if we are dealing with a conditional
//
//    if p.current().Type != "nl" {   // "nl" is "\n"
//        panic("[!] E: Expected \\n after 'if'")
//    }
//    p.adv() // consume \n
//    if p.current().Type != "indent" {
//        panic("[!] E: Expected indent after 'if'")
//    }
//    p.adv() //consume indent, i forgor this earlier
//
//    if result { // checks if there is or isnt a conditional, utilizing a bool
//
//        for p.pos < len(p.tokens) && p.current().Type != "dedent" {
//            if p.current().Type == "nl" {
//                p.adv()
//                continue
//            }
//            p.stmt()    // process the statement
//        }
//        if p.current().Type != "dedent" {
//            panic("[!] E: Expected end of indent body")
//        }
//        p.adv()
//        return
//    }
//    depth := 1  // how deep the indentation goes
//    for p.pos<len(p.tokens)&&depth>0{ // valid go loop??
//        switch p.current().Type {
//        case "indent":
//            depth++
//        case "dedent":
//            depth--
//        }
//
//        p.adv()
//    }
//    if depth != 0 {
//        panic("[!] E: Expected end of indent body")
//    }
