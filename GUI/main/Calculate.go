package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func infixToPostfix(infix string) string {
	var postfix strings.Builder
	var stack []rune

	//for _, token := range infix {
	for i := 0; i < len(infix); i++ {
		token := rune(infix[i])
		switch {
		case token >= '0' && token <= '9' || token == '.':
			for i < len(infix) {
				token = rune(infix[i])
				if token >= '0' && token <= '9' || token == '.' {
					postfix.WriteRune(token)
					i++
				} else {
					break
				}
			}
			i--
			postfix.WriteRune(' ')
		case token == '(':
			stack = append(stack, token)
		case token == ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		default:
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		postfix.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix.String()
}

func evaluatePostfix(postfix string) (float64, error) {
	var stack []float64 // 创建一个栈来存储操作数

	for i := 0; i < len(postfix); i++ {
		token := rune(postfix[i])
		switch {
		case token >= '0' && token <= '9' || token == '.':
			var temp strings.Builder
			for i < len(postfix) {
				token = rune(postfix[i])
				if token >= '0' && token <= '9' || token == '.' {
					temp.WriteRune(token)
					i++
				} else {
					break
				}
			}
			operand, _ := strconv.ParseFloat(temp.String(), 64)
			stack = append(stack, operand)
		case token == '+', token == '-', token == '*', token == '/':
			if len(stack) < 2 { // 检查栈中是否至少有两个操作数
				return 0, fmt.Errorf("后缀表达式无效")
			}
			// 从栈中弹出两个操作数
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 执行运算
			switch token {
			case '+':
				stack = append(stack, operand1+operand2)
			case '-':
				stack = append(stack, operand1-operand2)
			case '*':
				stack = append(stack, operand1*operand2)
			case '/':
				if operand2 == 0 { // 检查除数是否为零
					return 0, fmt.Errorf("除数为零")
				}
				stack = append(stack, operand1/operand2)
			}
		}
	}

	if len(stack) != 1 { // 检查栈中是否只剩下一个元素
		return 0, fmt.Errorf("后缀表达式无效")
	}

	return stack[0], nil
}

func remove(ans float64) string {
	result := fmt.Sprintf("%f", ans)

	for {
		if result[len(result)-1] == '0' {
			result = result[:len(result)-1]
		} else if result[len(result)-1] == '.' {
			result = result[:len(result)-1]
			return result
		} else {
			return result
		}
	}
}

func Calculate(infix string) string {
	re := regexp.MustCompile(`\s+`)
	infix = re.ReplaceAllString(infix, "")
	if infix == "" {
		return "error"
	}

	postfix := infixToPostfix(infix)

	ans, err := evaluatePostfix(postfix)
	if err != nil {
		return "error"
	}

	result := remove(ans)
	return result
}
