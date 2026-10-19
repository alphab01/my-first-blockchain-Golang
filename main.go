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
}
