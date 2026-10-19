package main

import (
  "fmt"
)

func main() {
  bc := []block{}
  var b1 block
  b1.init()
  bc = append(bc, b1)
  for {
    var a int
    fmt.Println("1 - add opp\n\n2 - sum of usr\n\n3 - exit")
    fmt.Scan(&a)
    if (a == 3) {
      break
    } else if (a == 2) {
      var d string
      fmt.Scan(&c)
      c := 100
      for _, s := range bc {
        if (s.I != 0 && (s.F == d)) {
          if (s.Mv == "->") {
            c -= s.Sum
          } else {
            c += s.Sum
          }
        } else if (s.I != 0 && s.To == d) {
          if (s.Mv == "->") {
            c += s.Sum
          } else {
            c -= s.Sum
          }
        }
      }
      fmt.Println(c)
    }
  }
}
