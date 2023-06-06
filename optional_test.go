package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultValueString = ""
	stringVal          = "an example string"
	elseString         = "something else"
)

var stringVar = stringVal

func TestNew(t *testing.T) {
	t.Run("non-pointer type", func(t *testing.T) {
		o := New(stringVal)

		assert.True(t, o.isSet)
		assert.Equal(t, stringVal, o.data)
	})

	t.Run("pointer type", func(t *testing.T) {
		o := New(&stringVar)

		assert.True(t, o.isSet)
		assert.Equal(t, &stringVar, o.data)
	})
}

func TestEmpty(t *testing.T) {
	o := Empty[string]()

	assert.False(t, o.isSet)
	assert.Equal(t, defaultValueString, o.data)
}

func TestNewPointer(t *testing.T) {
	t.Run("set pointer", func(t *testing.T) {
		o := NewPointer(&stringVar)

		assert.True(t, o.isSet)
		assert.Equal(t, stringVar, o.data)
	})

	t.Run("nil pointer", func(t *testing.T) {
		var s *string
		o := NewPointer(s)

		assert.False(t, o.isSet)
		assert.Equal(t, defaultValueString, o.data)
	})
}

func TestT_IsSet(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := T[string]{isSet: true}

		assert.True(t, o.IsSet())
	})

	t.Run("empty optional", func(t *testing.T) {
		o := T[string]{isSet: false}

		assert.False(t, o.IsSet())
	})
}

func TestT_Get(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)

		output, ok := o.Get()
		assert.True(t, ok)
		assert.Equal(t, stringVal, output)
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()

		output, ok := o.Get()
		assert.False(t, ok)
		assert.Equal(t, defaultValueString, output)
	})

	t.Run("default constructed optional", func(t *testing.T) {
		var o T[string]

		output, ok := o.Get()
		assert.False(t, ok)
		assert.Equal(t, defaultValueString, output)
	})
}

func TestT_GetOrZero(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)

		assert.Equal(t, stringVal, o.GetOrZero())
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()

		assert.Equal(t, defaultValueString, o.GetOrZero())
	})

	t.Run("default constructed optional", func(t *testing.T) {
		var o T[string]

		assert.Equal(t, defaultValueString, o.GetOrZero())
	})
}

func TestT_GetOrElse(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)

		assert.Equal(t, stringVal, o.GetOrElse(stringVal))
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()

		assert.Equal(t, elseString, o.GetOrElse(elseString))
	})

	t.Run("default constructed optional", func(t *testing.T) {
		var o T[string]

		assert.Equal(t, elseString, o.GetOrElse(elseString))
	})
}

func TestT_GetPointer(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)

		assert.True(t, o.isSet)
		assert.Equal(t, stringVal, *o.GetPointer())
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()

		var nilString *string

		assert.False(t, o.isSet)
		assert.Equal(t, nilString, o.GetPointer())
	})
}

func TestT_Set(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)
		assert.True(t, o.isSet)

		o.Set(elseString)

		assert.True(t, o.isSet)
		assert.Equal(t, elseString, o.data)
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()
		assert.False(t, o.isSet)

		o.Set(elseString)

		assert.True(t, o.isSet)
		assert.Equal(t, elseString, o.data)
	})
}

func TestT_Clear(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		o := New(stringVal)
		assert.True(t, o.isSet)

		o.Clear()

		assert.False(t, o.isSet)
		assert.Equal(t, defaultValueString, o.data)
	})

	t.Run("empty optional", func(t *testing.T) {
		o := Empty[string]()
		assert.False(t, o.isSet)

		o.Clear()

		assert.False(t, o.isSet)
		assert.Equal(t, defaultValueString, o.data)
	})
}
