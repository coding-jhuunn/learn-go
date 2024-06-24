package main

// if the package is called main, tells go compiler that our code should be compiled in a executable code
import "fmt"

// formatting and printing messages

//if the func name main, it fires automatically
func main(){
	fmt.Println(("hello "))

	//strings
	var nameOne string = "Mario";

	var nameTwo = "string"

	var nameThree string

	fmt.Println(nameOne,nameTwo,nameThree);

	//shorcut for declaring var in variables
	//you cant use ":=" outise of a function
	nameOur := "yoshi";
	fmt.Println(nameOur);

	// ints 

	var ageOne int = 23;
	var ageTwo = 30;

	ageThree := ageOne+ageTwo
	println(ageThree)

	//bits & memory
	//specifci the range of a number 
	var numOme int8 = 12
	var numTwo uint16= 129
	// uint specify we cant have a negative number

	var scoreOne float32 =25.98
	var scoreTwo float64 = 1243532452346234.343

	fmt.Println(numOme,nameTwo,scoreOne,scoreTwo,nameTwo)
}