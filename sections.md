# Simple programs that use common packages

## Variable, values, and types

- Packages
- Short declaration operator
- The var keyword
- Exploring types
- Zero value
- The fmt package
- Creating your own types
- Conversion, not casting

### Short declaration operator

- Limit the scope of a variable
- Use the short operator has often as possible

```go
x = 45
x:= 45
var x int
```

### Creating you own types

- You can create you own data type using the ***type***

keyword.
Example:  

```go
var a int
type hotbar int
var b hotdog

funct main(){
  a = 1
  b = 2
  a = b // The values are not equivalent
  a = int(b) // This works
}
```

### Conversion, not casting
