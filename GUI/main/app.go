package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme()) //设置窗口 亮/暗 主题
	myWindow := myApp.NewWindow("超级无敌计算器 第二代")    //创建窗口 设置窗口名称
	myWindow.Resize(fyne.NewSize(300, 205))       //设置窗口大小
	myWindow.SetFixedSize(true)                   //设置窗口大小不可调整
	myWindow.CenterOnScreen()                     //设置窗口居中

	// 给窗口添加计算器图标
	iconResource, err := fyne.LoadResourceFromPath("GUI/image/calculatorIcon.png") // 确保路径正确
	if err != nil {
		fyne.LogError("icon加载失败", err)
		return
	}
	myWindow.SetIcon(iconResource)

	// 创建一个文本框用于显示计算结果
	display := widget.NewEntry() //新的输入框组件
	display.Text = "0"           //初始默认为0
	display.SetPlaceHolder("0")  //占位符为0
	display.Disable()            //禁用状态，禁止输入

	// 创建按钮
	btns := [][]string{
		{"7", "8", "9", "=", "<-"},
		{"4", "5", "6", "+", "-"},
		{"1", "2", "3", "*", "/"},
		{"0", ".", "C", "(", ")"},
	}

	// 创建按钮并设置点击事件
	grid := container.NewGridWithColumns(len(btns[0])) // 创建一个新的网格布局
	b := true
	for _, row := range btns {
		for _, label := range row {
			btn := widget.NewButton(label, func(label string) func() { //根据遍历，创建多个按钮
				return func() { // 在按钮被点击时执行函数
					switch label {
					case "C":
						display.SetText("0")
						b = true
					case "=":
						// 计算结果
						result := Calculate(display.Text)
						display.SetText(result)
						b = false
					case "<-":
						if b == false || len(display.Text) == 1 {
							display.SetText("0")
						} else if display.Text != "0" {
							display.SetText(display.Text[:len(display.Text)-1])
						}
						b = true
					default:
						if display.Text == "0" || b == false {
							display.SetText(label)
						} else {
							display.SetText(display.Text + label)
						}
						b = true
					}
				}
			}(label))
			grid.Add(btn)
		}
	}
	myWindow.SetContent(container.NewVBox(display, grid)) //将文本框和按钮添加到窗口

	myWindow.ShowAndRun()
}
