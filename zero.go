// Package is provides some assertions on reflect.Values and reflect.Types.
//
// Go's reflect package can sometimes leave one asking basic questions, like
// "how do I know if a value is a zero value", or "is this field exported?".
// Although there are well-documented ways to answer these kinds of questions,
// discoverability can be hard.
//
// I hope this package comes in handy for quick and dirty assertions, but would
// encourage you to use this code as inspiration to writing your own reflect
// code more accurately and efficiently.
package is

import "reflect"

// Zero returns true when the underlying value of v is either nil or zero for
// its underlying type.
func Zero(v interface{}) bool {
	if v == nil {
		return true
	}
	return ZeroForType(v, reflect.TypeOf(v))
}

// ZeroForType returns true when v  matches the zero value of typ.
// If typ is nil, always returns false.
func ZeroForType(v interface{}, typ reflect.Type) bool {
	if typ == nil {
		return false
	}
	if v == nil {
		return true
	}
	z := reflect.Zero(typ).Interface()
	return reflect.DeepEqual(z, v)
}
