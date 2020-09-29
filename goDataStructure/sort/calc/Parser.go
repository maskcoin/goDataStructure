package calc

import "fmt"

type (
	PrefixParseFn func() Expression
	InfixParseFn  func(Expression) Expression
)

type Parser struct {
	L         *Lexer
	CurToken  *Token // 当前的
	PeekToken *Token

	//解析
	PrefixParseFns map[string]PrefixParseFn
	InfixParseFns  map[string]InfixParseFn

	//处理错误
	errors []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		L:              l,
		PrefixParseFns: make(map[string]PrefixParseFn),
		InfixParseFns:  make(map[string]InfixParseFn),
	}

	p.NextToken()
	p.NextToken()

	return p
}

//map中插入数据
func (p *Parser) RigisterPrefix(tokenType string, fn PrefixParseFn) {
	p.PrefixParseFns[tokenType] = fn
}

func (p *Parser) RigisterInfix(tokenType string, fn InfixParseFn) {
	p.InfixParseFns[tokenType] = fn
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) NextToken() {
	p.CurToken = p.PeekToken
	p.PeekToken = p.L.NextToken()
}

// 处理解析错误
func (p *Parser) PeekError(t string) {
	msg := fmt.Sprintf("bug %s %s", t, p.PeekToken.Type)
	p.errors = append(p.errors, msg)
}

// 取出优先级
func (p *Parser) PeekPrecedence() int {
	if p, ok := PRECEDENCES[p.PeekToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

func (p *Parser) CurPrecedence() int {
	if p, ok := PRECEDENCES[p.CurToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

func (p *Parser) PeekTokenIs(t string) bool {
	return p.PeekToken.Type == t
}

func (p *Parser) ExpectPeek(t string) bool {
	if p.PeekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.PeekError(t)
		return false
	}
}

func (p *Parser) ParseExpression(precedence int) Expression {
	prefix := p.PrefixParseFns[p.CurToken.Type]
	retExp := prefix()

	for precedence < p.PeekPrecedence() {
		infix := p.InfixParseFns[p.PeekToken.Type]
		if infix == nil {
			return retExp
		}
		p.NextToken()
		retExp = infix(retExp)
	}

	return retExp
}

func (p *Parser) ParsePrefixExpression() Expression {
	exp := &PrefixExpression{
		Token:    p.CurToken,
		Operator: p.CurToken.Literal,
		Right:    nil,
	}
	p.NextToken()

	return exp
}
