package main

import (
  "fmt"
  //"crypto/ed25519"
)

func main() {
  fmt.Println("enter login")
  var s string
  fmt.Scan(&s)
  c := 0
  for i := range s {
    c += int(s[i])
  }
  fmt.Println(c)
}
