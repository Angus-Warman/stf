package stf

import (
	"sync"
)

// Parallelise distributes tasks across a pool of workers. If any worker's
// process call returns an error, that worker stops, all other workers are
// signalled to stop before picking up their next task, and the first error
// is returned. Tasks already in-flight are allowed to finish naturally.
func Parallelise[T any](workers int, tasks []T, process func(T) error) error {
	if len(tasks) == 0 {
		return nil
	}
	if workers <= 0 {
		workers = 1
	}

	taskCh := make(chan T, len(tasks))
	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	var (
		wg       sync.WaitGroup
		once     sync.Once
		firstErr error
		stopped  = make(chan struct{}) // closed when the first error occurs
	)

	stop := func(err error) {
		once.Do(func() {
			firstErr = err
			close(stopped)
		})
	}

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for task := range taskCh {
				// Check whether another worker has already failed.
				select {
				case <-stopped:
					return
				default:
				}

				if err := process(task); err != nil {
					stop(err)
					return
				}
			}
		}()
	}

	wg.Wait()
	return firstErr
}
