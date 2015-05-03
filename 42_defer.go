
package main

import "os"
import "fmt"

func main(){
	defer func(){
		if e := recover(); e != nil {
			fmt.Println("发生错误:",e)
		}
	}()
	
	const file_name = "c:/defer.txt"
	f := createFile(file_name)
	
	defer os.Remove(file_name)
	defer closeFile(f)
	
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}
