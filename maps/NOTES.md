#maps NOTES

- a `map` is an **unordered** collection of key-value pairs providing for fast data retrieval, updates and deletions base on unique keys.
- maps are implemented as hash tables.
- They a **reference typed**: when passed to a function or assigned to a new variable, it passes a reference to the same underlying data structure.
- They have a **dynamic size**
- Keys must be of a comparable type (e.g. `int`, `string`, `float64`, `pointers`, `structs`, `arrays`). NO slices, functions or maps.
- Maps are **NOT safe for concurrent use (e.g goroutine)**. For concurrent scenarios use `sync.Map` OR protect map access with `sync.Mutex`
