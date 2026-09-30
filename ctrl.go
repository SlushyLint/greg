package main

// condition parses two numeric expressions separated by < or >.
func (p *Pars) condition() bool {
    left := p.expr()
    op := p.current()

    if op.Type != "op" || (op.Value != ">" && op.Value != "<") {
        panic("expected > or < in condition")
    }
    p.adv()

    right := p.expr()
    if op.Value == ">" {
        return left > right
    }
    return left < right
}

// ifstmt starts parsing an if/elif/else chain at the current token.
func (p *Pars) ifstmt() {
    p.ifbranch(true)
}

// block parses a braced block, executing its statements only when run is true.
func (p *Pars) block(run bool) {
    if p.current().Type != "brace" || p.current().Value != "{" {
        panic("[!] E: Expected {")
    }
    p.adv() //consume {
    if !run {
        // Keep nested blocks balanced while skipping an inactive branch.
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

        // Indentation tokens are informational; braces delimit this block.
        if token.Type == "nl" || token.Type == "indent" ||  token.Type == "dedent" {
            p.adv()
            continue
        }
        p.stmt()
    }
    panic("[!] E: Expected }")
}

// ifbranch parses one if/elif branch and any following alternatives.
func (p *Pars) ifbranch(allowed bool) {
    p.adv()
    matches := p.condition()
    // allowed is false after an earlier branch has already matched.
    run := allowed && matches
    p.block(run)

    for p.current().Type == "nl" ||
    p.current().Type == "indent" ||
    p.current().Type == "dedent" {
        p.adv()
    }
    if p.current().Type == "l" && p.current().Value == "elif" {
        p.ifbranch(allowed && !matches)
        return
    }
    if p.current().Type == "l" && p.current().Value == "else" {
        p.adv()

        // An else-if continues the chain with the same branch eligibility.
        if p.current().Type == "l" &&  p.current().Value == "if" {
            p.ifbranch(allowed && !matches)
            
        } else{
            p.block(allowed && !matches)
            
        }
    }
}
