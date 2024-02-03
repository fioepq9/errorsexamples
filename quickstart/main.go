package main

import (
	"fmt"

	"github.com/fioepq9/errors"
)

func foo(method string) error {
	switch method {
	case "new":
		return errors.New("foo")
	case "wrap":
		return errors.Wrap(foo("new"), "this is wrapeed foo")
	default:
		panic("invalid method")
	}
}

func bar(method string) error {
	switch method {
	case "new":
		return foo("new")
	case "wrap":
		return errors.Wrap(foo("wrap"), "this is wrapeed bar")
	default:
		panic("invalid method")
	}
}

func baz(method string) error {
	switch method {
	case "new":
		return bar("new")
	case "wrap":
		return errors.Wrap(bar("wrap"), "this is wrapeed baz")
	default:
		panic("invalid method")
	}
}

func main() {
	errors.C.Style = errors.StyleStack

	fmt.Println(baz("new"))
	fmt.Println(baz("wrap"))
}

/*
foo
  main.go:12 (0x658aba) main.foo()
  main.go:23 (0x658b7a) main.bar()
  main.go:34 (0x658c3a) main.baz()
  main.go:46 (0x658d24) main.main()
  D:/apps/scoop/apps/go/current/src/runtime/proc.go:267 (0x605411) runtime.main()

foo
  main.go:12 (0x658aba) main.foo()
this is wrapeed foo
  main.go:14 (0x658af5) main.foo()
this is wrapeed bar
  main.go:25 (0x658bb5) main.bar()
this is wrapeed baz
  main.go:36 (0x658c75) main.baz()
  main.go:47 (0x658d6e) main.main()
  D:/apps/scoop/apps/go/current/src/runtime/proc.go:267 (0x605411) runtime.main()

*/
