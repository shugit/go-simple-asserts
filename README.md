# go-simple-asserts


A very simple assertion using golang's default testing package


* Could print file name or absolute file path
* No other dependencies
* Support multiple common types
* WILL Support pointer comparison, support fmt.Sprintf("%x",ptr) and pass as string for now

```
import (
	"testing"
	asserts "github.com/shugit/go-simple-asserts"
)

func Test_FailCase1(t *testing.T) {
	asserts.FileType = asserts.PrintFileName
	asserts.AssertEqual("a,", "b", t)
}

func Test_FailCase2(t *testing.T) {
	asserts.FileType = asserts.PrintAbsolutePath
	asserts.AssertEqual(1, 2, t)
	asserts.AssertNotNull(nil, t)
	asserts.AssertGreaterThan(3, 2, t)
}
```