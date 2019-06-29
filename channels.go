package main

import "fmt"

func insertWord(word string, c chan string) {
	c <- word
}

func main() {
	go1 := make(chan string)
	arrayWeGo := [6]string{"dog", "cat", "mice", "alice", "in", "chains"}
	var x []string
	var y []string
	for _, val := range arrayWeGo {
		go insertWord(val, go1)
		if len(x) < len(y) {
			x = append(x, <-go1)
		} else {
			y = append(y, <-go1)
		}
	}
	fmt.Println(x)
	fmt.Println(y)
}
