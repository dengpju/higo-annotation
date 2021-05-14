package anno

import (
	"fmt"
	"reflect"
)

type Value struct {
	Tag reflect.StructTag
}

func NewValue(tag reflect.StructTag) *Value {
	return &Value{Tag: tag}
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.Tag = tag
}

func (this *Value) String() string {
	fmt.Println(this.Tag)
	fmt.Println(Config)
	fmt.Println(this.Tag.Get("prefix"))
	fmt.Println(Config.Exist(this.Tag.Get("prefix")))
	if Config.Exist(this.Tag.Get("prefix")) {
		return Config.Get(this.Tag.Get("prefix")).(string)
	}
	return ""
}
