package main

import (
	"fmt"
)

func main() {
	var str = "aa"
	fmt.Println(&str)
	hello(&str)
}

func hello(str *string) {
	fmt.Println(&*str)
}
