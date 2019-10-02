# Value and Reference Types in GOLANG

## Value Types

The following types are copy-on-call; unless indirection and address operators are used, original value may not be mutated from within a function or method.

* int
* float
* string
* bool
* struct

## Reference Types

The following types are essentially a struct containing a pointer to a backing store. Thus, the underlying data can be mutated from within a function (without indirection or address operators)

* slices
* maps
* channels
* pointers
* functions