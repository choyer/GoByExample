#WaitGroups NOTES

- sync.WaitGroup is a synchronization primitive used to wait for a collection of goroutines to FINISH executing.
- it maintains an internal counter that tracks how many concurrant tasks are currently running.
- Go 1.25 introduced the `Go(f func())` method to simplify the standard pattern and reduce the need to increment (`wg.Add(1)`) 
