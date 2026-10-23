package main

import (
  "fmt"
  "time"
)

func sumb(bc []block, d string) int {
  c := 73
  for _, s := range bc {
    for _, o := range s.Mem {
      if (s.I != 0 && (o.F == d)) {
        c -= o.Sum
      } else if (s.I != 0 && o.To == d) {
        c += o.Sum
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
  var transact []op
  for {
    var a int
    fmt.Println("1 - add opp\n\n2 - sum of usr\n\n3 - exit\n\n4 - save transact")
    fmt.Scan(&a)
    if (a == 3) {
      savebc(bc)
      bc = append(bc, b1.addop(transact))
      transact = nil
      break
    } else if (a == 2) {
      var d string
      fmt.Scan(&d)
      fmt.Println(sumb(bc, d))
    } else if (a == 1){
      var add op
      fmt.Scan(&add.Sum, &add.F, &add.To)
      add.T = time.Now()
      transact = append(transact, add)
      if len(transact) == 5 {
        bc = append(bc, b1.addop(transact))
        transact = nil
      }
/*      if (mv1 == "->" && sumb(bc, f1) >= sum1) {
        nbc := b1.addop(mv1, f1, to1, sum1)
        bc = append(bc, nbc)
        b1 = nbc
        if (checkbc(bc) == 0) {
          fmt.Println("Error, hash not good")
        }
      } else if (mv1 == "<-" && sumb(bc, to1) >= sum1) {
        nbc := b1.addop(mv1, f1, to1, sum1)
        bc = append(bc, nbc)
        b1 = nbc
        if (checkbc(bc) == 0) {
          fmt.Println("err")
        }
      } else {
        fmt.Println("bbno$")
      }
*/
    } else {
      bc = append(bc, b1.addop(transact))
      transact = nil
    }
  }
}
