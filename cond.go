package main

import "fmt"

func main() {
   /* local variable definition */
   var marks int = 90
   var grade string
   if (marks>=90) {
      grade="A"
   }else if (marks>=80 && marks<90) {
      grade="B"
   }else{
      grade="C"
   }

   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n", grade );      
}