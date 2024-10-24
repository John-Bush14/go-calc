package main

import "fmt" 

func main() {
   fmt.Print(calculatePostfix(infixToPostfix("(3+5^2-2*16√2)/(4+6-2^3)*(9-3*25√2)+(7^2-3^3+2*49√2)")))
}
