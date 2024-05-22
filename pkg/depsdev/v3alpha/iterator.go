package depsdev

type batchJob[T any] struct {
	Item T
	Err  error
}

type Iterator[T any] struct {
	cIn     <-chan batchJob[T]
	item    T
	err     error
	hasNext bool

	cancel func()
}

func (v *Iterator[T]) Next() bool {
	bj, ok := <-v.cIn
	if !ok {
		return false
	}

	v.err = bj.Err
	v.item = any(bj.Item).(T)

	return true
}

func (v *Iterator[T]) Item() (T, error) {
	return v.item, v.err
}

func (v *Iterator[T]) Close() {
	if v.cancel != nil {
		v.cancel()
	}
}
