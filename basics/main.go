package main

import "fmt"

func main() {
	//specificy the range of array and its type
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names :=[4]string{"abc","dce","efg","hij"}

	fmt.Println(ages,len(ages))
	fmt.Println(names,len(names))
	//len() is a built in functiomn to check the lenght of an array

	//slices (use arrays under the hood)
	var scores = []int{100,50,60}
	scores[2]=100;
	scores=append(scores,22)
	fmt.Println(scores);
	// you can't append values to an array, it can only be applied to slices


	//slices ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree:= names[:3]
	fmt.Print(rangeOne,rangeTwo,rangeThree)
 }