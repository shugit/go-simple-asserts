package go_simple_asserts

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"strings"
)

var PrintFileName = "file_name"
var PrintAbsolutePath = "absolute_path"

var FileType = PrintFileName

func getFileName(path string) string {
	var splits = strings.Split(path, "/")
	return splits[len(splits)-1]
}

func getLocation() *string {
	_, fn, line, _ := runtime.Caller(2)
	if FileType == PrintFileName {
		fn = getFileName(fn)
	}
	loc := fmt.Sprintf("%v:%v", fn, line)
	return &loc
}

func AssertEqual(actual, expected interface{}, t *testing.T) {
	loc := getLocation()
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
			printFail(t, loc, string(actualBytes), comparison, string(exp))
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func AssertNotEqual(actual, expected interface{}, t *testing.T, str string) {
	loc := getLocation()
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
			printFail(t, loc, string(actualBytes), comparison, string(exp))
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func AssertGreaterThan(actual, expected interface{}, t *testing.T) {
	loc := getLocation()
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

func printFail(t *testing.T, loc *string, actual interface{}, comparison string, expected interface{}) {
	fmt.Printf("%v:[FAIL]%v should %v %v\n", *loc, actual, comparison, expected)
	t.Fail()
}

func AssertNull(actual interface{}, t *testing.T) {
	loc := getLocation()
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
	loc := getLocation()
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
	loc := getLocation()
	if actual != true {
		printFail(t, loc, actual, "be", true)
	}
}
func AssertFalse(actual bool, t *testing.T) {
	loc := getLocation()
	if actual != false {
		printFail(t, loc, actual, "be", false)
	}
}

func AssertLengthEqual(actual interface{}, expLen int, t *testing.T) {
	loc := getLocation()
	comparison := "equal to length"
	switch act := actual.(type) {
	case string:
		if len(act) != expLen {
			printFail(t, loc, actual, comparison, expLen)
		}
	case []byte:
		actualBytes := actual.([]byte)
		if len(act) != expLen {
			printFail(t, loc, string(actualBytes), comparison, expLen)
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}

func AssertLengthGreaterThan(actual interface{}, expLen int, t *testing.T) {
	loc := getLocation()
	comparison := "greater or equal to length"
	switch act := actual.(type) {
	case string:
		if len(act) < expLen {
			printFail(t, loc, actual, comparison, expLen)
		}
	case []byte:
		actualBytes := actual.([]byte)
		if len(act) < expLen {
			printFail(t, loc, string(actualBytes), comparison, expLen)
		}
	default:
		fmt.Printf("%v:[FAIL]:Not Supported type comparison expected\n", loc)
		t.Fail()
	}
}
