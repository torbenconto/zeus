package zeus

import "sync"

func NonConcurrentConcat(strings ...string) string {
	out := ""

	for _, s := range strings {
		out += s
	}
	return out
}

func Concat(strings ...string) string {
	var (
		out   string
		mutex sync.Mutex
	)

	for _, s := range strings {
		s := s
		go func() {
			mutex.Lock()
			out += s
			mutex.Unlock()
		}()
	}

	for range strings {
		mutex.Lock()
		mutex.Unlock()
	}

	return out
}
