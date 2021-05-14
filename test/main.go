package main

import (
	"fmt"
	"github.com/dengpju/higo-annotation/anno"
	"reflect"
)

func main() {
	t := reflect.TypeOf(&Test{})
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("prefix"))
	fmt.Println(anno.Get("*anno.Value"))
}

type Test struct {
	Age *anno.Value `prefix:"user.age"`
}
