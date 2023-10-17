package main

import (
	"log"
	"strconv"
)

type Expr interface {
	Eval() float64
}

const (
	LITERAL_Number = iota
	LITERAL_Input
)

type Literal struct {
	Value int64
	Type  uint
}

const (
	BINOP_ADDITION = iota
	BINOP_SUBTRACTION
	BINOP_MULTIPLICATION
	BINOP_DIVISION
	BINOP_RAISED
	BINOP_ROOT
	BINOP_EQUALS
)

type BinaryOperation struct {
	Left      Expr
	Operation uint
	Right     Expr
}

func TokensShift(tokens []Token) (Token, []Token) {
	return tokens[0], tokens[1:]
}

func Parse(tokens []Token) Expr {
	ast, tokens := ParseTokens(tokens)
	return ast
}

func ParseTokens(tokens []Token) (Expr, []Token) {
	if len(tokens) == 0 {
		log.Panicln("ERROR: No tokens to parse.")
	}

	token, remainingTokens := TokensShift(tokens)

	switch token.Type {
	case TOK_NUM:
		value, err := strconv.Atoi(token.Lexme)
		if nil != err {
			log.Fatalln(err)
		}
		return Literal{Value: int64(value), Type: LITERAL_Number}, remainingTokens
	case TOK_IN:
		return Literal{Value: 0, Type: LITERAL_Input}, remainingTokens
	case TOK_PLUS:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_ADDITION, Right: right}, remainingTokens
	case TOK_MINUS:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_SUBTRACTION, Right: right}, remainingTokens
	case TOK_MULTIPLICATION:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_MULTIPLICATION, Right: right}, remainingTokens
	case TOK_DIVISON:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_DIVISION, Right: right}, remainingTokens
	case TOK_RAISED:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_RAISED, Right: right}, remainingTokens
	case TOK_ROOT:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_ROOT, Right: right}, remainingTokens
	case TOK_EQUALS:
		left, remainingTokens := ParseTokens(remainingTokens)
		right, remainingTokens := ParseTokens(remainingTokens)
		return BinaryOperation{Left: left, Operation: BINOP_EQUALS, Right: right}, remainingTokens
	default:
		log.Panicln("ERROR: Unexpected token type.")
	}

	return nil, nil
}
