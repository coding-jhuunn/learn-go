package main

import "fmt"

func main() {
	//formating strings
	numOne := 35
	nameOne := "me"
	fmt.Printf("my age is %v and my name is %v \n", numOne, nameOne)
	fmt.Printf("my age is %q and my name is %q \n", numOne, nameOne)
	fmt.Printf("age is of type %T",numOne)
	//this how you specificy fomrat

	//%v is for concatenating variables with strings
	//%q is for
	//%T is for gettung the type of variable
	//%f is for bolean , can specficy the range of boolean of decimal


	//Sprintf(save formatted strings)
	// use for saving strings
	var str = fmt.Sprintf("my age is %v and my name is %v \n", numOne, nameOne)
	fmt.Println(str)
}