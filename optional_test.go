package optional

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultValueString = ""
	stringVal          = "an example string"
	elseString         = "something else"
)

var stringVar = stringVal

type testUserStruct struct {
	Name         string // Sanity check
	EmailAddress T[string]
}

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

func TestT_MarshalJSON(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		t.Run("bare value", func(t *testing.T) {
			o := New(stringVal)

			data, err := o.MarshalJSON()

			assert.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("%s%s%s", `"`, stringVal, `"`), string(data))
		})

		t.Run("value as field in a struct", func(t *testing.T) {
			user := testUserStruct{
				Name:         "John Smith",
				EmailAddress: New(stringVal),
			}

			data, err := json.Marshal(user)

			assert.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("%s%s%s", `{"Name":"John Smith","EmailAddress":"`, stringVal, `"}`), string(data))
		})
	})

	t.Run("empty optional", func(t *testing.T) {
		t.Run("bare value", func(t *testing.T) {
			o := Empty[string]()

			data, err := o.MarshalJSON()

			assert.NoError(t, err)
			assert.Equal(t, "null", string(data))
		})

		t.Run("value as field in a struct", func(t *testing.T) {
			user := testUserStruct{
				Name:         "John Smith",
				EmailAddress: Empty[string](),
			}

			data, err := json.Marshal(user)

			assert.NoError(t, err)
			assert.Equal(t, `{"Name":"John Smith","EmailAddress":null}`, string(data))
		})
	})
}

func TestT_UnmarshalJSON(t *testing.T) {
	t.Run("initialized optional", func(t *testing.T) {
		t.Run("bare value", func(t *testing.T) {
			jsonValue := fmt.Sprintf("%s%s%s", `"`, stringVal, `"`)

			var o T[string]
			err := o.UnmarshalJSON([]byte(jsonValue))

			assert.NoError(t, err)
			assert.True(t, o.isSet)
			assert.Equal(t, stringVal, o.data)
		})

		t.Run("value as field in a struct", func(t *testing.T) {
			jsonValue := fmt.Sprintf("%s%s%s", `{"Name":"John Smith","EmailAddress":"`, stringVal, `"}`)

			var user testUserStruct
			err := json.Unmarshal([]byte(jsonValue), &user)

			assert.NoError(t, err)
			assert.True(t, user.EmailAddress.isSet)
			assert.Equal(t, stringVal, user.EmailAddress.data)
		})
	})

	t.Run("empty optional", func(t *testing.T) {
		t.Run("bare value", func(t *testing.T) {
			jsonValue := "null"

			var o T[string]
			err := o.UnmarshalJSON([]byte(jsonValue))

			assert.NoError(t, err)
			assert.False(t, o.isSet)
			assert.Equal(t, defaultValueString, o.data)
		})

		t.Run("value as field in a struct", func(t *testing.T) {
			jsonValue := `{"Name":"John Smith","EmailAddress":null}`

			var user testUserStruct
			err := json.Unmarshal([]byte(jsonValue), &user)

			assert.NoError(t, err)
			assert.False(t, user.EmailAddress.isSet)
			assert.Equal(t, defaultValueString, user.EmailAddress.data)
		})
	})
}
