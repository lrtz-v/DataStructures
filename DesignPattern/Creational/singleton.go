package creationaldesignpattern

import "sync"

// 单例模式

// Singleton type
type Singleton map[string]string

var (
	once sync.Once

	instance Singleton
)

// NewSingleton init instance
func NewSingleton() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})

	return instance
}
