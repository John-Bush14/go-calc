package main

import (
	"math"
	"strconv"
)


func calculatePostfix(postfix string) float64 {
   var stack Stack

   stack.append("")

   for _, char := range postfix {
      if isOperator(char) {
         y, _ := strconv.ParseFloat(stack.pop(), 64)
         x, _ := strconv.ParseFloat(stack.pop(), 64)

         stack.append(strconv.FormatFloat(calculate(OperatorEnum(char), x, y), 'f', 2, 64))

         continue
      }

      if char == ' ' {
         stack.append("")

         continue
      }

      stack.Array[len(stack.Array)-1] += string(char)
   }

   ans, _ := strconv.ParseFloat(stack.pop(), 64)

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
