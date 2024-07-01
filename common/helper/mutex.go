package helper

import "sync"

var (
	mu    sync.Mutex
	locks = make(map[string]*sync.Mutex)
)

func GetOrCreateMutex(key string) *sync.Mutex {
	mu.Lock()
	defer mu.Unlock()
	
	if _, exists := locks[key]; !exists {
		locks[key] = &sync.Mutex{}
	}
	return locks[key]
}