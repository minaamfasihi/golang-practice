package main
import "fmt"

func main() {
  if 7 % 2 == 0 {
    fmt.Println("true")
  } else {
    fmt.Println("fasle")
  }

  if 8 % 4 == 0 {
    fmt.Println("true")
  }

  if num := 9; num < 0 {
    fmt.Println(num, " is neg")
  } else if num > 0 {
    fmt.Println(num, " has 1 digit")
  } else {
    fmt.Println(num, " has multiple digits")
  }
}
