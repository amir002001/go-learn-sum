package main

import "fmt"

func while_loop() {

	i := 0
	for x := true; x; x = (i < 10) {

		fmt.Print("hello ooo")
		i++
	}
}
