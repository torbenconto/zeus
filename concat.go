package zeus

import "sync"

func NonConcurrentConcat(strings ...string) string {
	out := ""

	for _, s := range strings {
		out += s
	}
	return out
}

// Concatenates multiple strings together using goroutines. Using this as opposed to a normal concatenation is less efficient with small concatenations but works very well with large concatenations.
func Concat(strings ...string) string {
	var (
		out   string
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	ch := make(chan string, len(strings))

	for _, s := range strings {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			ch <- s
		}(s)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		mutex.Lock()
		out += s
		mutex.Unlock()
	}

	return out
}
