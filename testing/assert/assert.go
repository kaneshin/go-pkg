package assert

import (
	"fmt"
	"reflect"
)

type Interface interface {
	Errorf(format string, args ...interface{})
}

func check(t Interface, v interface{}, f string, msg ...string) bool {
	ok := true
	switch v := v.(type) {
	case bool:
		if !v {
			ok = false
		}
	case *bool:
		if !*v {
			ok = false
		}
	case error:
		if v != nil {
			ok = false
		}
	}
	if !ok {
		if len(msg) > 0 {
			t.Errorf("%s\n%s", f, msg[0])
		} else {
			t.Errorf("%s", f)
		}
	}
	return ok
}

func deepEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	return reflect.DeepEqual(expected, actual)
}

func Equal(t Interface, expected, actual interface{}, msg ...string) bool {
	v := deepEqual(expected, actual)
	f := fmt.Sprintf(`
Expected: %#v
Actual:   %#v`, expected, actual)
	return check(t, v, f, msg...)
}

func NotEqual(t Interface, expected, actual interface{}, msg ...string) bool {
	v := !deepEqual(expected, actual)
	f := fmt.Sprintf(`
Not Expected: %#v
Actual:       %#v`, expected, actual)
	return check(t, v, f, msg...)
}

func True(t Interface, v bool, msg ...string) bool {
	return Equal(t, true, v, msg...)
}

func False(t Interface, v bool, msg ...string) bool {
	return Equal(t, false, v, msg...)
}
