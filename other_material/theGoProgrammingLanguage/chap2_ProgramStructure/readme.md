# Program Structure

## Go has 25 key words:

break
case
chan
const
continue dedault
defer
else
fallthrough
for
func
go
goto
if 
import 
interface
map
package
range
return
select 
struct
switch 
type
var

## Predeclared names

Constants: true false iota nil
Types: 	   int int8 int16 ...
Functions: make len cap new append copy ...


## 2.2 Declarations 

Four kind of declarations: var, const, type, func

## 2.3 Variables
- Sets of variable can be initialized by calling a function that return multiple values
var f, err = os.Open(name) // opens a file and returns an error

- It is optional to intialize a set a variables 
var x, y, int
var x, y = true, 2.3 // bool, float64 

### 2.3.1 Short Variable Declaration

:= // is an declaration
= // is an assignment 

### 2.3.2 Pointers 
- pointer evaluate to true if and only if the addresses are the same. 

- pointer evaluate to fasle is the variable is not set. 

- Go has no pointer arithmetric 

### 2.3.3 The new Function




