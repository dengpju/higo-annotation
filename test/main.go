package main

import (
	"fmt"
	"github.com/dengpju/higo-annotation/anno"
	"reflect"
)

func main() {
	anno.Config.Set("user.age", "85")
	t := reflect.TypeOf(&Test{})
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("prefix"))
	anno.AnnoList.Append(anno.NewValue(field.Tag))
	fmt.Println(anno.AnnoList)
	fmt.Println(anno.Get("*anno.Value"))
}

type Test struct {
	Age *anno.Value `prefix:"user.age"`
}
