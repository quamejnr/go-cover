package main

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Div(num int, denom int) int {
  if denom == 0 {
    panic("invalid division error")
  }
	return num / denom
}
