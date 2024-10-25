package main

import (
	"fmt" 
)

func main() {
   for true {
      fmt.Print("What would you like to calculate?\n")

      
      var calculation string; 

      fmt.Scan(&calculation)


      fmt.Print("Result: ", calculatePostfix(infixToPostfix(calculation)))

      fmt.Scanln()

      fmt.Print("\033[H\033[2J")
   }
}
