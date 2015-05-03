
package main

import "fmt"

func main(){

	messages := make(chan string)
	signals := make(chan bool)
	
	select{
	case msg := <- messages: fmt.Println("received message",msg)
	default: fmt.Println("no message received")
	}
	
	// 注意缓冲与非缓冲的差别
	// 如果上面创建messages的语句是: messages := make(chan string,1)
	// 则创建的是缓冲的程道，则这里的发送操作可以完成
	msg := "hi"
	select{
	case messages <- msg: fmt.Println("sent message",msg)
	default: fmt.Println("no message sent")
	}
	
	select{
	case msg := <- messages: fmt.Println("received message",msg)
	case sig := <- signals: fmt.Println("received signal",sig)
	default: fmt.Println("no activity")
	}
}
