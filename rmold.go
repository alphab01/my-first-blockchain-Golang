package main

import (
  "os"
  "fmt"
)

func main() {
  os.Remove("Genesis")
  for i := 1; ; i++{
    s := fmt.Sprintf("b%d", i)
    err := os.Remove(s)
    if (err != nil) {
      break
    }
  }
}
