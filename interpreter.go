package main

import (
	"log"
	"math"
	"strconv"
)

func (n BinaryOperation) Eval() float64 {
	switch n.Operation {
	case BINOP_ADDITION:
		return n.Left.Eval() + n.Right.Eval()
	case BINOP_SUBTRACTION:
		return n.Left.Eval() - n.Right.Eval()
	case BINOP_MULTIPLICATION:
		return n.Left.Eval() * n.Right.Eval()
	case BINOP_DIVISION:
		return n.Left.Eval() / n.Right.Eval()
	case BINOP_RAISED:
		return math.Pow(n.Left.Eval(), n.Right.Eval())
	case BINOP_ROOT:
		return math.Pow(n.Right.Eval(), 1/n.Left.Eval())
	case BINOP_EQUALS:
		var val float64 = 0
		if n.Left.Eval() == n.Right.Eval() {
			val = 1
		}
		return val
	default:
		log.Panicln("WARNING: Operation not implemented! ( " + strconv.Itoa(int(n.Operation)) + " )")
	}
	return 0
}

func (n Literal) Eval() float64 {
	switch n.Type {
	case LITERAL_Number:
		return float64(n.Value)
	case LITERAL_Input:
		value, err := strconv.ParseFloat(ReadUserInput("\t> "), 64)
		if nil != err {
			log.Fatalln(err)
		}
		return value
	}
	return 0
	// c := sqrt(a^2 + b^2)
}
