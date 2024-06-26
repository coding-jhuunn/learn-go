package main

import "fmt"

func updateName(n *string)  {
	// now we are passing a pointer
	//if use * in front of argument that means accpeting in pointer
	*n = "wedege"
	
}

func main() {

	name := "tifa"
	
	

	//fmt.Println("memory address of name is: ",&name)
	//to get the memory address of a variable
	
	m:=&name
	fmt.Println("memory address",m)
	// use * to get the value of the pointer
	fmt.Println("value at memory address: ",*m)
	fmt.Println(name)
	
	updateName(m)
	fmt.Println(name)
}