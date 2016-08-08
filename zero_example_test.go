package is

import (
	"fmt"
	"reflect"
)

func ExampleZero_true() {
	a := ""
	b := 0
	c := ((*int)(nil))
	fmt.Println("a:", Zero(a), "b:", Zero(b), "c:", Zero(c))
	// output:
	// a: true b: true c: true
}

func ExampleZero_false() {
	a := "a"
	b := 1
	c := &b
	fmt.Println("a:", Zero(a), "b:", Zero(b), "c:", Zero(c))
	// output:
	// a: false b: false c: false
}

func ExampleZeroForType_false() {
	// i is the type interface{}
	i := reflect.TypeOf((*interface{})(nil)).Elem()
	fmt.Println(
		"a:", ZeroForType("a", i),
		"b:", ZeroForType(1, i),
		"c:", ZeroForType(i, i),
		"d:", ZeroForType(1, nil),
	)
	// output:
	// a: false b: false c: false d: false
}

func ExampleZeroForType_true() {
	// i is the type interface{}
	i := reflect.TypeOf((*interface{})(nil)).Elem()
	fmt.Println(
		"a:", ZeroForType("", reflect.TypeOf("a")),
		"b:", ZeroForType(0, reflect.TypeOf(1)),
		"c:", ZeroForType(nil, reflect.TypeOf(&i)),
		"d:", ZeroForType(nil, reflect.TypeOf(i)),
	)
	// output:
	// a: true b: true c: true d: true
}
