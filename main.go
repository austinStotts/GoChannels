package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Println("running... :)")

// 	numbers := make(chan int, 10)

// 	for i := 1; i <= 5; i++ {
// 		numbers <- i
// 	}

// 	close(numbers)

// 	for num := range numbers {
// 		fmt.Println(num)
// 	}

// }

func main() {

	numbers := make(chan int, 10)
	stop := make(chan string, 1)

	go func() {
		for i := 1; i <= 10; i++ {
			numbers <- i
			time.Sleep(500 * time.Millisecond)
		}
		stop <- "stop"
		close(numbers)
		close(stop)
	}()

	for {
		select {
		case num := <-numbers:
			fmt.Println(num)
		case <-stop:
			fmt.Println("STOP")
			return
		}
	}
}
