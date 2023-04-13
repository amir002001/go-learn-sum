package main

import "fmt"

func Switch() {
	// x := 2

	switch 1 {
	case 1:
		fmt.Println("hi")
		fallthrough
		// will fall through to next guy
	case 0:
		fmt.Println("hi2")
	default:
		fmt.Println("dude")
	}
}
