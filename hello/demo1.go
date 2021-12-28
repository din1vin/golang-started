// go脚本传参
package hello

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name","everyone","greeting to")
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
}
