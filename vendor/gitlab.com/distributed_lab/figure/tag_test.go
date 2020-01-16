package figure

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFieldTag(t *testing.T) {
	cases := []struct {
		name        string
		field       reflect.StructField
		expectedTag *Tag
	}{
		{
			name:        `field name set as tag key in snake case`,
			field:       reflect.StructField{Name: `FooBar`, Tag: ``},
			expectedTag: &Tag{Key: `foo_bar`},
		},
		{
			name:        `check value was recognized`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"foo"`},
			expectedTag: &Tag{Key: `foo`},
		},
		{
			name:        `check ignore tag`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"-"`},
			expectedTag: nil,
		},
		{
			name:        `recognition the tag and attribute`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"foo,required"`},
			expectedTag: &Tag{Key: `foo`, Required: true},
		},
		{
			name:        `tag with multiple attributes`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"foo,required,non_zero"`},
			expectedTag: &Tag{Key: `foo`, Required: true, NonZero: true},
		},
		{
			name: "implicit is still required",
			field: reflect.StructField{Name: `Implicit`, Tag: `fig:",required"`},
			expectedTag: &Tag{Key: "implicit", Required: true},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualTag, err := parseFieldTag(c.field, keyTag)

			assert.Equal(t, actualTag, c.expectedTag)
			assert.NoError(t, err)
		})
	}
}

func TestParseFieldTagErr(t *testing.T) {
	cases := []struct {
		name        string
		field       reflect.StructField
		expectedErr error
	}{
		{
			name:        `Conflicting attributes`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"-,required"`},
			expectedErr: ErrConflictingAttributes,
		},
		{
			name:        `Unknown attribute`,
			field:       reflect.StructField{Name: ``, Tag: `fig:"foo,yoba"`},
			expectedErr: ErrUnknownAttribute,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := parseFieldTag(c.field, keyTag)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
