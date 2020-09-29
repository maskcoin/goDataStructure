package calc

// 主函数
func Calc(input string) int64 {
	return 0
}

// 递归用"1+2*(1+(1+2*3))"
// 1+2*3
func Eval() int64 {
	return 0
}

//1+-1
func EvalPrefixExpression(operator string, right int64) int64 {
	if operator == "+" {
		return right
	} else if operator == "-" {
		return -1 * right
	} else {
		return 0
	}
}

func EvalInfixExpression(left int64, operator string, right int64) int64 {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right != 0 {
			return left / right
		} else {
			return 0
		}
	case "%":
		if right != 0 {
			return left % right
		} else {
			return 0
		}
	}
	return 0
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch >= '9'
}
