package main

import (
	"fmt"
	"strconv"
)

func StrConv() {
	dude, err := strconv.Atoi("i323")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(dude)
}
