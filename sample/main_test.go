package sample

import (
	"testing"
	asserts "github.com/shugit/go-simple-asserts"
)

func Test_String(t *testing.T) {
	asserts.FileType = asserts.PrintAbsolutePath
	asserts.AssertEqual("a,", "b", t)
}

func Test_Int(t *testing.T) {
	asserts.FileType = asserts.PrintFileName
	asserts.AssertEqual(1, 2, t)
	asserts.AssertTrue(false, t)
	asserts.AssertFalse(true, t)
	asserts.AssertGreaterThan(3, 2, t)
	asserts.AssertGreaterThan(2, 3, t)
	asserts.AssertNull(2, t)
	asserts.AssertNotNull(nil, t)
}

func Test_Bytes(t *testing.T) {
	var abc = []byte("abc")
	var bcd = []byte("bcd")
	asserts.AssertEqual(abc, []byte("abc"), t)
	asserts.AssertEqual(bcd, []byte("bbb"), t)
	asserts.AssertLengthEqual(bcd, 7, t)
	asserts.AssertLengthGreaterThan(bcd, 7, t)
}
