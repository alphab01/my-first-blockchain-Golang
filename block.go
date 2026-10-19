package main

import time

type block struct {
  I uint
  t time.Time
  mv string
  myh []byte
  prh []byte
  sum int
}
