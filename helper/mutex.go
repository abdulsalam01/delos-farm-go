package helper

import "sync"

func MutexLockUnLock() {
	mu := sync.Mutex{}

	mu.Lock()
	defer mu.Unlock()
}
