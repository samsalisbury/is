# is – Reflective Assertions for Golang

[![CircleCI](https://circleci.com/gh/samsalisbury/is.svg?style=svg)](https://circleci.com/gh/samsalisbury/is)
[![codecov](https://codecov.io/gh/samsalisbury/is/branch/master/graph/badge.svg)](https://codecov.io/gh/samsalisbury/is)
[![Go Report Card](https://goreportcard.com/badge/github.com/samsalisbury/is)](https://goreportcard.com/report/github.com/samsalisbury/is)
[![GoDoc](https://godoc.org/github.com/samsalisbury/is?status.svg)](https://godoc.org/github.com/samsalisbury/is)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Package `is` provides some assertions on `reflect.Values` and `reflect.Types`.

Go's reflect package can sometimes leave one asking basic questions, like

> Is this a zero value?

or

> Is this field exported?

Although there are well-documented ways to answer these kinds of questions,
discoverability can be hard.

I hope this package comes in handy for quick and dirty assertions, but would
encourage you to use this code as inspiration to writing your own reflect
code more accurately and efficiently.

## Is this a zero value?

`is.Zero` checks if a value is equal to the zero value of its underlying type.

`is.ZeroForType` checks if a value is equal to the zero value of the provided type.

These two functions are subtly different.

## TODO

- `is.ExportedField(reflect.StructField)`
- `is.ExportedType(reflect.Type)`
