package assert

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFunc asserts the given two functions.
func AssertFunc(t *testing.T, expected any, actual any) {
	expectedFunc := runtime.FuncForPC(reflect.ValueOf(expected).Pointer()).Name()
	actualFunc := runtime.FuncForPC(reflect.ValueOf(actual).Pointer()).Name()

	assert.Equal(t, expectedFunc, actualFunc)
}
