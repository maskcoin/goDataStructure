package calc

import "bytes"

type Expression interface {
	String() string
}

// 整数求值
type IntegerLiteralExpression struct {
	Token Token
	Value int64
}

func (il *IntegerLiteralExpression) String() string {
	return il.Token.Literal
}

// 前缀 1+ -1
type PrefixExpression struct {
	Token *Token
	Operator string
	Right Expression
}

// 括号内部计算
// (+1)
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.Write([]byte("("))
	out.Write([]byte(pe.Operator))
	out.Write([]byte(pe.Right.String()))
	out.Write([]byte(")"))
	return out.String()
}

type InfixExpression struct {
	Token Token
	Left Expression
	Operator string
	Right Expression
}

//(1+2)
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.Write([]byte("("))
	out.Write([]byte(ie.Left.String()))
	out.Write([]byte(" "))
	out.Write([]byte(ie.Operator))
	out.Write([]byte(" "))
	out.Write([]byte(ie.Right.String()))
	out.Write([]byte(")"))
	return out.String()
}
