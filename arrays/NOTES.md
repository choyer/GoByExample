#arrays NOTES

- **Fixed Length** The length is parts of an arrays type, thus it CANNOT be resized once declared. Hence the reason why `slice` is more commonly used.
- **Homogeneous Type** All elements must be of the same data type
- **Zero-Indexed** Indices start from 0 (zero)
- **Value Types** When an array is assigned to a new variable OR passed to a function, a **copy of the entire array is created**. Thus, changes in the function or new variable DO NOT affect the original array.
- **Contiguous Memory** Elements are stored in adjacent memory locations, making access very efficient.
