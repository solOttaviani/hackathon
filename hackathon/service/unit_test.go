package service

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeveloper(t *testing.T) {

	t.Run("Test getting one developer assert type Developer", func(t *testing.T) {
		developer, _ := GetDeveloper()

		want := Developer{}
		want.Results.Name.Title = "Mr"
		want.Results.Name.First = "Juan"

		if reflect.DeepEqual(developer, want) {
			assert.Equal(t, developer, want)
		} else {
			t.Errorf("got %q, wanted %q", developer, want)
		}
	})
}
