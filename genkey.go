package main

import (
  "fmt"
  "crypto/ed25519"
  "crypto/sha256"
  "os"
)

func main() {
  fmt.Println("enter login and password")
  var s, s1 string
  fmt.Scan(&s, &s1)
  s2 := sha256.Sum256([]byte(s + s1))
  priv := ed25519.NewKeyFromSeed(s2[:])
  fmt.Printf("%x", priv)
  f, _ := os.Create(s)
  f.Write(priv)
}
