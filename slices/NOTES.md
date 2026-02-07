#slices NOTES

- A thorough write-up of slices can be found in the [Go Blog article - Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- A slice type is an abstraction built on-top of Go's array type. It is really a slice descriptor (header) which is composed of an array pointer (&a), segment length (int), capacity (int)
- A slice is declared like an array, EXCEPT you leave out the element count.
