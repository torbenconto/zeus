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
