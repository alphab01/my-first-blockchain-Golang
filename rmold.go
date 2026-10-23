package main

import (
  "os"
  "fmt"
)

func main() {
  for i := 1; ; i++{
    s := fmt.Sprintf("b%d", i)
    err := os.Remove(s)
    if (err != nil) {
      break
    }
  }
}
