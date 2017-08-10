package main

import (
	"fmt"
	"strings"
)

func main()  {
	var tag = `jf:"#.JPUSH.AppKey"`
	fmt.Println(tag[:strings.Index(tag, ".")])
	fmt.Println(tag[strings.Index(tag, ".")+1:])

}
