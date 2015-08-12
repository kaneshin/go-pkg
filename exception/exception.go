package exception

type Exception interface {
	Message() interface{}
}

// throwable is
type throwable struct {
	rescue bool
	msg    interface{}
}

// Message ...
func (t throwable) Message() interface{} {
	return t.msg
}

// New ...
func New() *throwable {
	t := &throwable{
		rescue: false,
		msg:    nil,
	}
	return t
}

// Try ...
func Try(try func()) *throwable {
	t := New()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.rescue = true
				t.msg = r
			}
		}()
		try()
	}()
	return t
}

// Catch ...
func (t *throwable) Catch(catch func(Exception)) *throwable {
	if t.rescue {
		catch(*t)
	}
	return t
}

// Finally ...
func (e *throwable) Finally(finally func()) {
	finally()
}
