
package main

import(
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct{
	key int
	resp chan int
}

type writeOp struct{
	key int
	val int
	resp chan bool
}

func main(){

	var state = make(map[int]int)
	var ops int64
	
	read_chnl := make(chan *readOp)
	write_chnl := make(chan *writeOp)
	
	go func(){
		for{select{
		// 读取操作: 将请求的数据写入到应答程道中
		case read_op := <- read_chnl:
			read_op.resp <- state[read_op.key]
		// 写入操作: 写入数据,在应答程道中写入true表示写入完成
		case write_op := <- write_chnl:
			state[write_op.key] = write_op.val
			write_op.resp <- true
		}}
	}()
	
	for r := 0; r < 100; r++ {
		go func(){
			for{
				read_op := &readOp{
					key : rand.Intn(5),
					resp: make(chan int),
				}
				// 将读取请求写入到程道中
				read_chnl <- read_op
				// 等待读取完成: 读取到的数据应该是写入到了应答程道的
				<- read_op.resp
				atomic.AddInt64(&ops,1)
			}
		}()
	}
	
	for w := 0; w < 10; w++ {
		go func(){
			for{
				write_op := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp:make(chan bool),					
				}
				// 将写入请求写入到程道中
				write_chnl <- write_op
				// 等待写入完成：完成后应该可以从程道收取到一个布尔值
				<- write_op.resp
				atomic.AddInt64(&ops,1)
			}
		}()	
	}
	
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadInt64(&ops)
	
	fmt.Println("ops:",opsFinal)
}
