package main

import "fmt"

func updateName(n string) string {
	n = "wedege"
	return n
	//it should have return to update the value of name var
}

func updateMap (n map[string]float64){
	n["coffee"]=6.23
}


func main() {
	name := "tifa"

	// updateName(name)
	// will not change the value of the name
	name = updateName(name)
	fmt.Println(name)

	// key first then the value if you use map
	menu := map[string]float64{
		"pie": 5.95,
		"ice cream":3.99,

	}

	updateMap(menu)
	fmt.Println(menu)

}