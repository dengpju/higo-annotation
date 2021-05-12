package main

import (
	"fmt"
	"github.com/dengpju/higo-annotation/annotation"
	"reflect"
)

func main() {
	t := reflect.TypeOf(&Test{})
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("prefix"))
}

type Test struct {
	Age *anno.Value `prefix:"user.age"`
}
