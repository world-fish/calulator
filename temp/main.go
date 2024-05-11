package main

import (
	"bufio"
	"fmt"
	"os"
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

	for _, token := range infix {
		switch {
		case token >= '0' && token <= '9':
			postfix.WriteRune(token)
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

	for _, token := range postfix {
		switch {
		case token >= '0' && token <= '9': // 如果 token 是操作数
			operand, _ := strconv.ParseFloat(string(token), 64) // 将字符串转换为浮点数
			stack = append(stack, operand)                      // 将操作数压入栈中
		case token == '+', token == '-', token == '*', token == '/': // 如果 token 是运算符
			if len(stack) < 2 { // 检查栈中是否至少有两个操作数
				return 0, fmt.Errorf("无效的后缀表达式")
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
		return 0, fmt.Errorf("无效的后缀表达式")
	}

	return stack[0], nil // 返回结果
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入中缀表达式:")
	infix, _ := reader.ReadString('\n')
	re := regexp.MustCompile(`\s+`)
	infix = re.ReplaceAllString(infix, "")

	postfix := infixToPostfix(infix)
	fmt.Println("后缀表达式:", postfix)

	result, err := evaluatePostfix(postfix)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}
}
