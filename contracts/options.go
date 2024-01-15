package contracts

func ApplyOptions[T any, S any](o *S, options ...T) {
	if o == nil {
		o = new(S)
	}
	for _, fn := range options {
		fn(o)
	}
}
