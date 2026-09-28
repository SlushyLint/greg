package main

import(
    "fmt"
    "strconv"
)

// var  vars = make(map[string]float64) this  will be a huge issue when we worry about var conversions
// lets make a struct

type Value struct {
    Type    string
    Int     int
    Float   float64
    Str     string
}
var vars = make(map[string]Value)

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
        value := vars[token.Value]
        p.adv()
    
        if value.Type == "int" {
            return float64(value.Int)
        }
        if value.Type == "float" {
            return value.Float
        }

        return 0
    }
    return 0    // returns zero
}

func (p *Pars) current() Token {
    if p.pos >= len(p.tokens) {
        return Token{}
    }
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
    tname := p.current().Value


    if p.pos+1 < len(p.tokens) && p.tokens[p.pos+1].Type == "dc" {
        p.adv()     //type
        p.adv()     // ^
        name := p.adv().Value
        p.adv()     // =
        value := p.expr()
//        fmt.Println("\ntype:", tname, "\name:", name)

        if tname == "i" {
            vars[name] = Value{
                Type:   "int",
                Int:    int(value),
            }
        }
        if tname == "f" {
            vars[name] = Value{
                Type:   "flt",
                Float:  (value),
            }
        }

        if tname == "s" {
            value := p.str()    // we cant convert an f64 to a str as easily, 
                                // so we'll reuse a function from earlier
            vars[name] = Value{
                Type:   "str",
                Str: string(value),
            }
        }
        return
    }
    
    name := p.current().Value
    if name == "print" {
        p.adv()
        value := p.expr()
        fmt.Println(value)
        return
    }
    for i := 0; i < 2; i++ {
        p.adv() // simultaniously advance and fix the bug where we forgot to skip '='
    }

    value := p.expr()
    vars[name] = Value{
        Type: "float",
        Float: value,
    }
//    value := p.expr()
//    vars[name] = value
}


func (p *Pars) str() string { // this returns a string, right...
    token := p.current()

    if token.Type == "s" {
        p.adv()
        return token.Value
    }
    return "" // ... here
}
