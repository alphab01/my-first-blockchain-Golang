package main

import time

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
  b2 := block(I: b.i + 1, t: time.Now(), mv: mv1, myh: , prh: , sum: sum1, f: f1, to: to1)
  return b2
}
