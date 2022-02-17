package main

import (
	"dingliang/p1introduction/c04ProgramInitOrder/pkg1"
	"fmt"
)

func init() {
	fmt.Println("init method calling")
}
func main() {
	fmt.Println("running")
	pkg1.Main()
}
