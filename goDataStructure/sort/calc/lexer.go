package calc

type Lexer struct {
	Input        string //输入字符串
	Position     int    //位置
	ReadPosition int    //读取位置
	Ch           byte   //读取的一个字节
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		Input: input,
	}

	l.ReadChar() // 前进一个字符
	return l
}

// 分离数据，操作
func (l *Lexer) NextToken() *Token {
	var token *Token
	l.SkipWhiteSpace() //跳过垃圾字符
	switch l.Ch {
	case '(':
		token = NewToken(LPAREN, l.Ch)
	case ')':
		token = NewToken(RPAREN, l.Ch)
	case '+':
		token = NewToken(PLUS, l.Ch)
	case '-':
		token = NewToken(MINUS, l.Ch)
	case '*':
		token = NewToken(ASTERISK, l.Ch)
	case '/':
		token = NewToken(SLASH, l.Ch)
	case '%':
		token = NewToken(MOD, l.Ch)
	case '0':
		token.Type = EOF
		token.Literal = ""
	default:
		if IsDigit(l.Ch) {
			token.Type = INT
			token.Literal = l.ReadNumber()
			return token
		} else {
			token = NewToken(ILLEGAL, l.Ch) //非法字符
		}
	}
	l.ReadChar()
	return token
}

//提取一个字符
func (l *Lexer) ReadChar() {
	if l.ReadPosition >= len(l.Input) {
		l.Ch = '0'
	} else {
		l.Ch = l.Input[l.ReadPosition]
	}
	l.Position = l.ReadPosition
	l.ReadPosition += 1
}

//123+2,切割数字出来
func (l *Lexer) ReadNumber() string {
	pos := l.Position
	for IsDigit(l.Ch) {
		l.ReadChar() //连续提取数字
	}
	return l.Input[pos:l.Position]
}

func (l *Lexer) SkipWhiteSpace() {
	if l.Ch == '\t' || l.Ch == ' ' || l.Ch == '\r' || l.Ch == '\n' {
		l.ReadChar()
	}
}
