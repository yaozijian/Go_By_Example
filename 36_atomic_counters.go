
package main

import "fmt"
import "time"
import "sync/atomic"

func main(){

	var ops uint64
	
	for i := 0; i < 50; i++ {
		go func(){
			for{
				time.Sleep(time.Millisecond)
				atomic.AddUint64(&ops,1)
			}
		}()
	}
	
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)
}
