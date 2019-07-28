package main

import (
  "fmt"
  "math"
)

const s string = "constant"

var foo int = 123 // this is declared at the package level. at this level, variables can only be declared like this.
// cant use abc := 99 syntax here

// go allows us to group our declarations in a block like below.

var (
  actorName string = "Elisabeth Sladen"
  companion string = "Sarah Jane Smith"
  doctorNumber int = 3
  season int = 11
)

var (
  counter int = 0
)

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
  k := 23.0 // drawback of this style is that you cant control the type of variable too finely. Like you cant ever declare float32
          // it will always be float64. 
  fmt.Println(k)
  // below %v stands for value, %T for type and Printf() allows us to format and lets us use %v and %T
  fmt.Printf("%v, %T", k, k) // 23.0, float64
  fmt.Printf("%v, %T", i, i) // 42, int

}
