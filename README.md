# ![dson.png](dson.png)

**dson.go** provides encoding, decoding, marshaling, unmarshaling, and verification of the DSON (Doge Serialized Object Notation) as defined [here](https://dogeon.xyz/).

## Installing dson.go package

```sh
go get muzzammil.xyz/dson.go
```

## Examples

### Common imports

```go
import (
    "fmt"

    "muzzammil.xyz/dson.go"
)
```

### Encoding JSON into DSON

```go
func main() {
    d := dson.Encode(`{"foo":"bar"}`)
    fmt.Println(d) // such "foo" is "bar" wow
}
```

### Decoding DSON into JSON

```go
func main() {
    j := dson.Decode(`such "foo" is "bar" wow`)
    fmt.Println(j) // {"foo":"bar"}
}
```

### Validating DSON

```go
func main() {
    if dson.Valid(`such "foo" is "bar" wow`) {
        fmt.Println("Valid DSON")
    } else {
        fmt.Println("Invalid DSON")
    }
}
```

### Marshalling DSON

```go
func main() {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
    }
    RedGroup := ColorGroup{
        ID:     1,
        Name:   "Reds",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }
    r, err := dson.Marshal(RedGroup)
    if err == nil && dson.Valid(r) {
        fmt.Println(r) // such "ID" is 1! "Name" is "Reds". "Colors" is so "Crimson" and "Red" and "Ruby" also "Maroon" many wow
    }
}
```

### Unmarshalling DSON

```go
func main() {
    d := `so such "Name" is "Platypus" and "Order" is "Monotremata" wow and such "Name" is "Quoll" and "Order" is "Dasyuromorphia" wow many`
    if !dson.Valid(d) {
        fmt.Println("DSON is not valid")
        return
    }
    type Animal struct {
        Name  string
        Order string
    }
    var animals []Animal
    err := dson.Unmarshal(d, &animals)
    if err == nil {
        fmt.Printf("%+v", animals) // [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
    }
}
```
