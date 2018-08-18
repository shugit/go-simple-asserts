package go_simple_asserts

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

var pathType = "file_name"
var pathType2 = "absolute_path"

func AssertEqual(expected, actual interface{}, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	comparison := "be"
	switch exp := expected.(type) {
	case int:
		if exp != actual.(int) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int32:
		if exp != actual.(int32) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int64:
		if exp != actual.(int64) {
			printFail(t, loc, actual, comparison, exp)
		}
	case uint64:
		if exp != actual.(uint64) {
			printFail(t, loc, actual, comparison, exp)
		}
	case string:
		if exp != actual.(string) {
			printFail(t, loc, actual, comparison, exp)
		}
	case bool:
		if exp != actual.(bool) {
			printFail(t, loc, actual, comparison, exp)
		}
	case []byte:
		actualBytes := actual.([]byte)
		if bytes.Compare(exp, actualBytes) != 0 {
			printFail(t, loc, actual, comparison, exp)
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func AssertNotEqual(expected, actual interface{}, t *testing.T, str string) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	comparison := fmt.Sprintf("not equal to")
	switch exp := expected.(type) {
	case int:
		if exp == actual.(int) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int32:
		if exp == actual.(int32) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int64:
		if exp == actual.(int64) {
			printFail(t, loc, actual, comparison, exp)
		}
	case uint64:
		if exp == actual.(uint64) {
			printFail(t, loc, actual, comparison, exp)
		}
	case string:
		if exp == actual.(string) {
			printFail(t, loc, actual, comparison, exp)
		}
	case bool:
		if exp == actual.(bool) {
			printFail(t, loc, actual, comparison, exp)
		}
	case []byte:
		actualBytes := actual.([]byte)
		if bytes.Compare(exp, actualBytes) == 0 {
			printFail(t, loc, actual, comparison, exp)
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func AssertGreaterThan(expected, actual interface{}, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	comparison := "greater than or equal to"
	switch exp := expected.(type) {
	case int:
		if exp > actual.(int) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int32:
		if exp > actual.(int32) {
			printFail(t, loc, actual, comparison, exp)
		}
	case int64:
		if exp > actual.(int64) {
			printFail(t, loc, actual, comparison, exp)
		}
	case string:
		if exp > actual.(string) {
			printFail(t, loc, actual, comparison, exp)
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func printFail(t *testing.T, loc string, actual interface{}, comparison string, expected interface{}) {
	fmt.Printf("%v:[FAIL]%v should %v %v\n", loc, actual, comparison, expected)
	t.Fail()
}

func AssertNull(actual interface{}, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	if reflect.ValueOf(actual).Kind() == reflect.Ptr {
		var pointerValue = fmt.Sprintf("%v", actual)
		if pointerValue != "<nil>" {
			printFail(t, loc, actual, "be", "nil")
		}
	} else {
		if actual != nil {
			printFail(t, loc, actual, "be", "nil")
		}
	}
}

func AssertNotNull(actual interface{}, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	if reflect.ValueOf(actual).Kind() == reflect.Ptr {
		if fmt.Sprintf("%v", actual) == "<nil>" {
			printFail(t, loc, actual, "not be", "nil")
		}
	} else {
		if actual == nil {
			printFail(t, loc, actual, "not be", "nil")
		}
	}
}

func AssertTrue(actual bool, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	if actual != true {
		printFail(t, loc, actual, "be", true)
	}
}
func AssertFalse(actual bool, t *testing.T) {
	_, fn, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf("%v:%v", fn, line)
	if actual != false {
		printFail(t, loc, actual, "be", false)
	}
}
