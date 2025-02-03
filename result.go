package result

type Result[A any] interface {
	IsOk() bool
	IsError() bool
	GetValue() A
	GetError() error
	Unwrap() (A, error)
}

type OkType[A any] struct {
	value A
}

func (o OkType[A]) IsOk() bool {
	return true
}
func (o OkType[A]) IsError() bool {
	return false
}
func (o OkType[A]) GetValue() A {
	return o.value
}
func (o OkType[A]) GetError() error {
	return nil
}
func (o OkType[A]) Unwrap() (A, error) {
	return o.value, nil
}

func Ok[A any](value A) Result[A] {
	return OkType[A]{value}
}

type ErrorType[A any] struct {
	err error
}

func (e ErrorType[A]) IsOk() bool {
	return false
}
func (e ErrorType[A]) IsError() bool {
	return true
}
func (e ErrorType[A]) GetValue() A {
	var zero A
	return zero
}
func (e ErrorType[A]) GetError() error {
	return e.err
}
func (e ErrorType[A]) Unwrap() (A, error) {
	return e.GetValue(), e.err
}

func Error[A any](err error) Result[A] {
	return ErrorType[A]{err}
}
