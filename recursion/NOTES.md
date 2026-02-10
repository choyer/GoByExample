#recursion NOTES

- Each recursive call adds a new frame to the call stack. In Go, this can be inefficient, as the language DOES NOT implement [tail-call optimization](https://deepsource.com/blog/common-mistakes-in-go). Thus large recursion depths will exhaust memory.
- ANY problem that can be solved with recursion can also be solved with iteration, which can sometimes be more efficient in Go due to memory management.
- When using recursion, the base case is crucial. Without a proper base case, a recursive function will encounter an infinite-loop, consuming excessive memory and ultimately leading to a stack overflow.
