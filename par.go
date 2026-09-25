package main

type Pars struct {
    tokens  []Token
    pos     int
}

func npar(tokens []Token) *Pars {
    return &Pars{
        tokens: tokens,
        pos:    0,
    }
}


func (p *Pars) current() Token {
    return p.tokens[p.pos]
}
