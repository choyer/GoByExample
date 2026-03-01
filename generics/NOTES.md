#generics NOTES

- Go generics, introduced in Go 1.18, allow functions and types to work with a variety of data types while maintaining type safety. This feature improves code reusability and helps prevent type-related errors at compile time, eliminating the need for repetitive code or unsafe type assertions with interface{}
- [Further tutorial](https://go.dev/doc/tutorial/generics)
- `S ~[]E` defines `s` as any type whose *underlying type* is `[]E`. The `~` (tilde) means "this type, OR any named type built on top of it." eg. not just plain `[]int`, but also custom type like `type MyList []int`.
