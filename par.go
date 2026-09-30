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
	//	if token.Type == "paren" { // recursive math
	//		p.adv()           // we have to accept the parenthesis
	//		value := p.expr() // main math func
	//		if p.current().Type != "paren" ||
	//			p.current().Value != "(" {
	//			panic("[!] E: Expected )")
	//		}
	//		p.adv() //consume closing parenthesis
	//		return value
	//	}

	if token.Type == "paren" && token.Value == "(" {
		p.adv() // we have to grab the "("
		value := p.expr()
		if p.current().Type != "paren" ||
			p.current().Value != ")" {
			panic("[!] E: Expected )")

		}
		p.adv() //continue
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

		// non-numbers eval as zero
		// deal with it
		return 0
	}
	// other token types,like strings, are not intigers or floats, incase you are stupid

	return 0 // <-- returns zero
}

// current() returns the next token, then...
// uh idk
// but it does something useful.
func (p *Pars) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{}
	}
	return p.tokens[p.pos]
}

// term parses multiplication and division, which bind more tightly than + and -.

func (p *Pars) term() float64 { // parses * and / more strictly than + and -
	// safety float64
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
func (p *Pars) expr() float64 { //yes, this shitshow will always return a float
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

func (p *Pars) forstmt() {
	// c-style looping
	// example, for (i^i = 0; i < 20; i++) to do something 20 times
	p.adv() // read "for"
	if p.current().Type != "paren" || p.current().Value != "(" {
		panic("[!] L: Expected ( after 'for'")
	}
	p.adv() //read (

	p.stmt() //initializer
	// e.g., i^i = 0
	if p.current().Value != ";" {
		panic("[!] L: Expected ; after initializer")
	}
	p.adv() //continue

	constart := p.pos
	matches := p.condition()
	if p.current().Value != ";" {
		panic("[!] L: Expected ; after condition")
	}
	p.adv() //continue

	updstart := p.pos
	closeparen := updstart
	for closeparen < len(p.tokens) &&
		!(p.tokens[closeparen].Type == "paren" &&
			p.tokens[closeparen].Value == ")") {
		closeparen++
	}

	if closeparen >= len(p.tokens) {
		panic("[!] L: Expected ) after update")

	}
	bodystart := closeparen + 1 // this plus one pissed me off
	bodyend := p.blockend(bodystart)

	for matches {
		p.pos = bodystart
		p.block(true)

		p.pos = updstart
		p.stmt() // update, example x++

		p.pos = constart
		matches = p.condition()

	}

	p.pos = bodyend
}

func (p *Pars) whilestmt() {
	// ?????????
	// i made this at 3am and it works??
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
		panic("[!] E: Expected {") // you missed the starting brace

	}
	depth := 0 // we should reset the depth for every brace for recursion
	for index := start; index < len(p.tokens); index++ {
		token := p.tokens[index]
		if token.Type != "brace" {
			continue
		}
		if token.Value == "{" { // ALE, stfu
			depth++
		} else if token.Value == "}" {
			depth--
			if depth == 0 {
				return index + 1

			}
		}
	}
	panic("[!] E: Expected }") // in other words, "shit"
}

// stmt parses and executes one declaration, print, assignment, or conditional.
func (p *Pars) stmt() {

	if p.current().Value == "if" {
		p.ifstmt()
		return
	}
	if p.current().Value == "for" {
		p.forstmt()
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
			// why the hell did i make float the default variable????
			vars[name] = Value{
				Type:  "float",
				Float: (value),
			}
		}

		if tname == "s" {
			// Expressions do not read string tokens
			// so read the bullshit directly
			value := p.str()
			vars[name] = Value{
				Type: "str",
				Str:  string(value),
			}
		}
		return
	}

	// print and  printf should do the same thing
	name := p.current().Value
	if name == "print" || name == "printf" {
		p.printstmt()
		return
	}

	// now this is an actual peice of shit
	// deadass forgot what this does
	if p.pos+1 < len(p.tokens) &&
		p.current().Type == "l" &&
		p.tokens[p.pos+1].Type == "op" &&
		(p.tokens[p.pos+1].Value == "++" ||
			p.tokens[p.pos+1].Value == "--") {
		if p.invars() {
			return
		}
	}

	if !p.invars() {
		return
	}
	// Skip the assignment target and '=' before evaluating the side.
	for i := 0; i < 2; i++ { // shut up ALE
		p.adv()
	}

	value := p.expr()
	//  represents every god damn value as a float because why tf not
	vars[name] = Value{
		Type:  "float",
		Float: value,
	}
}

func (p *Pars) invars() bool {
	if p.pos+1 >= len(p.tokens) ||
		p.current().Type != "l" {
		return false
	}
	name := p.current().Value // p global current token value
	op := p.tokens[p.pos+1].Value
	if op != "++" && op != "--" {
		return false
	}
	value, ok := vars[name]
	if !ok {
		panic("[!] V: Unknown variable")
	}
	delta := 1
	if op == "--" {
		delta = -1

	}

	switch value.Type {

	case "int":
		value.Int += delta

	case "float":
		value.Float += float64(delta)

	default:
		panic("[!] V: Can only incriment i^ or ^f vars")
		// i^ variables are int
		// f^ variables are floats (specifically float64)
	}
	vars[name] = value
	p.adv()
	p.adv()
	return true
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
		return fmt.Sprintf("%d", v.Int) // why is printf so fucking weird
	case "float":
		return fmt.Sprintf("%v", v.Float) // THIS is sensible
	case "str":
		return v.Str
	}

	return "" // we have to return a string
}

// printstmt parses print/printf(format, arguments) and writes the formatted output.

func (p *Pars) printstmt() {
	p.adv() // read "print"
	if p.current().Type != "paren" || p.current().Value != "(" {
		panic("[!] P: Expected ( after print")
	}
	p.adv() //read (

	formtoken := p.adv()
	if formtoken.Type != "s" {
		panic("[!] FM: Expected a format string")
	}
	format := formtoken.Value
	args := make([]any, 0)

	for p.current().Type == "c" {
		p.adv()
		token := p.current()

		switch token.Type {
		case "s":
			args = append(args, p.adv().Value)
		case "l":
			p.adv()
			value, ok := vars[token.Value]
			if !ok {
				panic("[!] PV: Unknown variable in print")
			}
			switch value.Type {
			case "int":
				args = append(args, value.Int)
			case "float":
				args = append(args, value.Float)
			case "str":
				args = append(args, value.Str)
			default:
				panic("[!] PV: Unsupported variable type in print")
			}
		case "n", "paren":
			args = append(args, p.expr())
		default:
			panic("[!] P: Expected print argument")
		}
	}
	if p.current().Type != "paren" || p.current().Value != ")" {
		panic("[!] P: Expected ) after print args")
	}
	p.adv()
	fmt.Printf(format, args...)
}
