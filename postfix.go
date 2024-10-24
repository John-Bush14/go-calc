package main


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
      default: return 9999
   }
}

func isOperator(char_ rune) bool {
   var char OperatorEnum = OperatorEnum(char_)

   return char == Plus || char == Minus || char == Asterix || char == Slash || char == Circumflex || char == Radical
}
