
package main

import "fmt"

func main(){

	jobs := make(chan int,5)
	done := make(chan bool)
	
	/*
	go func(){
		for{
			j,more := <- jobs
			if more {
				fmt.Println("received job",j)
			}else{
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	
	for j := 1; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	*/
	
	// 用下面的代码就不需要close了
	// 默认的收发操作是阻塞的,使用select则变成非阻塞的
	for j := 1; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	
	go func(){
		for{select{
		case j := <- jobs: fmt.Println("received job",j)
		default:
			fmt.Println("received all jobs")
			done <- true
			return
		}}
	}()	
	
	//close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
