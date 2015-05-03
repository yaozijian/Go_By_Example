
package main

import "fmt"

func main(){
	// 注意: make(chan string) 不等于 make(chan string,1)
	// 前者是无缓冲的程道，后者是缓冲的程道
	messages := make(chan string,2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
