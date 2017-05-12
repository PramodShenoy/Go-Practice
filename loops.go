package main

import "fmt"

func main() {
	var a int=10
	var ch string= "*"
	numbers := [6]int{1, 2, 3, 5} 

	//c style loop
	for i := 0; i < a; i++ {
		fmt.Printf("%s\n",ch)
	}

	//while style loop

	for a<20 {
		a++
		fmt.Println(a)
	}

	// python style loop
	/*first impressions:
	i is iterator variable
	x is variable in the range */
	for i,x:= range numbers {
		fmt.Println(i,x)
	}

	
}