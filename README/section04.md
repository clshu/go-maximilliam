## Section 4: Understanding Pointers
### Avoid Unnecessary Value Copies
- #### By default, Go creates a copy when passing values to functions
- #### For very large & complex values, this may take up too much memory space unnecessarily
- #### With pointers, only one va;ue is stored in memory( and the address is passed around)

### Directly Mutate Values
- #### Pass a pointer(address) instead of a value to a function
- #### The function can then directly edit the underlying value - no return value is required
- #### Cab lead to less code (but also to less understandle code or unexoected behaviors)