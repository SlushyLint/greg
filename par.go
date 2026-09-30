package main

import (
	"fmt"
	"strconv"
)

// Value stores a runtime value. Type determines which value field is in use.
type Value struct {
	Type  string
	Int   int
	Float float64
	Str   string
}

// vars contains values declared or assigned while the current program runs.
var vars = make(map[string]Value)

// Pars holds the token stream and the index of the next token to parse.
type Pars struct {
	tokens []Token
	pos    int
}

// par creates a parser positioned at the first token.
func par(tokens []Token) *Pars {
	return &Pars{
		tokens: tokens,
		pos:    0,
	}
}

// adv returns the current token and advances the parser, or an empty token at EOF.
func (p *Pars) adv() Token {
	if p.pos >= len(p.tokens) {
		return Token{}
	}
	token := p.tokens[p.pos]
	p.pos++
	return token
}

// num reads a numeric literal or numeric variable as a float64.
func (p *Pars) num() float64 {
	token := p.current()
	if token.Type == "paren" { // recursive math
		p.adv()           // we have to accept the parenthesis
		value := p.expr() // main math func
		if p.current().Type != "paren" ||
			p.current().Value != "(" {
			panic("[!] E: Expected )")
		}
		p.adv() //consume closing parenthesis
		return value
	}

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

		// Non-numeric or unknown identifiers currently evaluate as zero.
		return 0
	}
	// Other token types, including strings, are not numeric expressions.
	return 0
}

// current returns the next token without consuming it, or an empty token at EOF.
func (p *Pars) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{}
	}
	return p.tokens[p.pos]
}

// term parses multiplication and division, which bind more tightly than + and -.
func (p *Pars) term() float64 {
	l := p.num()

	for p.pos < len(p.tokens) {
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

// expr parses addition and subtraction over terms.
func (p *Pars) expr() float64 {
	l := p.term()

	for p.pos < len(p.tokens) {
		op := p.current().Value

		if op != "+" && op != "-" {
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

func (p *Pars) whilestmt() {
	p.adv()
	condstart := p.pos
	matches := p.condition()
	bodystart := p.pos
	bodyend := p.blockend(bodystart)

	for matches {
		p.pos = bodystart
		p.block(true)
		p.pos = condstart
		matches = p.condition()

	}
	p.pos = bodyend
}

func (p *Pars) blockend(start int) int {
	if start >= len(p.tokens) ||
		p.tokens[start].Type != "brace" ||
		p.tokens[start].Value != "{" {
		panic("[!] E: Expected {")

	}
	depth := 0
	for index := start; index < len(p.tokens); index++ {
		token := p.tokens[index]
		if token.Type != "brace" {
			continue
		}
		if token.Value == "{" {
			depth++
		} else if token.Value == "}" {
			depth--
			if depth == 0 {
				return index + 1

			}
		}
	}
	panic("[!] E: Expected }")
}

// stmt parses and executes one declaration, print, assignment, or conditional.
func (p *Pars) stmt() {
	if p.current().Value == "if" {
		p.ifstmt()
		return
	}
	if p.current().Value == "for" {
		p.whilestmt()
		return
	}
	tname := p.current().Value

	if p.pos+1 < len(p.tokens) && p.tokens[p.pos+1].Type == "dc" {
		// A declaration has the form type ^ name = value.
		p.adv() //type
		p.adv()
		name := p.adv().Value
		p.adv() // =
		value := p.expr()

		if tname == "i" {
			// Integer declarations truncate the evaluated numeric value.
			vars[name] = Value{
				Type: "int",
				Int:  int(value),
			}
		}
		if tname == "f" {
			vars[name] = Value{
				Type:  "float",
				Float: (value),
			}
		}

		if tname == "s" {
			// Expressions do not consume string tokens, so read this initializer directly.
			value := p.str()
			vars[name] = Value{
				Type: "str",
				Str:  string(value),
			}
		}
		return
	}

	name := p.current().Value
	if name == "print" {
		// Print accepts a literal token or a previously stored variable.
		p.adv()
		if p.current().Type != "paren" ||
			p.current().Value != "(" {
			panic("[!] E: Expected (")

		}
		p.adv()
		item := p.current() // current position

		switch {
		case item.Type == "s": // string litterals in "" and plain ass numbers
			fmt.Println(p.adv().Value)
		case item.Type == "l":
			printval(item.Value) // usually a variable name
		case item.Type == "n" || item.Type == "l":
			fmt.Println(p.expr())

		default: // error handling;
			// just printing newline should be fine here too
			panic("[!] E: expected a statement")
		}
		if p.current().Type != "paren" ||
			p.current().Value != ")" {
			panic("[!] E: Expected )")
		}
		p.adv()
		return
	}
	// Skip the assignment target and '=' before evaluating the right-hand side.
	for i := 0; i < 2; i++ {
		p.adv()
	}

	value := p.expr()
	// Reassignment currently represents every value as a float.
	vars[name] = Value{
		Type:  "float",
		Float: value,
	}
}

// str reads and consumes the current string token, or returns an empty string.
func (p *Pars) str() string {
	token := p.current()

	if token.Type == "s" {
		p.adv()
		return token.Value
	}
	return ""
}

// printval prints a variable according to the type stored in vars.
func printval(name string) {
	value := vars[name]
	switch value.Type {
	case "int":
		fmt.Println(value.Int)

	case "float":
		fmt.Println(value.Float)

	case "str":
		fmt.Println(value.Str)
	}
}

// String formats a Value for Go's stringer interface.
func (v Value) String() string {
	switch v.Type {
	case "int":
		return fmt.Sprintf("%d", v.Int)
	case "float":
		return fmt.Sprintf("%v", v.Float)
	case "str":
		return v.Str
	}

	return ""
}
