package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Println("hello ",n)
}

func cycleNames(n []string,f func(string)){
	for _, v:=range n{
		f(v)
	}
}
func circleArea(r float64)float64{
	return math.Pi*r *r
}
func main() {
sayGreeting("hello")
	cycleNames([]string{"aasd","asdasd","asdasdsa"},sayGreeting)
} 