
package main

import "fmt"
import "time"

func main(){
	
	requests := make(chan int,5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	// 每隔200毫秒取下一个请求
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<- limiter
		fmt.Println("每隔200毫秒:",req,time.Now())
	}
	
	// 先连续输入三个当前时间
	burstyLimiter := make(chan time.Time,3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	
	// 然后每隔200毫秒输入当前时间
	go func(){
		for t := range time.Tick(time.Millisecond * 200){
			burstyLimiter <- t
		}
	}()
	
	// 突发的请求
	burstyRequests := make(chan int,5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	
	// 先连续取三个请求,然后每隔200毫秒取一个请求
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("连续3个,然后每隔200毫秒:",req,time.Now())
	}
}
