package appCommon

import "reflect"

// Generic Contain function
func Contain(s interface{}, value interface{}) bool {
	// We will iterate over our array or slice using reflection.
	iter := reflect.ValueOf(s)

	// Check whether the passed variable is a slice or an array
	if iter.Kind() != reflect.Slice && iter.Kind() != reflect.Array {
		panic("Invalid data-type. The first argument should be an array or a slice")
	}

	// Iterate over the slice or array
	for i := 0; i < iter.Len(); i++ {
		// If any slice or array element is equal to the passed value then return true
		if reflect.DeepEqual(value, iter.Index(i).Interface()) {
			return true
		}
	}

	// If no slice or array element is equal to the passed value then return false
	return false
}
