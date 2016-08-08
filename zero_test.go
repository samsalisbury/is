package is

import (
	"reflect"
	"testing"
)

type ZeroTestStruct struct {
	Nested            *ZeroTestStruct
	String            string
	notExportedString string
	notExportedNested *ZeroTestStruct
	Embedded          struct {
		Int         int
		notExported string
	}
}

type ZeroTestUnexportedStruct struct {
	nested *ZeroTestStruct
	str    string
}

var zeroValues = []interface{}{
	"", 0, nil, 0.0, 0x000, 00,
	ZeroTestStruct{},
	ZeroTestStruct{Embedded: struct {
		Int         int
		notExported string
	}{0, ""}},
}

func TestIsZero_isZero(t *testing.T) {
	for _, v := range zeroValues {
		if !Zero(v) {
			t.Errorf("%+v is non-zero; want zero", v)
		}
	}
}

var nonZeroValues = []interface{}{
	"a", 1,
	ZeroTestStruct{Nested: &ZeroTestStruct{}},
	ZeroTestStruct{String: "a"},
	ZeroTestStruct{notExportedString: "a"},
	ZeroTestStruct{notExportedNested: &ZeroTestStruct{}},
	ZeroTestStruct{Embedded: struct {
		Int         int
		notExported string
	}{1, ""}},
	ZeroTestStruct{Embedded: struct {
		Int         int
		notExported string
	}{0, "a"}},
}

func TestIsZero_notZero(t *testing.T) {
	for _, v := range nonZeroValues {
		if Zero(v) {
			t.Errorf("%+v is zero; want non-zero", v)
		}
	}
}

var zeroForTypeValues = []interface{}{nil}

func TestIsZeroForType_zero(t *testing.T) {
	typ := reflect.TypeOf(zeroForTypeValues).Elem()
	for _, v := range zeroForTypeValues {
		if !ZeroForType(v, typ) {
			t.Errorf("%+v is non-zero; want zero", v)
		}
	}
}

var nonZeroForTypeValues = []interface{}{
	"", 0, 0.0, 0x000, 00,
	ZeroTestStruct{},
	ZeroTestStruct{Embedded: struct {
		Int         int
		notExported string
	}{0, ""}},
}

func TestIsZeroForType_nonZero(t *testing.T) {
	typ := reflect.TypeOf(nonZeroForTypeValues).Elem()
	for _, v := range nonZeroForTypeValues {
		if ZeroForType(v, typ) {
			t.Errorf("%+v is zero; want non-zero", v)
		}
	}
}
