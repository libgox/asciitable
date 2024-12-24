package asciitable

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal parses an ASCII table into a slice of the specified type.
func Unmarshal[E any](asciiTable string, v E) ([]string, []E, error) {
	var delimitter = "*+|"
	var skip = []string{"---"}
	lines := strings.Split(strings.TrimSpace(asciiTable), "\n")
	if len(lines) < 4 {
		return nil, nil, errors.New("invalid ascii table format, must have at least 4 lines")
	}

	// Look for headers in the first 10 lines
	var headers []string
	var headerLineNumber int

	var defaultLines = 10
	if len(lines) < defaultLines {
		defaultLines = len(lines) - 1
	}

	for i, line := range lines[0:defaultLines] {
		headerLine := strings.Trim(line, delimitter+" ")
		if skipLine(headerLine, skip) {
			continue
		}
		headers = splitRow(headerLine, delimitter)
		if len(headers) > 1 {
			headerLineNumber = i
			break
		}
	}

	// Reflect on the type of E
	var results []E
	elemType := reflect.TypeOf(v)
	if elemType.Kind() != reflect.Struct {
		return nil, nil, errors.New("v must be a struct type")
	}

	// Process rows
	for _, line := range lines[headerLineNumber+1 : len(lines)-1] { // Begin after header
		if skipLine(line, skip) {
			continue
		}

		row := splitRow(strings.Trim(line, delimitter+" "), delimitter)
		if len(row) != len(headers) {
			return nil, nil, fmt.Errorf("row length does not match header length: %v", row)
		}

		// Map header to struct fields
		elem := reflect.New(elemType).Elem()
		for i, header := range headers {
			for j := 0; j < elem.NumField(); j++ {
				field := elemType.Field(j)
				tag := field.Tag.Get("asciitable")
				if tag == header {
					value := row[i]
					err := setFieldValue(elem.Field(j), value)
					if err != nil {
						return nil, nil, fmt.Errorf("failed to set value for field '%s': %v", field.Name, err)
					}
				}
			}
		}

		// Append to results
		results = append(results, elem.Interface().(E))
	}

	return headers, results, nil
}

// Helper function to split a row by the first character in the delimitter string and trim whitespace.
func splitRow(row string, delimitter string) []string {
	var cells []string
	for _, char := range strings.Split(delimitter, "") {
		cells = strings.Split(row, char)
		for i := range cells {
			cells[i] = strings.TrimSpace(cells[i])
		}
		if len(cells) > 1 {
			break
		}
	}
	return cells
}

// Helper function to set a field value with proper type conversion
func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return errors.New("cannot set value to the field")
	}

	switch field.Kind() {
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intVal))
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatVal)
	case reflect.String:
		field.SetString(value)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}

// Return `true` if any string in the skip slice is in the line.
func skipLine(line string, skip []string) bool {
	for _, s := range skip {
		if strings.Contains(line, s) {
			return true
		}
	}
	return false
}
