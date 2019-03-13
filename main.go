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

// *** *** *** *** *** *** *** ***

// func main() {

// 	numbers := make(chan int, 10)
// 	stop := make(chan string, 1)

// 	go func() {
// 		for i := 1; i <= 10; i++ {
// 			numbers <- i
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 		stop <- "stop"
// 		close(numbers)
// 		close(stop)
// 	}()

// 	for {
// 		select {
// 		case num := <-numbers:
// 			fmt.Println(num)
// 		case <-stop:
// 			fmt.Println("STOP")
// 			return
// 		}
// 	}
// }

func main() {
	odd := make(chan int, 10)
	even := make(chan int, 10)
	stop := make(chan int, 1)
	results := make(chan int, 10)
	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
			time.Sleep(500 * time.Millisecond)
		}
		stop <- -1
	}()

	for {
		select {
		case e := <-even:
			results <- (e + <-odd)
			fmt.Println(<-results)
		case <-stop:
			fmt.Println("stop")
			return
		}
	}
}
