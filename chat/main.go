package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 预设的问题和答案
	answers := map[string]string{
		"你叫什么名字？":      "我是一个聊天机器人，你可以随时给我取一个名字。",
		"你是谁开发的？":      "我是由三月软件团队开发的，三月软件致力于为用户提供更好的交互体验。",
		"你会说多少种语言？":    "我是智障，现在只会讲中文",
		"你能告诉我今天的天气吗？": "当然，我可以帮你查询今天的天气情况。",
		"你会做饭吗？":       "抱歉，我只是一个聊天机器人，不会做饭。",
		"你喜欢什么样的音乐？":   "我是一个机器人，没有自己的喜好，但我可以帮你推荐各种类型的音乐。？",
		"你知道如何学习吗？":    "我是一个机器学习的聊天机器人，可以根据用户的反馈不断学习和改进。",
		"你会玩游戏吗？":      "我可以和你玩一些简单的文字游戏。",
		"你知道如何做运动吗？":   "我不会做运动，但我可以帮你提供一些健身建议。",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("你有什么问题想问我吗？(输入 q 退出) ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if question == "q" {
			fmt.Println("再见！")
			break
		}

		answer, ok := answers[question]
		if ok {
			fmt.Println("回答：", answer)
		} else {
			fmt.Println("回答：我还在学习中，您的问题我回答不了")
		}
	}
}
