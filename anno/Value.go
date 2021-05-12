package anno

import (
	"github.com/dengpju/higo-config/config"
	"reflect"
)

var Config *config.Configure

type Value struct {
	tag reflect.StructTag
}

func (this *Value) SetTag(tag reflect.StructTag) {
	this.tag = tag
}

func (this *Value) String() string {
	if Config.Exist(this.tag.Get("prefix")) {
		return Config.Get(this.tag.Get("prefix")).(string)
	}
	return ""
}
