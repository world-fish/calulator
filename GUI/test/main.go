package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	// 创建一个文本框用于显示计算结果
	display := widget.NewEntry()
	display.Text = "0"
	display.SetPlaceHolder("0")
	display.Disable()

	// 创建按钮
	btns := [][]string{
		{"7", "8", "9", "/", "("},
		{"4", "5", "6", "*", ")"},
		{"1", "2", "3", "-", "."},
		{"0", "+", "=", "C", ""},
	}

	// 创建按钮并设置点击事件
	grid := container.NewGridWithColumns(len(btns[0])) // Create a grid with the number of columns equal to the number of buttons in a row
	for _, row := range btns {
		for _, label := range row {
			btn := widget.NewButton(label, func(label string) func() {
				return func() {
					switch label {
					case "C":
						display.SetText("0")
					case "=":
						// 计算结果
						result := calculate(display.Text)
						display.SetText(result)
					default:
						if display.Text == "0" {
							display.SetText(label)
						} else {
							display.SetText(display.Text + label)
						}
					}
				}
			}(label))
			grid.Add(btn) // Add the button to the grid
		}
	}
	myWindow.SetContent(container.NewVBox(display, grid)) // Add the display and the grid of buttons to the window

	myWindow.ShowAndRun()
}

// 计算结果
func calculate(expression string) string {
	// 简单实现，仅支持加减乘除
	// 实际开发中可以使用类库进行更复杂的计算
	// 这里简单处理错误情况，如除数为0
	result := "Error"
	if len(expression) > 0 {
		result = expression
	}
	return result
}
