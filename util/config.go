package util

import (
	"fmt"
	"reflect"
	"strconv"
)

const KEY_NAME = "name"
const KEY_DEFAULT_VALUE = "default"

func ResolveStruct(dict *map[string]interface{}, prefix string, p interface{}) error {
	pt := reflect.TypeOf(p)
	if pt.Kind() != reflect.Ptr {
		return fmt.Errorf("the param [p] must be &struct")
	}
	structType := pt.Elem()
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("the param [p] must be &struct")
	}

	pv := reflect.ValueOf(p).Elem()
	return resolveStruct(dict, prefix, structType, &pv)
}

func FlatMap(dict *map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range *dict {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			tmp := value.(map[interface{}]interface{})
			for k, v := range FlatMap(&tmp) {
				result[fmt.Sprintf("%s.%s", key, k)] = v
			}
		} else {
			result[fmt.Sprintf("%s", key)] = value
		}
	}
	return result
}

func FlatStringKeyMap(dict *map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range *dict {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			tmp := value.(map[string]interface{})
			for k, v := range FlatStringKeyMap(&tmp) {
				result[key+"."+k] = v
			}
		} else {
			result[key] = value
		}
	}
	return result
}

func resolveStruct(dict *map[string]interface{}, prefix string, structType reflect.Type, structValue *reflect.Value) error {
	count := structType.NumField()
	if count == 0 {
		return nil
	}
	if prefix != "" {
		prefix += "."
	}
	for index := 0; index < count; index++ {
		field := structType.Field(index)
		tag := field.Tag
		name := tag.Get(KEY_NAME)
		if name == "" {
			continue
		}
		key := prefix + name
		fieldValue := structValue.Field(index)
		if field.Type.Kind() == reflect.Struct {
			resolveStruct(dict, key, field.Type, &fieldValue)
			continue
		}
		value, ok := (*dict)[key]
		if !ok {
			value = tag.Get("default")
			required := tag.Get("required")
			if value == "" && required == "true" {
				return fmt.Errorf("not found value by [" + key + "]")
			}
		}
		setValue(&fieldValue, value)
	}
	return nil
}

func setValue(fieldValue *reflect.Value, value interface{}) error {
	switch fieldValue.Type().Kind() {
	case reflect.Int:
		val, err := convertInt64(value)
		if err != nil {
			return err
		}
		fieldValue.SetInt(val)
	case reflect.Float64, reflect.Float32:
		val, err := convertFloat64(value)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(val)
	case reflect.String:
		val, ok := value.(string)
		if ok {
			fieldValue.SetString(val)
		} else {
			fieldValue.SetString(fmt.Sprintf("%v", value))
		}
	}
	return nil
}

func convertFloat64(i interface{}) (float64, error) {
	switch t := i.(type) {
	case float64:
		return t, nil
	case float32:
		return float64(t), nil
	case string:
		return strconv.ParseFloat(t, 10)
	case int64:
		return float64(t), nil
	default:
		return 0, fmt.Errorf("can't convert %v to float64", i)
	}
}

func convertInt64(i interface{}) (int64, error) {
	switch t := i.(type) {
	case int64:
		return t, nil
	case float32:
		return int64(t), nil
	case int:
		return int64(t), nil
	case string:
		tmp, err := strconv.Atoi(t)
		if err != nil {
			return 0, err
		}
		return int64(tmp), nil
	default:
		return 0, fmt.Errorf("can't convert %v to int64", i)
	}

}
