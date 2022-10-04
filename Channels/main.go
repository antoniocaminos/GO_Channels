package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

/*
	func main() {
		go DoSomething("hello, world!")
		fmt.Println("another message")
		for {
			//do nothing
		}
	}

	func DoSomething(s string) {
		until := 0
		for {
			fmt.Println("s is", s)
			until = until + 1
			if until >= 5 {
				break
			}
		}

}
*/
var keyPresChan chan rune

func main() {
	keyPresChan = make(chan rune)
	go listenForKeyPress()
	fmt.Println("press any key or Q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()
	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPresChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPresChan
		fmt.Println("you press", string(key))
	}
}
