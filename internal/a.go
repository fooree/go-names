package internal

import (
	"fmt"
	"reflect"
)

type Foo interface {
	Foo()
}

type Int int

var X Int

//go:noinline
func (i *Int) Foo() {
	t := reflect.TypeOf(i)
	go func() {
		fmt.Printf("i am Int, name=%s, string=%s\n", t.Name(), t.String())
	}()
}

func init() {
	X = Int(0x123)
}

func init() {
	fmt.Println("X =", X)
}
