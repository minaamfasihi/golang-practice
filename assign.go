package main

import "fmt"

func main() {
  a := 23
  fmt.Println("a:", a)
  const b = 12
  fmt.Println("b:", b)
  var i int
  for i = 1; i <= 5; i++ {
    fmt.Println(i)
  }

  var x = 10
  if x > 10 {
    fmt.Println("X is greater")
  } else if x < 10 {
    fmt.Println("X is smaller")
  } else {
    fmt.Println("X is equal")
  }
}
