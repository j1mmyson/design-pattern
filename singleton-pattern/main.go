package main

import (
	"fmt"
	"singleton-pattern/singleton"
)

func main() {
	for i := 0; i < 10; i++ {
		s := singleton.GetInstance()
		s.AddCount()
		fmt.Println(s.GetCount())
	}

	// Wait for the goroutines to finish
	fmt.Scanln()
}
