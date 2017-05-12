package main

import "fmt"

func main() {
   /* local variable definition */
   var a int 
   var b int 
   var c int
   var ret int

   fmt.Scan(&a)
   fmt.Scan(&b)

   /* calling a function to get max value */
   ret = max(a, b)

   fmt.Printf( "Max value is : %d\n", ret )

   a,b,c= swap(10,20,30)
   fmt.Printf("%d\t%d\t%d",a,b,c)


}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}


func swap(x, y ,z int) (int, int,int ) {
   return z,y,x
}
