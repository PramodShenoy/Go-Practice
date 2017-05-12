package main 

import "fmt"

func main() {
	var a int=10
	var b string="*"

	for i := a; i >=0; i-- {
		for j := 0; j<=i ; j++ {
			fmt.Printf(b)
		}
		fmt.Println()
	}


	for i := 0; i <=a; i++ {
		for j := 0; j<=i ; j++ {
			if i==3 {
				continue;//or break
			}
			fmt.Printf(b)
		}
		fmt.Println()
	}


}
