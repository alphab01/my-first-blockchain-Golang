package main

import (
  "time"
  "crypto/sha256"
  "fmt"
  "encoding/json"
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

func eqb(p []byte, m []byte) bool {
  a := len(p)
  for i := range a {
    if (p[i] != m[i]) {
      return false
    }
  }
  return true
}

func checkbc(bc []block) int {
  b1 := bc[0]
  for _, s := range bc {
    if (s.I > 0) {
      if (!(eqb(s.Prh, b1.Myh))) {
        return 0
      }
      b1 = s
    }
  }
  return 1
}

func begin() block {
  h := sha256.New()
  s := fmt.Sprintf("%d %s %s %x %d %s %s", 0, time.Now(), "init", []byte("all usrs have 73"), 0, "x", "x")
  h.Write([]byte(s))
  b := block{I: 0, T: time.Now(), Mv: "init", Myh: h.Sum(nil), Prh: []byte("all usrs have 73"), Sum: 0, F: "x", To: "x"}
  return b
}

func (b block) addop(mv1 string, f1 string, to1 string, sum1 int) block {
  h := sha256.New()
  s := fmt.Sprintf("%d %s %s %x %d %s %s", b.I + 1, time.Now(), mv1, b.Myh, sum1, f1, to1)
  h.Write([]byte(s))
  b2 := block{I: b.I + 1, T: time.Now(), Mv: mv1, Myh: h.Sum(nil), Prh: b.Myh, Sum: sum1, F: f1, To: to1}
  return b2
}

func (b block) savebc() {
  dat, err := json.Marshal(b)
  if (err != nil) {
    return
  }
  fmt.Println(string(dat))
}
