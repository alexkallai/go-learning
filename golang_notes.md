# Data types:

### Integers:
```go
int8
int16
int32
int64

uint8 (byte)
uint16
uint32
uint64

int (32 or 64 bits)
uint (32 or 64 bits)

rune (synonym of int32, represents Unicode code points)

uintptr
```

### Floats:
```go
float32
float64
```

### Complex numbers:
```go
complex64
complex128
```

### Booleans:
```go
bool (true, false)
```

Operations:
AND – &&
OR  – ||
Negation – !

### Strings:
```go
string
```

# Composite types:
## Non-Reference Types
### Arrays




# Variable declaration:
Declared variables will contain zero value for that type!
```go
var variablename int (variablename is now == 0)
variablename = 45

or

var variablename int = 45

or

// implicit variable declaration
variablename := 45

or

var length, width float32
length, width = 32, 64
```

# Const declaration:
```go
const variablename int = 45
```

## Working with different typed variables:
```go
var length float64 = 1.2
var width int = 2
length = float64(width)
fmt.Println(length)
```

## If condition:
```go
if grade == 100 {
    fmt.Println("Perfect!")
} else if grade >= 60 {
    fmt.Println("You pass!")
} else {
    fmt.Println("You fail!")
}
```

## Loops:

Simple for loop:
```go
for x := 4; x <= 6; x++ {
    fmt.Println("x is now", x)
}
```

There is no while loop, but a while loop is possible with the for loop, if only the condition expression is supplied
Simple while loop:
```go
x := 1
for x <= 6 {
    fmt.Println("x is now", x)
    x++
}
```

Breaking out of loop
```go
x := 1
for true {
    fmt.Println("x is now", x)
    if x >= 2 {
        fmt.Println("The condition is met, breaking out of the loop...")
        break
    }

    x++
}
```

## Working with strings:
Printf: print with formatting
Sprintf : return the formatted string

Formatting verb example:
```go
floatvar := 5.55555555555555555
fmt.Printf("The value of the variable is %0.2f\n", floatvar)
// Output: The value of the variable is 5.55
```
Formatting verbs:
| Verb | Output |
| ---- | ------ |
| %f  | Floating-point number |
| %d  | Decimal integer |
| %s  | String |
| %t  | Bool |
| %v  | Any value (chooses appropriate format based on supplied value's type) |
| %#v | Any value, formatted as it would appear in Go program code |
| %T  | Type of the supplied value (int, string...) |
| %%  | Literal percent sign |

## %5.3f
% - start of the formatting verb
5 - minimum width of the __entire__ number
3 - width after the decimal point
f - formatting verb type

__%.3f__ is also correct

# Functions:
Simple function with a single return type
```go
func double(number float64) float64 {
    return number * 2
}
```

Multiple return types
```go
func manyReturns() (int, bool, string) {
    return 1, true, "hello"
}
```

Multiple return types with error
```go
func manyReturns() (int, error) {
    if false {
        return 0, fmt.Errorf("Something went wrong")
    } else {
        return 1, nil
    }
}
```

## Pointers TODO:
```go

```

## Arrays:
Initialized arrays will have 0 values for the variable type where they are not specified
```go
var notes[7]string
notes[0] = "do"
notes[1] = "re"
notes[2] = "mi"
fmt.Println(notes[0])
fmt.Println(notes[1])
```
do
re

### Array literals
```go
[3]int{3, 5, 44}
```

```go
text := [3]string{
    "First",
    "Second",
    "Third", // the comma is mandatory!
}
```

### Using arrays in loops:
```go
for index, value := range myArray {
    // Loop block
}

// Unused variables can be left out
for _, value := range myArray {
    // Loop block
}
```

## Slices:
Slices are a view into arrays (changes will also change slices)
New elements can be added to slices
Declaring a slice doesn't automatically create a slice
The built-in __make__ function can create it
You pass the type of the slice and the length of the slice

Unlike arrays, the slice variable itself also has a zero value: it's **nil**
-> a slice variable that has no slice assigned to, will have a value of nil
len() function will return 0 if passed a nil slice
append() will add the specified slice to the 0 length slice, and thus create one
```go
var notes []string
notes = make([]string, 10)
```
After the slice is created, the same syntax can be used for assigning and retrieving elements
```go
notes[0] = "do"
notes[1] = "re"
```
Just as with other variables, a slice doesn't have to be declared in separate steps:
```go
primes := make([]int, 5)
```

```go
// __Slice__ declaration -> empty pair of brackets
var mySlice []string
// Array declaration
var mySlice [3]string
```

### Slice literals:
A slice can be created with the elements already present (no length specified)
```go
sliceLiteral := []int{1, 2, 3}

// Multi-line slice literal:
sliceLiteralMulti := []int{
    1, 
    2, 
    3,
}
```
### Slice operator:
```go
// 1: index of array where slice shoud start
// 3: index of array where slice shoud stop !!before!!
mySlice := myArray[1:3]
```

### Append to slices:
If there is no room in the array to add elements, all its elements will be copied to a new, larger array, and the slice will be updated to refer to this new array.
```go
slice := []string{"a", "b"}
slice = append(slice, "c")
slice = append(slice, "d", "e")
```
---
Keeping both slices can lead to unpredictable behavior, because a newly created array can be created, but the original is preserved too
When calling **append**, it's conventional to just assign the return value back to the same slice variable one passed to append

---



## CodeSection:
```go

```





