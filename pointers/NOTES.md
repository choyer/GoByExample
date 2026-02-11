#pointers NOTES

**Pointers** in Go are *variables that store the memory address of another variable, enabling direct access and modification of the original data*.

## Concepts

| Operator | Name | Purpose | Example |
|----------|------|---------|---------|
| & | **Address-of Operator** | Returns the memory address of a variable | `p := &v` stores the address of `v` in `p` |
| * | **Dereferencing Operator** | Accesses the value stored at the pointer's memory address | `fmt.Println(*p)` prints the value at the address `p` holds |
| *T | **Type Declaration** | Declares a variable as a pointer to a value of type `T` | `var p *int` declares `p` as a pointer to an integer |

## Uses and Best Practices

- **Mutation**: a primary reason to use a pointer is when you want a function or method to modify the **original variable**. When a variable is passed without a pointer, Go uses *pass-by-value*, thus a copy.
- **Performance**: For large structs, passing a pointer can be more efficient than copying the entire object, saving memory and CPU time.
- **Methods with Pointer Receivers**: Methods defined with a pointer receiver ( `func (p *T) MethodName()` ) can modify the underlying struct instance. Commonly found for methods that need to change the object's state.
- Built-in types such as *slices, maps, and channels* are already reference types which hold pointers to the underlying data structure, so you don't typically need to use explicit pointers to them.
