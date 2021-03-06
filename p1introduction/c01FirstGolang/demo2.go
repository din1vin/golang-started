package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	var name = flag.String("name", "everyone", "greeting to")
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
}
