package utilities

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
	"unicode"
)

func ParseMultipartForm(body any) (*bytes.Buffer, string, error) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	val := reflect.ValueOf(body)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	t := val.Type()
	fields := reflect.VisibleFields(t)

	for _, field := range fields {
		fieldVal := val.FieldByName(field.Name)
		if field.Type.Kind() == reflect.Ptr && fieldVal.IsNil() {
			continue
		}

		tag := field.Tag.Get("form")
		if tag == "" || tag == "-" {
			continue
		}

		if strings.Contains(tag, ",") {
			tagParts := strings.Split(tag, ",")
			tag = tagParts[0]
		}

		val := fieldVal.Interface()
		if field.Type.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				continue
			}

			val = fieldVal.Elem().Interface()
		}

		if field.Type.Kind() == reflect.String && !isASCII(val.(string)) {
			val = []byte(val.(string))
		}

		valStr := fmt.Sprintf("%v", val)
		if valStr == "" {
			continue
		}

		bodyWriter.WriteField(tag, valStr)
	}

	err := bodyWriter.Close()
	if err != nil {
		return nil, "", err
	}

	return bodyBuffer, bodyWriter.FormDataContentType(), nil
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}
