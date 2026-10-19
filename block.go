package main

import (
  "time"
  "crypto/sha256"
  "fmt"
)

type block struct {
  I uint
  T time.Time
  Mv string
  Myh []byte
  Prh []byte
  Sum int
  F string
  To string
}

func (b block) addop(mv1 string, f1 string, to1 string, sum1 int) block {
  h := sha256.New()
  s := fmt.Sprintf("%d %s %x %d %s %s", b.i + 1, mv1, b.myh, sum1, f1, to1)
  b2 := block{I: b.i + 1, T: time.Now(), Mv: mv1, Myh: s, Prh: b.myh, Sum: sum1, F: f1, To: to1}
  return b2
}
