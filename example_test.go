package tcg_test

import (
	"errors"
	"fmt"

	"github.com/ysmood/tcg"
)

type data struct{}

func (d data) bar() string {
	tcg.Throw(errors.New("err"))
	return ""
}

func foo() data {
	tcg.Throw(errors.New("err"))
	return data{}
}

func Example_catch() {
	defer tcg.Catch(func(err error) {
		fmt.Println(err)
	})

	s := foo().bar()
	fmt.Println(s)

	// Output: err
}

func Example_guard() {
	var d data
	err := tcg.Guard(func() {
		d = foo()
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)

	// Output: err
}
