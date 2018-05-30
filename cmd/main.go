package main

import (
	"flag"
	"fmt"

	"github.com/sunreaver/heartrate"
)

func main() {
	var age = flag.Int("age", 31, "年龄")
	var static = flag.Int("static", 56, "静态心率")
	flag.Parse()

	fmt.Printf("年龄:%d 静态心率:%d\n", *age, *static)
	out, e := heartrate.KarvonenFormula(*static, *age)
	if e != nil {
		fmt.Println("error: ", e.Error())
	} else {
		fmt.Println(out)
	}
}
