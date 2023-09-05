package zeus

import "sync"

// NonConcurrentConcat concatenates multiple strings together without concurrency.
// This function is usually less efficient for small concatenations but works well for large concatenations.
//
// Parameters:
//   - strings: A variadic list of strings to concatenate.
//
// Returns:
//   - The concatenated string.
func NonConcurrentConcat(strings ...string) string {
	out := ""

	for _, s := range strings {
		out += s
	}

	return out
}

// Concat concatenates multiple strings together using goroutines and a mutex for synchronization.
// It is more efficient for large concatenations compared to NonConcurrentConcat.
//
// Parameters:
//   - strings: A variadic list of strings to concatenate.
//
// Returns:
//   - The concatenated string.
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
