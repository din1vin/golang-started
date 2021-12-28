package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "greeting to")
}

func main() {
	flag.Parse()
	hello(name)
}
