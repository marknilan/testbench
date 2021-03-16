package main

import "fmt"
import "github.com/marknilan/testbench"
import "github.com/marknilan/testbench/clientlib"

func main() {
  fmt.Println(testbench.Config())
  fmt.Println(clientlib.Hello())
}
