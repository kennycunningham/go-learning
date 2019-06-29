package main

import "fmt"

func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	*p = 10
	fmt.Println(v)
}
