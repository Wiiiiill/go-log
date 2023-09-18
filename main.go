package main

import (
	"fmt"
	"log4go/log4go"
	"time"
)

func main() {
	fmt.Println("asd")
	for i := 0; i < 100; i++ {
		go log4go.Println(i)
	}
	for i := 0; i < 100; i++ {
		go log4go.Printf(" %d %d %d \n", i, i, i)
	}
	log4go.Println("4")
	time.Sleep(3000)
	fmt.Println(1, 2, 3)
}
