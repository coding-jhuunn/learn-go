package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])


	//looop
	//key then value variable
	for k, v:=range menu{
		fmt.Println(k,"-",v)
	}

	// ints as key type

	phonebook :=map[int] string{
		232123:"Mario",
		123123:"asd",
		454454:"asda",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[123123])

	//update an item in a map

	phonebook[123123]="bowser"
	fmt.Println(phonebook[123123])


}