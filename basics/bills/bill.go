package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
// if put parethesis before the name of the fcuntion, we limit the function only to bill objects
func (b *bill) format() string {
	fs := fmt.Sprintf("%-25v \n",b.name)
	var total float64 = 0

	// lists items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%v \n",k+":",v)
		total +=v
	}
    // total
	fs += fmt.Sprintf("%-25v ...%0.2f \n","tip:",b.tip)

	fs += fmt.Sprintf("%-25v ...%0.2f \n","total:",total+b.tip)
	return fs
}
//update the tip

func(b *bill) updateTip(tip float64){
	b.tip =tip
}

func (b *bill) addItem(name string,price float64){
	b.items[name]=price
}


func (b*bill)save(){
	data :=[]byte(b.format())

	err := os.WriteFile("text/"+b.name+".txt",data,0644)

	if err !=nil{
		panic(err)
	}
	fmt.Println("bills was saved to file")
}