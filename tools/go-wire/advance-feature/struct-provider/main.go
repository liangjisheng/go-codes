package main

import "github.com/google/wire"

// Foo ...
type Foo int

// Bar ...
type Bar int

// ProvideFoo ...
func ProvideFoo() Foo {
	return 1
}

// ProvideBar ...
func ProvideBar() Bar {
	return 2
}

// FooBar ...
type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

// Set ...
var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"))

func main() {

}
