package console

import (
	"os"
	"strconv"
	"unsafe"
)

// Stringer - Useful for custom types or structs formatting.
type Stringer interface {
	String() string
}

// WriteLine - Prints out its arguments, returns the number of bytes written and any write error encountered.
func WriteLine(args ...interface{}) (int, error) {
	var bytes []byte

	for i, value := range args {
		bytes = append(bytes, stringify(value)...)

		// Add a spacing
		if i != len(args)-1 {
			bytes = append(bytes, ' ')
		}
	}
	return os.Stdout.Write(append(bytes, '\n'))
}

// Write - Works the same as WriteLine except this doesn't put a newline at the end
func Write(args ...interface{}) (int, error) {
	var bytes []byte

	for i, value := range args {
		bytes = append(bytes, stringify(value)...)

		// Add a spacing
		if i != len(args)-1 {
			bytes = append(bytes, ' ')
		}
	}
	return os.Stdout.Write(bytes)
}

// WriteBytes - Output raw bytes to the stdout.
// Returns the number of bytes written and an error, if any. Returns a non-nil error when int != len(bytes).
func WriteBytes(bytes []byte) (int, error) {
	return os.Stdout.Write(bytes)
}

// stringify - returns the value as a string.
func stringify(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"

	// Integers
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)

	// Floats
	case float32:
		return strconv.FormatFloat(float64(v), 'v', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'v', -1, 64)

	// Unsigned Integers
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)

	default:
		if value, isStringer := value.(Stringer); isStringer {
			return value.String()
		}
	}
	return "<UNKNOWN>"
}

// Group - Groups a bunch of values for you and surrounds them with '{...}'.
// Useful for 'String' method on structs.
func Group(args ...interface{}) string {
	bytes := []byte{'{'}

	for i, value := range args {
		// Add value
		bytes = append(bytes, stringify(value)...)

		// Add a spacing
		if i != len(args)-1 {
			bytes = append(bytes, ',', ' ')
		}
	}

	//return string(append(bytes, '}'))
	bytes = append(bytes, '}')
	return *(*string)(unsafe.Pointer(&bytes))
}
