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

	//for _, token := range infix {
	for i := 0; i < len(infix); i++ {
		token := rune(infix[i])
		switch {
		case token >= '0' && token <= '9' || token == '.':
			for i < len(infix) {
				token = rune(infix[i])
				if token >= '0' && token <= '9' || token == '.' || token == '!' {
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
		case token >= '0' && token <= '9' || token == '.' || token == '!':
			nega := 1.0
			if token == '!' {
				nega = -1.0
				i++
			}
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
			operand *= nega
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

// 判断负数
func negative(infix string) string {
	infix = strings.Replace(infix, "(-", "(!", -1)
	if infix[0] == '-' {
		infix = "!" + infix[1:]
	}
	return infix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\s+`)
	for {
		fmt.Print("请输入算数表达式:")
		infix, _ := reader.ReadString('\n')
		infix = re.ReplaceAllString(infix, "")
		if infix == "" {
			break
		}

		infix = negative(infix)
		fmt.Println(infix)

		postfix := infixToPostfix(infix)
		fmt.Println("后缀表达式:", postfix)

		result, err := evaluatePostfix(postfix)
		if err != nil {
			fmt.Println("错误:", err)
		} else {
			fmt.Printf("结果:%.2f", result)
		}

		fmt.Println()
	}
}
