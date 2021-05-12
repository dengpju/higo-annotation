package anno

import (
	"reflect"
)

type Annotation interface {
	SetTag(tag reflect.StructTag)
}

type Annotations []Annotation

// 是否是注解
func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnoList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

func (this Annotations) Append(anno Annotation) Annotations {
	this = append(this, anno)
	return this
}
