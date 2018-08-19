package main

import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
)

func main() {
  in := bufio.NewReader(os.Stdin)
  fmt.Println("First operand: ")
  s1, _ := in.ReadString('\n')
  s1 = strings.TrimRight(s1, "\n")
  op1, _:= strconv.ParseFloat(s1, 64)
  fmt.Println("Second operand: ")
  s2, _ := in.ReadString('\n')
  s2 = strings.TrimRight(s2, "\n")
  op2, _:= strconv.ParseFloat(s2, 64)
  res := Adder(op1, op2)
  fmt.Println("Result is ", res)
}

func Adder(op1, op2 float64) float64 {
  return op1 + op2
}
