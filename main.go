package main

import "fmt"

func greeting() string {
  s := fmt.Sprintf("Hello world!")
  return s
}

func main() {
    fmt.Println(greeting())
}
