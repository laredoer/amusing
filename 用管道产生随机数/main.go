package main

import "fmt"

// 随机数可控 0,1
func main() {
	for i := range random(100) {
		fmt.Println(i)
	}
}

func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			case c <- 2:
			}
		}
	}()
	return c
}
