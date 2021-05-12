package anno

import "sync"

var (
	AnnoList Annotations
	once           sync.Once
)

func init() {
	once.Do(func() {
		AnnoList = make([]Annotation, 0)
	})
	AnnoList.Append(new(Value))
}
