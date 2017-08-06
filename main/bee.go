package main

import "fmt"

func main() {

	var a circle
	b := circle{}
	c := new (circle)
	d := &circle{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

type circle struct {
	x int
	z int
	y int
}