# Laws of Reflection

## 1. Reflection goes from interface value to reflection object.
```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    t := reflect.TypeOf(x)
    fmt.Println("type:", t)


    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    // Note that .Kind() can't differentiate a Float64 from MyFloat64, even tho .Type() can
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())
    }
```

## 2. Reflection goes from reflection object to interface value.
Like physical reflection, reflection in Go Generates its own inverse.

## 3. To modify a reflection object, the value must be settable.
Settability
- Is a bit like addressability, but stricter. 
- Its the property that a reflection object can modify the actual storage that was used to create the reflection object. 
- Settability is determined by whether the reflection object holds the original item.

Key take-away:  Reflection `Values` need the address of something in order to modify what they represent.