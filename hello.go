package main

import "fmt"

func main() {
  var array [1000]int

  for i := 0; i < 1000; i++ {
  array[i] = i
  }

  fmt.Println(array)
}
