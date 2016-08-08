# is: reflective assertions for golang

Package `is` provides some assertions on reflect.Values and reflect.Types.

Go's reflect package can sometimes leave one asking basic questions, like
"how do I know if a value is a zero value", or "is this field exported?".
Although there are well-documented ways to answer these kinds of questions,
discoverability can be hard.

I hope this package comes in handy for quick and dirty assertions, but would
encourage you to use this code as inspiration to writing your own reflect
code more accurately and efficiently.

## is.Zero? is.ZeroForType?

`is.Zero` checks if a value is equal to the zero value of its underlying type.

`is.ZeroForType` checks if a value is equal to the zero value of the provided type.

These two functions are subtly different.

## TODO

- `is.ExportedField(reflect.StructField)`
- `is.ExportedType(reflect.Type)`
