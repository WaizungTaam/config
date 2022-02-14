package config

import "reflect"

func isZero(value reflect.Value) bool {
	val := value.Interface()
	zero := reflect.Zero(value.Type()).Interface()
	return reflect.DeepEqual(val, zero)
}
