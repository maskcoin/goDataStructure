package calc

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"
	LPAREN   = "("
	RPAREN   = ")"
)

const (
	_ = iota //优先级
	LOWEST
	SUM     //+ -
	PRODUCT //* /
	PREFIX  //-1
	CALL    //1+(3)
)

var PRECEDENCES = map[string]int{
	PLUS:     SUM,
	MINUS:    SUM,
	ASTERISK: PRODUCT,
	SLASH:    PRODUCT,
	MOD:      PRODUCT,
	LPAREN:   CALL,
	RPAREN:   CALL,
}
