package main

// Token is one piece of source recognized by the lexer. Type identifies the
// category and Value stores the source text when applicable.
type Token struct {
	Type  string
	Value string
}

// isdigit reports whether c is an ASCII decimal digit.
func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// isquote reports whether c is a double quote.
func isquote(c byte) bool {
	return c == '"'
}

// isletter reports whether c is allowed in an identifier.
func isletter(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

// lex converts source text into tokens and emits indentation changes as tokens.
func lex(src string) []Token {
	var tokens []Token

	// Each entry is an indentation width; the last entry is the active level.
	indents := []int{0}
	lnstart := true

	for i := 0; i < len(src); i++ {
		if lnstart {
			indent := 0
			for i < len(src) && src[i] == ' ' {
				indent++
				i++
			}

			if i >= len(src) {
				break
			}

			if src[i] == '\n' {
				tokens = append(tokens, Token{
					Type:  "nl", // newline
					Value: "\n",
				})
				lnstart = true
				continue
			}

			curindent := indents[len(indents)-1]
			if indent > curindent {
				indents = append(indents, indent)
				tokens = append(tokens, Token{
					Type: "indent",
				})
			} else {
				for indent < curindent {

					indents = indents[:len(indents)-1]
					curindent = indents[len(indents)-1]
					tokens = append(tokens, Token{
						Type: "dedent",
					})
				}
				if indent != curindent {
					panic("[!] E: Inconsistent indentation")
				}
			}
			lnstart = false
			// Revisit the first non-space byte in the main scanning section.
			i--
			continue
		}

		if isquote(src[i]) {
			// Strings end at the next quote; escape sequences are not processed.
			start := i + 1
			i++
			for i < len(src) && !isquote(src[i]) {
				i++
			}
			if i >= len(src) {
				break
			}

			tokens = append(tokens, Token{
				Type:  "s",
				Value: src[start:i],
			})
			continue
		}
		if isletter(src[i]) {
			// Identifiers contain ASCII letters and underscores.
			start := i
			for i < len(src) && isletter(src[i]) {
				i++
			}
			tokens = append(tokens, Token{
				Type:  "l",
				Value: src[start:i],
			})
			i--
			continue
		}
		if isdigit(src[i]) {
			// Number tokens may contain a decimal point.
			start := i
			for i < len(src) && isdigit(src[i]) || src[i] == '.' {
				i++
			}

			tokens = append(tokens, Token{
				Type:  "n",
				Value: src[start:i],
			})
			i--
			continue
		}

		if src[i] == '\n' {
			tokens = append(tokens, Token{
				Type:  "nl",
				Value: "\n",
			})
			lnstart = true
			continue
		}
		if src[i] == ' ' || src[i] == '\t' {
			continue
		}

		// Recognize single-character operators, the declaration marker, and braces.
		switch src[i] {
		case '+':
			value := "+"
			if i+1 < len(src) && src[i+1] == '+' { // "++" for incrimenting variables
				value = "++" // useful, as we're gonna use a basic c 'for' loop
				i++

			}
			tokens = append(tokens, Token{
				Type:  "op",
				Value: value,
			})
		case '-':
			value := "-"
			if i+1 < len(src) && src[i+1] == '-' {
				value = "--"
				i++
			}
			tokens = append(tokens, Token{
				Type:  "op",
				Value: value,
			})
		case '*':
			tokens = append(tokens, Token{
				Type:  "op",
				Value: "*",
			})
		case '/':
			tokens = append(tokens, Token{
				Type:  "op",
				Value: "/",
			})
		case '=':
			tokens = append(tokens, Token{
				Type:  "op",
				Value: "=",
			})
		case ':':
			tokens = append(tokens, Token{
				Type:  "op",
				Value: ":",
			})
		case ';':
			tokens = append(tokens, Token{
				Type:  "op",
				Value: ";",
			})
		case '^':
			tokens = append(tokens, Token{
				Type:  "dc", // declaration
				Value: "^",
			})
		case '>':
			tokens = append(tokens, Token{
				Type:  "op", //opperator
				Value: ">",
			})
		case '<':
			tokens = append(tokens, Token{
				Type:  "op", //operator
				Value: "<",
			})
		case '{', '}':
			tokens = append(tokens, Token{
				Type:  "brace",
				Value: string(src[i]),
			})
		case '(', ')':
			tokens = append(tokens, Token{
				Type:  "paren",
				Value: string(src[i]),
			})
		case ',':
			tokens = append(tokens, Token{
				Type:  "c",
				Value: string(src[i]),
			})

		}
	}
	// Close any indentation levels left open at end of input.
	for len(indents) > 1 {
		indents = indents[:len(indents)-1]
		tokens = append(tokens, Token{
			Type: "dedent",
		})
	}

	return tokens
}
