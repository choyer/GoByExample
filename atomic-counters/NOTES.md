#atomic-counters NOTES

- Introduced in Go 1.19
- Go allows you to implement highly efficient, lock-free thread-safe counters using the `sync/atomic` package.
- Atomic counters allow multiple goroutines to increment, decrement, and read integer values concurrently without encountering data races or requiring heavy `sync.Mutex` locking mechanisms.
- See [Ralph Caraveo's Go 1.19 Atomic Wrappers article](https://medium.com/@deckarep/the-go-1-19-atomic-wrappers-and-why-to-use-them-ae14c1177ad8) -> https://medium.com/@deckarep/the-go-1-19-atomic-wrappers-and-why-to-use-them-ae14c1177ad8
