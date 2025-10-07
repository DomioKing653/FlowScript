package helpers

import (
	"fmt"
	"reflect"
)

func ExpectedType[T any](r any) T {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()

	if r == nil {
		panic(fmt.Sprintf("Error::Type->Expected %s but received nil", expectedType))
	}

	rv := reflect.ValueOf(r)
	// If we have a pointer, dereference to compare the underlying type
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			panic(fmt.Sprintf("Error::Type->Expected %s but received nil pointer", expectedType))
		}
		rv = rv.Elem()
	}

	receivedType := rv.Type()
	if receivedType == expectedType {
		// return the concrete value as T
		return rv.Interface().(T)
	}

	panic(fmt.Sprintf("Error::Type->Expected %s but received %s instead", expectedType, receivedType))
}
