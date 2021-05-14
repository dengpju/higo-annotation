package anno

import (
	"github.com/dengpju/higo-config/config"
	"sync"
)

var (
	Config   *config.Configure
	AnnoList *Annotations
	once     sync.Once
)

func init() {
	once.Do(func() {
		Config = config.New()
		AnnoList = NewAnnotations()
	})
	AnnoList.Append(new(Value))
}

func Get(key string) Annotation {
	return AnnoList.Get(key)
}
