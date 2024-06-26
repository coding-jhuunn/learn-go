package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	opt,_ := getInput("Option: (a - add item , s -save bill, t- add tip)",reader)
	switch opt{
	case "a":
		name,_:=getInput("Item name",reader)
		price,_:=getInput("Item Price",reader)

		p,err :=strconv.ParseFloat(price,64)
		if err !=nil{
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name,p)
		fmt.Println("Item added -",name,p)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you save the file")
	case "t":
		tip,_:=getInput("Tip",reader)
		
		t,err :=strconv.ParseFloat(tip,64)
		if err !=nil{
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added -",t)
		promptOptions(b)
		fmt.Println(tip)
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