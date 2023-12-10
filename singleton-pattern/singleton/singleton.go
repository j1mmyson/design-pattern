package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
	count int
}

var singletonInstance *singleton

// GetInstance returns the singleton instance
func GetInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singletonInstance == nil {
			fmt.Println("Creating single instance now.")
			singletonInstance = new(singleton)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singletonInstance
}

// AddCount adds 1 to the count
func (s *singleton) AddCount() {
	s.count++
}

// GetCount returns the count
func (s *singleton) GetCount() int {
	return s.count
}
