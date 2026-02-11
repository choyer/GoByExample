#range-over-built-in-types NOTES

`range` can be used on **arrays, slices, strings, maps or channels**. New in Go 1.23 **Range over Function Types**

**Slices and Arrays**: Returns `index, value`
**Maps**: Returns `key, value`
**Strings**: Returns `index, rune` (rune = Unicode code point)
**Channels**: Returns elements sent to the channel (until closed)
**Function Types**: see [https://go.dev/blog/range-functions](https://go.dev/blog/range-functions)


## Considerations 

- **Value Copying**: `range` returns a *copy* of the element, not a reference, so modifying the value directly DOES NOT update the original slice/array.
- **Map Ordering**: Iteration order for maps is **random**.
