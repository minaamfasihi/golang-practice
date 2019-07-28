package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  // var a = "initial"
  var a string = "initial"
  fmt.Println(a)
  var b, c int = 1, 2
  fmt.Println(b)
  fmt.Println(c)
  var e int
  fmt.Println(e)
  f := "short"
  fmt.Println(f)
  g := 12
  fmt.Println(g)
  fmt.Println(s)
  const n = 50000
  const d = 3e20 / n
  fmt.Println(d)
  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))

  var i int
  i = 42
  fmt.Println(i)
  var j int = 42
  fmt.Println(j)
  k := 23 // drawback of this style is that you cant control the type of variable too finely. Like you cant ever declare float32
          // it will always be float64. 
  fmt.Println(k)

}
