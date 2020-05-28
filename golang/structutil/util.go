package main

import (
	"encoding/json"
	"errors"
	"reflect"
)

func ReflectTag(tagName string, valueFunc func(string) string, data interface{}) error {
	tag, ok := lookupTag(tagName, data)
	if !ok {
		return errors.New("")
	}
	value := valueFunc(tag)
	return setDynamic(tag, value, data)
}

func lookupTag(tag string, data interface{}) (value string, ok bool) {
	t := reflect.TypeOf(data)
	if t == nil {
		return "", false
	}
	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		tag, ok := field.Tag.Lookup(tag)
		if ok {
			return tag, ok
		}
	}
	return "", false
}

func setDynamic(key string, value interface{}, data interface{}) error {
	m, err := json.Marshal(map[string]interface{}{
		key: value,
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(m, data)
	return err
}
