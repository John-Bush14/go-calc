package main


func infixToPostfix(infix string) string {
   var postfix string

   var lastPrescedence int = 0

   var stack Stack

   for _, char := range infix {
      if char == ' ' {continue}

      if char == ')' {
         for stack.last() != "(" {
            postfix += stack.pop()
         }
         
         stack.pop()

         continue
      }

      if isOperator(char) {
         postfix += " ";

         for precedence(char) <= lastPrescedence && char != '(' && stack.len() > 0 {
            postfix += stack.pop()
            
            if stack.len() > 0 {
               lastPrescedence = precedence([]rune(stack.last())[0])
            }
         }
         
         stack.append(string(char))
         lastPrescedence = precedence(char)

         continue
      }

      postfix += string(char);
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
   OpenBracket = '('
)

func precedence(operator rune) int {
   switch OperatorEnum(operator) {
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

   return char == Plus || char == Minus || char == Asterix || char == Slash || char == Circumflex || char == Radical || char == OpenBracket
}
