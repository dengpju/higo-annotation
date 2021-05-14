package anno

import (
	"reflect"
)

type Value struct {
	Tag reflect.StructTag
}

func NewValue() *Value {
	return &Value{}
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.Tag = tag
}

func (this *Value) String() string {
	if Config.Exist(this.Tag.Get("prefix")) {
		return Config.Get(this.Tag.Get("prefix")).(string)
	}
	return ""
}
