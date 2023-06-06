package optional

type T[Type any] struct {
	data  Type
	isSet bool
}

func New[Type any](data Type) T[Type] {
	return T[Type]{
		data:  data,
		isSet: true,
	}
}

func Empty[Type any]() T[Type] {
	return T[Type]{
		isSet: false,
	}
}

func NewPointer[Type any](data *Type) T[Type] {
	if data == nil {
		return Empty[Type]()
	}

	return New[Type](*data)
}

func (o *T[Type]) IsSet() bool {
	return o.isSet
}

func (o *T[Type]) Get() (Type, bool) {
	if !o.isSet {
		return *new(Type), false
	}

	return o.data, true
}

func (o *T[Type]) GetOrZero() Type {
	if !o.isSet {
		return *new(Type)
	}

	return o.data
}

func (o *T[Type]) GetOrElse(elseValue Type) Type {
	if !o.isSet {
		return elseValue
	}

	return o.data
}

func (o *T[Type]) GetPointer() *Type {
	if !o.isSet {
		return nil
	}

	return &o.data
}

func (o *T[Type]) Set(data Type) {
	o.data = data
	o.isSet = true
}

func (o *T[Type]) Clear() {
	o.data = *new(Type)
	o.isSet = false
}
