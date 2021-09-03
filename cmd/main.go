package main

import (
	"debug/elf"
	"fmt"
	"github.com/fooree/go-names/internal"
	"github.com/fooree/go-names/internal/foo"
	"github.com/fooree/go-names/internal/foo/ree"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"
)

//go:noinline
func anonymousType() {
	t := reflect.TypeOf(struct {
		Name string
	}{
		Name: "Jack",
	})
	fmt.Printf("name=%s, string=%s, addres=%p\n", t.Name(), t.String(), t)
}

func main() {
	anonymousType()
	ree.Run()
	foo.Y.Foo()
	internal.X.Foo()

	name, _ := filepath.Abs(os.Args[0])
	file, err := elf.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	symbols, err := file.Symbols()
	if err != nil {
		panic(err)
	}

	slice := make([]string, 0, 100)
	for _, symbol := range symbols {
		const module = "github.com/fooree/go-names"
		const name = "main"
		if strings.HasPrefix(symbol.Name, module) || strings.HasPrefix(symbol.Name, name) {
			slice = append(slice, symbol.Name)
		}
	}

	go func() {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		go func() {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}()
	}()

	time.Sleep(time.Second)

	for _, sym := range slice {
		fmt.Println(sym)
	}
}
