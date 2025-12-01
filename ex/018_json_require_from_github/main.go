package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	data := []byte(`{
		"foo": [
			1.23,
			{
				"bar":33,
				"baz":null
			}
		]
	}`)

	fmt.Printf("exists(data.foo) = %v\n", fastjson.Exists(data, "foo"))
	fmt.Printf("exists(data.foo[0]) = %v\n", fastjson.Exists(data, "foo", "0"))
	fmt.Printf("exists(data.foo[1].baz) = %v\n", fastjson.Exists(data, "foo", "1", "baz"))
	fmt.Printf("exists(data.foobar) = %v\n", fastjson.Exists(data, "foobar"))
	fmt.Printf("exists(data.foo.bar) = %v\n", fastjson.Exists(data, "foo", "bar"))

}
