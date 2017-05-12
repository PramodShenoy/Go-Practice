package main

import (
   "fmt"
   "math")

//fucntion closure
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
	  return i  
   }
}

func main(){
   /* declare a function variable 
	similar to lambdas in python*/
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* use the function */
   var a float64
   fmt.Scan(&a)
   fmt.Println(getSquareRoot(a))


   number:=getSequence()
   fmt.Println(number())
   fmt.Println(number())
   fmt.Println(number())

}