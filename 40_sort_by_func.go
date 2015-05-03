
package main

import "fmt"
import "sort"

type SortByLength []string

func (p_array SortByLength) Len() int{
	return len(p_array)
}

func (p_array SortByLength) Less(x,y int) bool{
	return len(p_array[x]) < len(p_array[y])
}

func (p_array SortByLength) Swap(x,y int){
	p_array[x],p_array[y] = p_array[y],p_array[x]
}

func main(){
	fruits := []string{"peach","banana","kiwi"}
	sort.Sort(SortByLength(fruits))
	fmt.Println(fruits)
}
