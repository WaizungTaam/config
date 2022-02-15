package config

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

func setDefaultValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.Bool:
		return setBool(field, value)
	case reflect.Int:
		return setInt(field, value)
	case reflect.Int8:
		return setInt8(field, value)
	case reflect.Int16:
		return setInt16(field, value)
	case reflect.Int32:
		return setInt32(field, value)
	case reflect.Int64:
		err := setDuration(field, value)
		if err != nil {
			err = setInt64(field, value)
		}
		return err
	case reflect.Uint:
		return setUint(field, value)
	case reflect.Uint8:
		return setUint8(field, value)
	case reflect.Uint16:
		return setUint16(field, value)
	case reflect.Uint32:
		return setUint32(field, value)
	case reflect.Uint64:
		return setUint64(field, value)
	case reflect.Float32:
		return setFloat32(field, value)
	case reflect.Float64:
		return setFloat64(field, value)
	case reflect.String:
		return setString(field, value)
	case reflect.Slice:
		return setSlice(field, value)
	case reflect.Map:
		return setMap(field, value)
	case reflect.Struct:
		return setStruct(field, value)
	default:
		return nil
	}
}

func setBool(field reflect.Value, value string) error {
	val, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(val))
	return nil
}

func setInt(field reflect.Value, value string) error {
	val, err := strconv.ParseInt(value, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(int(val)))
	return nil
}

func setInt8(field reflect.Value, value string) error {
	val, err := strconv.ParseInt(value, 0, 8)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(int8(val)))
	return nil
}

func setInt16(field reflect.Value, value string) error {
	val, err := strconv.ParseInt(value, 0, 16)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(int16(val)))
	return nil
}

func setInt32(field reflect.Value, value string) error {
	val, err := strconv.ParseInt(value, 0, 32)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(int32(val)))
	return nil
}

func setInt64(field reflect.Value, value string) error {
	val, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(val))
	return nil
}

func setDuration(field reflect.Value, value string) error {
	val, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(val))
	return nil
}

func setUint(field reflect.Value, value string) error {
	val, err := strconv.ParseUint(value, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(uint(val)))
	return nil
}

func setUint8(field reflect.Value, value string) error {
	val, err := strconv.ParseUint(value, 0, 8)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(uint8(val)))
	return nil
}

func setUint16(field reflect.Value, value string) error {
	val, err := strconv.ParseUint(value, 0, 16)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(uint16(val)))
	return nil
}

func setUint32(field reflect.Value, value string) error {
	val, err := strconv.ParseUint(value, 0, 32)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(uint32(val)))
	return nil
}

func setUint64(field reflect.Value, value string) error {
	val, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(uint64(val)))
	return nil
}

func setFloat32(field reflect.Value, value string) error {
	val, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(float32(val)))
	return nil
}

func setFloat64(field reflect.Value, value string) error {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	field.Set(reflect.ValueOf(val))
	return nil
}

func setString(field reflect.Value, value string) error {
	field.Set(reflect.ValueOf(value))
	return nil
}

func setSlice(field reflect.Value, value string) error {
	s := reflect.MakeSlice(field.Type(), 0, 0)
	p := reflect.New(field.Type())
	p.Elem().Set(s)
	if err := json.Unmarshal([]byte(value), p.Interface()); err != nil {
		return err
	}
	field.Set(p.Elem())
	return nil
}

func setMap(field reflect.Value, value string) error {
	m := reflect.MakeMap(field.Type())
	p := reflect.New(field.Type())
	p.Elem().Set(m)
	if err := json.Unmarshal([]byte(value), p.Interface()); err != nil {
		return err
	}
	field.Set(p.Elem())
	return nil
}

func setStruct(field reflect.Value, value string) error {
	p := field.Addr().Interface()
	if err := json.Unmarshal([]byte(value), p); err != nil {
		return err
	}
	return nil
}
