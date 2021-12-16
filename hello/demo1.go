// go脚本传参
package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "greetings to")
}

func main() {
	flag.Parse()
	fmt.Printf("hello, %s\n", name)
}
