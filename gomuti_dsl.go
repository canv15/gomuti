package gomuti

import (
	"reflect"

	"github.com/xeger/gomuti/types"
)

// Allow accepts an instance of Mock, or any struct contains a Mock, and returns
// an object that can be used to program an expected method call to the mock.
func Allow(double interface{}) *types.Allowed {
	m := types.FindMock(reflect.ValueOf(double))
	return m.Allow()
}

// Â is an alias for Allow. Use Shift+Option+M to type this symbol on Mac; Alt+0194 on Windows.
func Â(double interface{}) *types.Allowed {
	return Allow(double)
}