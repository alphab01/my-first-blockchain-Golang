package main

import (
  "time"
  "crypto/sha256"
  "fmt"
  "encoding/json"
  "os"
)

type op struct {
  T time.Time
  Sum int
  F string
  To string
}

type block struct {
  I uint
  T time.Time
  Myh []byte
  Prh []byte
  Mem []op
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
  t := time.Now()
  s := fmt.Sprintf("%d %s %x %v", 0, t, "init", []byte("all usrs have 73"), nil)
  h.Write([]byte(s))
  b := block{I: 0, T: t, Myh: h.Sum(nil), Prh: []byte("all usrs have 73"), Mem: nil}
  return b
}

func (b block) addop(add []op) block {
  h := sha256.New()
  t := time.Now()
  s := fmt.Sprintf("%d %s %x %v", b.I + 1, t, b.Myh, add)
  h.Write([]byte(s))
  b2 := block{I: b.I + 1, T: t, Myh: h.Sum(nil), Prh: b.Myh, Mem: add}
  return b2
}

func (b block) saveb(s string) {
  dat, err := json.Marshal(b)
  if (err != nil) {
    return
  }
  f, err := os.Create(s)
  if (err != nil) {
    return
  }
  f.WriteString(string(dat))
  f.Close()
}

func savebc(bc []block) {
  for i := range len(bc) {
    if (bc[i].I == 0) {
      bc[i].saveb("Genesis")
    } else {
      s := fmt.Sprintf("b%d", bc[i].I)
      bc[i].saveb(s)
    }
  }
}
