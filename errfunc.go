package errfunc

// Func is a function that takes a single argument of type A and
// returns a single result of type R along with an error. If an error
// is returned, it becomes a no-op.
type Func[A, R any] struct {
	f   func(A) (R, error)
	err error
}

// New returns a new Func that wraps f.
func New[A, R any](f func(A) (R, error)) *Func[A, R] {
	return &Func[A, R]{f: f}
}

// Call calls the underlying function with the provided arguument and
// returns the result. If an error has occured on a previous call, it
// does not call the function and instead simply returns the zero-value
// of R.
func (f *Func[A, R]) Call(a A) (r R) {
	if f.err != nil {
		return r
	}

	r, f.err = f.f(a)
	return r
}

// Err returns the error that was returned by the underlying function
// during a previous call of it, or nil if no such error occurred.
func (f *Func[A, R]) Err() error {
	return f.err
}
