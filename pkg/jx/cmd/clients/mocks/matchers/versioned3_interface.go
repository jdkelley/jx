// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	versioned3 "github.com/knative/build/pkg/client/clientset/versioned"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyVersioned3Interface() versioned3.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned3.Interface))(nil)).Elem()))
	var nullValue versioned3.Interface
	return nullValue
}

func EqVersioned3Interface(value versioned3.Interface) versioned3.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned3.Interface
	return nullValue
}