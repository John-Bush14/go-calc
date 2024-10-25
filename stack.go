package main


type Stack struct {
   Array []string
}

func (s *Stack) pop() string {
   var lastItem string = s.Array[len(s.Array)-1]

   s.Array = s.Array[:len(s.Array)-1]

   return lastItem
}

func (s *Stack) append(item string) {
   s.Array = append(s.Array, item)
}

func (s *Stack) get(index int) string {
   return s.Array[index]
}

func (s *Stack) last() string {
   return s.get(len(s.Array)-1)
}

func (s *Stack) len() int {
   return len(s.Array)
}

