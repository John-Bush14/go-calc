package main

import "fmt"


func infixToPostfix(infix string) string {
   var postfix string

   var lastPrescedence int = 0

   var stack []Operator

   for _, char := range infix {
      fmt.Print(string(char), ": char ", postfix, ": postfix ", stack, ": stack \n")

      if char == '(' {
         stack = append(stack, Operator{Operator: OperatorEnum(char)})

         lastPrescedence = -1;

         continue
      }

      if char == ')' {
         for rune(stack[len(stack)-1].Operator) != '(' {
            postfix += string(rune(stack[len(stack)-1].Operator))

            stack = stack[:len(stack)-1]
         }
         
         stack = stack[:len(stack)-1]

         continue
      }

      if isOperator(char) {
         var operator Operator = Operator{
            Operator: OperatorEnum(char),
         }

         for operator.precedence() <= lastPrescedence && len(stack) > 0 {
            var lastItem Operator = stack[len(stack)-1]

            stack = stack[:len(stack)-1]

            postfix += string(rune(lastItem.Operator))
            
            if len(stack) > 0 {lastPrescedence = stack[len(stack)-1].precedence()}
         }
         
         stack = append(stack, operator);
         lastPrescedence = operator.precedence()

         continue
      }

      if char == ' ' {continue}

      postfix += string(rune(char));
   }

   for len(stack) > 0 {
      postfix += string(rune(stack[len(stack)-1].Operator))

      stack = stack[:len(stack)-1]
   }

   return postfix
}


type OperatorEnum rune

const (
   Plus OperatorEnum = '+'
   Minus = '-'
   Asterix = '*'
   Slash = '/'
   Circumflex = '^' 
   Radical = '√'
)

type Operator struct {
   Operator OperatorEnum
}

func (o *Operator) precedence() int {
   switch o.Operator {
      case Plus: return 0
      case Minus: return 0
      case Asterix: return 1
      case Slash: return 1
      case Circumflex: return 2
      case Radical: return 2
      default: return -1
   }
}

func isOperator(char_ rune) bool {
   var char OperatorEnum = OperatorEnum(char_)

   return char == Plus || char == Minus || char == Asterix || char == Slash || char == Circumflex || char == Radical
}
