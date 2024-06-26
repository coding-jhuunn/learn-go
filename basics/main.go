package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error){	
	fmt.Println(prompt)
	input,err :=r.ReadString('\n')
	return strings.TrimSpace(input),err
}



func  createBill() bill{
	reader:=bufio.NewReader(os.Stdin)


	name,_ := getInput("Create a new bill name:",reader)

	b:= newBill(name)
	fmt.Println("Created the bill - ",b.name)
	return b
}

func promptOptions(b bill){
	reader:=bufio.NewReader(os.Stdin)
	opt,_ := getInput("Option: (a - add item , s -save bill, t- add tip",reader)
	switch opt{
	case "a":
		name,_:=getInput("Item name",reader)
		price,_:=getInput("Item Price",reader)
		fmt.Println(name,price)
	case "s":
		tip,_:=getInput("Tip",reader)

		fmt.Println(tip)
	case "t":
		fmt.Println("you choose t")
	default:
		fmt.Println("not valid")
		promptOptions(b)
}
}

func main() {

	mybill:=createBill()
	promptOptions(mybill)

	fmt.Println(mybill)
	

}