# ![dson.png](dson.png)
[![Build Status](https://travis-ci.org/muhammadmuzzammil1998/dson.go.svg?branch=master)](https://travis-ci.org/muhammadmuzzammil1998/dson.go) [![CodeFactor](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/dson.go/badge)](https://www.codefactor.io/repository/github/muhammadmuzzammil1998/dson.go) [![Go Report Card](https://goreportcard.com/badge/github.com/muhammadmuzzammil1998/dson.go)](https://goreportcard.com/report/github.com/muhammadmuzzammil1998/dson.go) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/44f9622f23e748cf9733aeaef7ee6c40)](https://www.codacy.com/app/muhammadmuzzammil1998/dson.go?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=muhammadmuzzammil1998/dson.go&amp;utm_campaign=Badge_Grade) [![Maintainability](https://api.codeclimate.com/v1/badges/a2fa53afcf7a48d55660/maintainability)](https://codeclimate.com/github/muhammadmuzzammil1998/dson.go/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/a2fa53afcf7a48d55660/test_coverage)](https://codeclimate.com/github/muhammadmuzzammil1998/dson.go/test_coverage) [![GitHub license](https://img.shields.io/github/license/muhammadmuzzammil1998/dson.go.svg)](https://github.com/muhammadmuzzammil1998/dson.go/blob/master/LICENSE) [![Twitter](https://img.shields.io/twitter/url/https/github.com/muhammadmuzzammil1998/dson.go.svg?style=social)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fmuhammadmuzzammil1998%2Fdson.go)

**dson.go** provides encoding, decoding, marshaling, unmarshaling, and verification of the DSON (Doge Serialized Object Notation) as defined [here](https://dogeon.xyz/).

## Installing dson.go package

```sh
go get muzzammil.xyz/dson.go
```

## Documentation

- [DSON standard](https://dogeon.xyz/)
- [GoDoc for dson.go](https://godoc.org/muzzammil.xyz/dson.go)

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
