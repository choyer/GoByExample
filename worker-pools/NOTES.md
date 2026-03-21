#worker-pools NOTES

- *Worker Pools* are a concurrency pattern used to efficiently manage a **fixed number of goroutines** (workers) that process a queue of tasks (jobs).
- Instead of spawning a new goroutine for every task. a set number of worker goroutines are started once.
- A channel (typically buffered) acts as a shared queue where all the tasks are sent.
- Workers pull jobs from the job channel and send results to an optional results channel.
- Worker pools are ideal for situations where you have large or unbounded number of tasks and limited system resources.
