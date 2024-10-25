package main

import "fmt"


func infixToPostfix(infix string) string {
   var postfix string

   var lastPrescedence int = 0

   var stack Stack

   for _, char := range infix {
      fmt.Print(string(char), ": char ", postfix, ": postfix ", stack, ": stack \n")

      if char == '(' {
         stack.append(string(char))

         lastPrescedence = -1;

         continue
      }

      if char == ')' {
         for stack.last() != "(" {
            postfix += stack.pop()
         }
         
         stack.pop()

         continue
      }

      if isOperator(char) {
         postfix += " ";

         for precedence(OperatorEnum(char)) <= lastPrescedence && stack.len() > 0 {
            postfix += stack.pop()
            
            if stack.len() > 0 {
               lastPrescedence = precedence(lastStackOperator(stack))
            }
         }
         
         stack.append(string(char))
         lastPrescedence = precedence(OperatorEnum(char))

         continue
      }

      if char == ' ' {continue}

      postfix += string(rune(char));
   }

   for stack.len() > 0 {
      postfix += stack.pop()
   }

   return postfix
}

func lastStackOperator(stack Stack) OperatorEnum {
   return OperatorEnum([]rune(stack.last())[0])
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

func precedence(operator OperatorEnum) int {
   switch operator {
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
