package util

// ExecContext containing information about the current execution in the whole exec chain
type ExecContext struct {
	Index int
}

// Exec execute fn function n times with a concurrency level of c
func Exec(n int, c int, fn func(ExecContext)) {
	semaphore := make(chan bool, c)
	for i := 0; i < n; i++ {
		semaphore <- true
		go func(index int) {
			defer func() { <-semaphore }()
			fn(ExecContext{Index: index})
		}(i)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}
}
