package foo

import (
	"fmt"
	"reflect"
)

type Ree struct {
	Name string
}

//go:noinline
func (r Ree) Foo() {
	anonymousType()
	t := reflect.TypeOf(r)
	fmt.Printf("i am Ree, name=%s, string=%s\n", t.Name(), t.String())
}

//go:noinline
func anonymousType() {
	t := reflect.TypeOf(struct {
		Name string
	}{
		Name: "Jack",
	})
	fmt.Printf("name=%s, string=%s, addres=%p\n", t.Name(), t.String(), t)
}

var Y = Ree{"Rose"}

func init() {
	fmt.Println("Y =",Y)
}
