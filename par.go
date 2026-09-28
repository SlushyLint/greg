package main

import(
    "strconv"
)

var  vars = make(map[string]float64)

type Pars struct {
    tokens  []Token
    pos     int
}

func par(tokens []Token) *Pars {
    return &Pars{
        tokens: tokens,
        pos:    0,
    }
}


func (p *Pars) adv() Token{
    if p.pos >= len(p.tokens) {
        return Token{}
    }
    token := p.tokens[p.pos]
    p.pos++
    return token
}
func (p *Pars) num() float64 {
    token := p.current()
    if token.Type == "n" {
        n, _ := strconv.ParseFloat(token.Value, 64)
        p.adv()
        return n
    }
    if token.Type == "l" {
        value:= vars[token.Value]
        p.adv()
        return value
    }
    return 0
}
func (p *Pars) current() Token {
    return p.tokens[p.pos]
}

func (p *Pars) term() float64 {
    l := p.num()
    
    for  p.pos < len(p.tokens) {
        op := p.current().Value
        if op != "*" && op != "/" {
            break
        }

        p.adv()
        r := p.num()

        switch op {
        case "*":
            l *= r
        case "/":
            l /= r
        }
    }
    return l
}

func (p *Pars) expr() float64 {
    l := p.term()

    for p.pos < len(p.tokens) {
        op := p.current().Value

        if op != "+" &&  op != "-" {
            break
        }

        p.adv()

        r := p.term()

        switch op {
        case "+":
            l += r
        case "-":
            l -= r
        }
    }

    return l
}


func (p *Pars) stmt() { //statement
    name := p.current().Value
    for i := 0; i < 2; i++ {
        p.adv() // simultaniously advance and fix the bug where we forgot to skip '='
    }

    value := p.expr()
    vars[name] = value
}

