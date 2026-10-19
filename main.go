package main

import (
  "fmt"
)

type usr struct {
  B int
  Name string
}

func main() {
  U := []usr{}
  for i := 0; i < 5; i++ {
    s := fmt.Sprintf("test%d", i)
    test := usr{B: 100, Name: s}
    U = append(U, test)
  }
  bc := []block{}
  var b1 block
  b1.init()
  bc = append(bc, b1)
  for {
    var a int
    fmt.Println("1 - add opp\n\n2 - sum of usr\n\n3 - exit")
  }
}
