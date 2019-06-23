package main

import "fmt"

func visit (list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func print(v int) {
	fmt.Println(v)
}
func main(){
	f := print
	visit([]int{1,2,3,4,5}, f)
}
