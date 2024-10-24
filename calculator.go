package main

import (
	"math"
	"strconv"
)


func calculatePostfix(postfix string) float64 {
   var stack []string

   stack = append(stack, "")

   for _, char := range postfix {
      if isOperator(char) {
         x, _ := strconv.ParseFloat(stack[len(stack)-2], 64)
         y, _ := strconv.ParseFloat(stack[len(stack)-1], 64)

         print(x, " ", y, "", string(char), "\n")

         stack = append(stack[:len(stack)-2], strconv.FormatFloat(calculate(OperatorEnum(char), x, y), 'f', 2, 64))

         continue
      }

      if char == ' ' {
         stack = append(stack, "")

         continue
      }

      stack[len(stack)-1] += string(char)
   }

   ans, _ := strconv.ParseFloat(stack[0], 64)

   return ans
}

func calculate(operator OperatorEnum, x float64, y float64) float64 {
   switch rune(operator) {
      case '+': 
         return x + y
      case '-':
         return x - y
      case '*':
         return x * y
      case '/':
         return x / y
      case '^':
         return math.Pow(x, y)
      case '√':
         if y == 2 {return math.Sqrt(x)}
   }

   return 0
}
