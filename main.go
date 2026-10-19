package main

import (
  "fmt"
)

func sumb(bc []block, d string) int {
  c := 73
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
  return c
}

func main() {
  bc := []block{}
  var b1 block
  b1 = begin()
  bc = append(bc, b1)
  for {
    var a int
    fmt.Println("1 - add opp\n\n2 - sum of usr\n\n3 - exit")
    fmt.Scan(&a)
    if (a == 3) {
      break
    } else if (a == 2) {
      var d string
      fmt.Scan(&d)
      fmt.Println(sumb(bc, d))
    } else {
      var mv1, f1, to1 string
      var sum1 int
      fmt.Scan(&mv1, &f1, &to1, &sum1)
      if (mv1 == "->" && sumb(bc, f1) >= sum1) {
        nbc := b1.addop(mv1, f1, to1, sum1)
        bc = append(bc, nbc)
        b1 = nbc
      } else if (mv1 == "<-" && sumb(bc, to1) >= sum1) {
        nbc := b1.addop(mv1, f1, to1, sum1)
        bc = append(bc, nbc)
        b1 = nbc
      } else {
        fmt.Println("bbno$")
      }
    }
  }
}
