package main

import (
	"log"
	"unicode"
)

// Token Type
const (
	TOK_NUM = iota
	TOK_PLUS
	TOK_MINUS
	TOK_MULTIPLICATION
	TOK_DIVISON
	TOK_RAISED
	TOK_ROOT
	TOK_IN
	TOK_EQUALS
)

type Token struct {
	Type  uint
	Lexme string
}

func TokenizeString(str string) []Token {
	var tokens []Token
	i := 0
	for i < len(str) {
		if unicode.IsSpace(rune(str[i])) {
			i++
			continue
		}
		switch rune(str[i]) {
		case '+':
			tokens = append(tokens, Token{Type: TOK_PLUS, Lexme: "+"})
		case '-':
			tokens = append(tokens, Token{Type: TOK_MINUS, Lexme: "-"})
		case '*':
			tokens = append(tokens, Token{Type: TOK_MULTIPLICATION, Lexme: "*"})
		case '/':
			tokens = append(tokens, Token{Type: TOK_DIVISON, Lexme: "/"})
		case '^':
			tokens = append(tokens, Token{Type: TOK_RAISED, Lexme: "^"})
		case '\\':
			tokens = append(tokens, Token{Type: TOK_ROOT, Lexme: "\\"})
		case '$':
			tokens = append(tokens, Token{Type: TOK_IN, Lexme: "$"})
		case '=':
			tokens = append(tokens, Token{Type: TOK_EQUALS, Lexme: "="})
		default:
			if unicode.IsDigit(rune(str[i])) {
				lexme := ""
				for ; i < len(str) && unicode.IsDigit(rune(str[i])); i++ {
					lexme += string(str[i])
				}

				i--

				tokens = append(tokens, Token{Type: TOK_NUM, Lexme: lexme})
			} else {
				log.Fatalln("ERROR: Unknown character found! ( " + string(str[i]) + " )")
			}
		}
		i++
	}
	return tokens
}
