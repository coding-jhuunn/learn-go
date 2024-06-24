package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting,"hello"))
	//fnc that return boolean if contains the string
	fmt.Println(strings.ReplaceAll(greeting,"hello","hi"))
	//fnc that subtitute that target strings, does not replace the origina value
	fmt.Println(strings.ToUpper(greeting))
	//fnc upper case
	fmt.Println(strings.Index(greeting,"ll"))
	//find the index of targeted string
	fmt.Println(strings.Split(greeting,""))

	ages :=[]int{45,12,53,8,3,85,35}
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(sort.SearchInts(ages,30))

}