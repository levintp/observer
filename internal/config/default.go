package config

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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

	log.Printf("Enumerating object type=%v kind=%v", objType.Name(), objKind)

	switch objKind {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldValue := objValue.Field(i)
			fieldType := objType.Field(i)
			fieldKind := fieldType.Type.Kind()

			log.Printf("Processing field=%v type=%v", fieldType.Name, fieldKind)

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
				log.Printf("Descending into field=%v kind=%v", fieldType.Name, fieldKind)
				if err := setDefaults(fieldValue.Addr().Interface()); err != nil {
					return err
				}
				log.Printf("Ascending from field=%v", fieldType.Name)
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

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Failed to set field value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}
