package anno

import (
	"reflect"
)

type Annotation interface {
	SetTag(tag reflect.StructTag)
}

type Annotations struct {
	Sort  []Annotation
	Value map[string]Annotation
}

func NewAnnotations() *Annotations {
	return &Annotations{Sort: make([]Annotation, 0), Value: make(map[string]Annotation, 0)}
}

// 是否是注解
func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnoList.Value {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

func Apply()  {
	
}

func (this *Annotations) Append(anno Annotation) *Annotations {
	this.Value[reflect.TypeOf(anno).String()] = anno
	this.Sort = append(this.Sort, anno)
	return this
}

func (this *Annotations) Get(key string) Annotation {
	return this.Value[key]
}

