package main

import "fmt"

func swap(p *int ,q *int) {
	var t int
	t = *p
	*p = *q
	*q = t
}
func main() {
	var x int
	var y int
	x = 1
	y = 2
	swap(&x,&y)
	fmt.Println("x=",x,"y=",y)
	
}
