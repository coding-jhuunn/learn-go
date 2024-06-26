package main

import "fmt"

func main() {
	mybill := newBill("My bill")
	mybill.addItem("rice",2.5)
	mybill.addItem("drinks",2.1)
	mybill.addItem("main dish",8.4)
	mybill.tip = 10.22
	fmt.Println(mybill.format())

}