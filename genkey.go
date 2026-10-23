package main

import (
  "fmt"
  "crypto/ed25519"
  "crypto/sha256"
  "os"
)

func main() {
  fmt.Println("enter login")
  var s string
  fmt.Scan(&s)
  s2 := sha256.Sum256([]byte(s))
  priv := ed25519.NewKeyFromSeed(s2[:])
  fmt.Printf("%x", priv)
  f, _ := os.Create(s)
  f.Write(priv)
}
