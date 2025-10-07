package helpers

import (
	"fmt"
	"reflect"
)

func ExpectedType[T any](r any) T {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	recievedType := reflect.TypeOf(r)
	if expectedType == recievedType {
		return r.(T)
	}
	panic(fmt.Sprintf("Error::Type->Expected %T but recieved %T instead", expectedType, recievedType))
}
