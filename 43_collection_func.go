
package main

import "strings"
import "fmt"

func Index(vs []string,t string) int {
	for i,v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string,t string) bool {
	return Index(vs,t) >= 0
}

func Any(vs []string,f func(string) bool) bool {
	for _,x := range vs {
		if f(x) {
			return true
		}
	}
	return false
}

func All(vs []string,f func(string) bool) bool {
	for _,x := range vs {
		if !f(x) {
			return false
		}
	}
	return true
}

func Filter(vs []string,f func(string) bool) []string{
	vsm := make([]string,0)
	for _,x := range vs {
		if f(x) {
			vsm = append(vsm,x)
		}
	}
	return vsm
}

func Map(vs []string,f func(string) string) []string {
	vsm := make([]string,len(vs))
	for i,v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	
	var strs = []string{"peach","apple","pear","plum"}
	
	fmt.Println(Index(strs,"pear"))
	fmt.Println(Include(strs,"grape"))
	
	fmt.Println(Any(strs,func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	
	fmt.Println(All(strs,func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	
	fmt.Println(Filter(strs,func(v string) bool {
		return strings.Contains(v,"e")
	}))
	
	fmt.Println(Map(strs,strings.ToUpper))
}
