package config

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/levintp/observer/internal/logging"
)

// Function to set the default values of all fields in a structure, recursively.
func setDefaults(obj any) error {
	objValue := reflect.ValueOf(obj)

	// Dereference pointer if needed.
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()
	objKind := objType.Kind()

	logging.Logger.Tracef("Enumerating object type=%v kind=%v", objType.Name(), objKind)

	switch objKind {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldValue := objValue.Field(i)
			fieldType := objType.Field(i)
			fieldKind := fieldType.Type.Kind()

			logging.Logger.Tracef("Processing field=%v type=%v", fieldType.Name, fieldKind)

			// If the current field has a `default` tag, use it as default value.
			if defaultVal := fieldType.Tag.Get("default"); defaultVal != "" {
				// If the field is empty (zero value), set it to its default value.
				if fieldValue.Interface() == reflect.Zero(fieldType.Type).Interface() {
					if err := setField(fieldValue, defaultVal); err != nil {
						return err
					}
				}
			}
			if fieldKind == reflect.Struct || fieldKind == reflect.Map {
				logging.Logger.Tracef("Descending into field=%v kind=%v", fieldType.Name, fieldKind)
				if err := setDefaults(fieldValue.Addr().Interface()); err != nil {
					return err
				}
				logging.Logger.Tracef("Ascending from field=%v", fieldType.Name)
			}
		}
	case reflect.Map:
		// Iterate through the map and decend into structure elements.
		iter := objValue.MapRange()
		for iter.Next() {
			elem := iter.Value()
			setDefaults(elem.Interface())
		}
	default:
		return fmt.Errorf("Error while determining default value of field of kind %v", objKind)
	}

	return nil
}

// Function to set a field's value.
func setField(field reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("Failed to set field value\n")
	}

	switch field.Kind() {
	case reflect.Int:
		if val, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(value).Convert(field.Type()))
	}

	return nil
}
