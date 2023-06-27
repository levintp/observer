package common

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/levintp/observer/internal/log"
)

// Function to set the default values of all fields in a structure, recursively.
func SetDefaults(obj any) error {
	objValue := reflect.ValueOf(obj)

	// Dereference pointer if needed.
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()
	objKind := objType.Kind()

	switch objKind {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldValue := objValue.Field(i)
			fieldType := objType.Field(i)
			fieldKind := fieldType.Type.Kind()

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
				if err := SetDefaults(fieldValue.Addr().Interface()); err != nil {
					return err
				}
			}
		}
	case reflect.Map:
		// Iterate through the map and decend into structure elements.
		iter := objValue.MapRange()
		for iter.Next() {
			elem := iter.Value()
			SetDefaults(elem.Interface())
		}
	default:
		return fmt.Errorf("error while determining default value of field of kind %v", objKind)
	}

	return nil
}

// Function to set fields of a structure to matching values from the environment, recursively.
func SetEnvironment(obj any, prefix string) error {
	objValue := reflect.ValueOf(obj)

	// Dereference pointer if needed.
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()
	objKind := objType.Kind()

	if objKind == reflect.Struct {
		for i := 0; i < objType.NumField(); i++ {
			fieldValue := objValue.Field(i)
			fieldType := objType.Field(i)
			fieldKind := fieldType.Type.Kind()

			// If the current field has a `env` tag, use it as the environment variable name.
			if environVar := fieldType.Tag.Get("env"); environVar != "" {
				environVar = fmt.Sprintf("%s%s", prefix, environVar)
				// If the environment variable exists, set the field to it.
				log.Debugf("looking for environment variable %s", environVar)
				if environValue := os.Getenv(environVar); environValue != "" {
					log.Debugf("found environment variable: %s=%s", environVar, environValue)
					if err := setField(fieldValue, environValue); err != nil {
						return err
					}
				}
			}
			if fieldKind == reflect.Struct {
				if err := SetEnvironment(fieldValue.Addr().Interface(), prefix); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Function to set fields of a structure to matching values from commandline flags, recursively.
func SetFlags(obj any) error {
	objValue := reflect.ValueOf(obj)

	// Dereference pointer if needed.
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()
	objKind := objType.Kind()

	if objKind == reflect.Struct {
		for i := 0; i < objType.NumField(); i++ {
			fieldValue := objValue.Field(i)
			fieldType := objType.Field(i)
			fieldKind := fieldType.Type.Kind()

			// If the current field has a `flag` tag, use it as the flag name.
			if flagName := fieldType.Tag.Get("flag"); flagName != "" {
				// If the flag is present, set the field to it.
				log.Debugf("looking for flag %s", flagName)
				if flagValue := GetFlag(flagName); flagValue != "" {
					log.Debugf("found flag: %s=%s", flagName, flagValue)
					if err := setField(fieldValue, flagValue); err != nil {
						return err
					}
				}
			}
			if fieldKind == reflect.Struct {
				if err := SetFlags(fieldValue.Addr().Interface()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Function to set a field's value.
func setField(field reflect.Value, value string) error {
	if !field.CanSet() {
		return errors.New("failed to set field value")
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
