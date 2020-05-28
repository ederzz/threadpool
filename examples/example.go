package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	i, ch, scan := 0, make(chan string), bufio.NewScanner(os.Stdin)
	// goroutine1
	go func() {
		for {
			fmt.Print(<-ch)
			fmt.Println("goroutine1")
		}
	}()

	// goroutine2
	go func() {
		for {
			fmt.Print(<-ch)
			fmt.Println("goroutine2")
		}
	}()

	// goroutine3
	go func() {
		for {
			fmt.Print(<-ch)
			fmt.Println("goroutine3")
		}
	}()

	for {
		scan.Scan()
		ch <- fmt.Sprintf("%d : Hello ", i)
		i++
	}
}
