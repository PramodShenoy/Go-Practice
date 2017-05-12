package main 

import "fmt"

func main() {
	var a int=5
	var b string="*"

	for i := a; i >=0; i-- {
		for j := 0; j<=i ; j++ {
			fmt.Printf(b)
		}
		fmt.Println()
	}
}