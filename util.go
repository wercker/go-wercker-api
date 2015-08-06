package wercker

import (
	"reflect"
	"strings"

	"github.com/jtacoma/uritemplates"
)

func struct2map(v interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(v)
	switch value.Type().Kind() {
	case reflect.Ptr:
		return struct2map(value.Elem().Interface())
	case reflect.Struct:
		m := make(map[string]interface{})
		for i := 0; i < value.NumField(); i++ {
			tag := value.Type().Field(i).Tag
			omitEmpty := false

			var name string
			if strings.Contains(string(tag), ":") {
				s := strings.Split(tag.Get("map"), ",")
				name = s[0]
				if len(s) == 2 {
					omitEmpty = s[1] == "omitempty"
				}
			}

			if name == "-" {
				continue
			}

			if len(name) == 0 {
				name = value.Type().Field(i).Name
			}

			if omitEmpty && isEmptyValue(value.Field(i)) {
				continue
			}

			m[name] = value.Field(i).Interface()
		}
		return m, true
	}
	return nil, false
}

// isEmptyValue checks if v is is an "empty" value accoriding to its type.
// Taken from src/encoding/json/encode.go https://golang.org/src/encoding/json/encode.go?s=5367:5410#L280
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// addURITemplate adds rawTemplate to templates using name as the key.
func addURITemplate(templates map[string]*uritemplates.UriTemplate, name string, rawTemplate string) {
	uriTemplate, err := uritemplates.Parse(rawTemplate)
	if err != nil {
		panic(err)
	}
	templates[name] = uriTemplate
}
