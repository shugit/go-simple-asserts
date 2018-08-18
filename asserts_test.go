package go_simple_asserts

import "testing"

func Test_Main(t *testing.T) {
	AssertEqual("a,", "b", t)
}
