package calc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//(a+1)*b
//a=1, b=2
//4

const (
	//四则运算符号
	AVAILABLE_OPERATOR_CODE = "+-*/^"
	//小数点符号
	AVAILABLE_DECIMAL_CODE = "0123456789.E"
	//参数
	AVAILABLE_PARAMETER_CODE = "abcdefghijklmnopqrstuvwxyz"
)

type Operator struct {
	sentence         string   //1+(2*3)
	operators        []string //表达式存储，2*3 1+6
	suffixExpression []string //后缀表达式
}

func NewOperator(sentence string) (op *Operator, err error) {
	op = &Operator{
		sentence: sentence,
	}
	op.Init()
	err = op.SetSuffixExpression()
	return
}

// 字符串切割
func (op *Operator) Init() {
	//type rune = int32
	value := make([]rune, 0, len(op.sentence))
	flag := false
	for _, c := range op.sentence {
		value = append(value, c)
		flag = true
		if c == rune('E') {
			flag = false
		}
	}
	fmt.Println(flag)
}

// 中缀表达式转化为后缀表达式，自动具备了顺序
func (op *Operator) SetSuffixExpression() (err error) {

	return
}

//计算结果
func (op *Operator) Execute() {

}

//判断操作符
func IsAvailableCode(s string) bool {
	return strings.IndexAny(AVAILABLE_OPERATOR_CODE, s) != -1
}

//判断小数
func IsBelongToDecimal(s string) bool {
	return strings.IndexAny(AVAILABLE_DECIMAL_CODE, s) != -1
}

//判断参数
func IsParameterCode(s string) bool {
	return strings.IndexAny(AVAILABLE_PARAMETER_CODE, s) != -1
}

//计算
func ExeSingleExpression(left, right float64, exp string) (ret float64) {
	switch exp {
	case "+":
		ret = left + right
	case "-":
		ret = left - right
	case "*":
		ret = left * right
	case "/":
		ret = left / right
	case "^":
		ret = math.Pow(left, right)
	}

	return
}

//(1+a)*b a=1 b=2
//参数替换
func ChangeParameter(parameter string, str []string) (ret string) {
	for i := 0; i < len(str); i += 2 {
		if parameter == str[i] {
			ret = str[i+1]
		}
	}
	return
}

func Modify2int64(i interface{}) (ret int64) {
	if i == nil {
		panic("Modify2int64(i interface{}) (ret int64)")
	}
	switch i.(type) {
	case int64:
		ret = i.(int64)
	case int32:
		ret = int64(i.(int32))
	case int16:
		ret = int64(i.(int16))
	case int8:
		ret = int64(i.(int8))
	case int:
		ret = int64(i.(int))
	case float64:
		ret = int64(i.(float64))
	case float32:
		ret = int64(i.(float32))
	case byte:
		ret = int64(i.(byte))
	default:
		ret = Parse2int64(Modify2string(i))
	}

	return
}

func Modify2string(i interface{}) (ret string) {
	if i == nil {
		panic("Modify2string")
	}
	switch i.(type) {
	case string:
		ret = i.(string)
	default:
		ret = fmt.Sprintf("%v", i)
	}
	return
}

func Parse2int64(s string, def ...int64) (ret int64) {
	if n, err := strconv.ParseInt(strings.TrimSpace(s), 0, 0); err == nil {
		ret = n
	} else if len(def) > 0 {
		ret = def[0]
	} else {
		panic("Parse2int64(s string, def... int64) (ret int64)")
	}
	return
}

func Modify2float64(i interface{}) (ret float64) {
	if i == nil {
		panic("Modify2int64(i interface{}) (ret int64)")
	}
	switch i.(type) {
	case int64:
		ret = float64(i.(int64))
	case int32:
		ret = float64(i.(int32))
	case int16:
		ret = float64(i.(int16))
	case int8:
		ret = float64(i.(int8))
	case int:
		ret = float64(i.(int))
	case float64:
		ret = i.(float64)
	case float32:
		ret = float64(i.(float32))
	case byte:
		ret = float64(i.(byte))
	default:
		ret = Parse2float64(Modify2string(i))
	}

	return
}

func Parse2float64(s string, def ...float64) (ret float64) {
	if n, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
		ret = n
	} else if len(def) > 0 {
		ret = def[0]
	} else {
		panic("Parse2int64(s string, def... int64) (ret int64)")
	}
	return
}
