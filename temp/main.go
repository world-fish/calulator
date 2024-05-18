package main

import "fmt"

type neme struct {
	a int
	b string
}

func (n *neme) fun() {

	n.a = 100

}

func aaa(sle []int) {
	//sle = append(sle, 10)
	sle = append(sle, 1)
	fmt.Println(len(sle), cap(sle))

	return
}

func main() {
	nn := neme{
		10,
		"234",
	}
	nn.fun()
	fmt.Println(nn)
}
