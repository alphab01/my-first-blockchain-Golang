package main

import (
  "time"
  "crypto/sha256"
)

type block struct {
  i uint
  t time.Time
  mv string
  myh []byte
  prh []byte
  sum int
  f string
  to string
}

func (b block) addop(mv1 string, f1 string, to1 string, sum1 int) block {
  h := sha256.New()
  s := fmt.Sprintf("%d %s %x %d %s %s", b.i + 1, mv1, b.myh, sum1, f1, to1)
  b2 := block{I: b.i + 1, t: time.Now(), mv: mv1, myh: s, prh: b.myh, sum: sum1, f: f1, to: to1}
  return b2
}
