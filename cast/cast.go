package cast

import (
	"github.com/ohler55/ojg/oj"
	"github.com/spf13/cast"
)

func isFloatAnInteger(val float64) bool {
	return val == float64(int(val))
}

// ToCorrectType will convert the given type to correct type.
// We are only worried about types that are supported by JSON.
func ToCorrectType(input any) any {
	if input == nil {
		return nil
	}
	switch input.(type) {
	case bool:
		return cast.ToBool(input)
	case int:
		return cast.ToInt(input)
	case float64:
		value := cast.ToFloat64(input)
		// When parsing JSON without struct, the actual int will be marked as float64
		// To check if float64 is actually int we have to add this check
		if isFloatAnInteger(value) {
			return cast.ToInt(input)
		}
		return value
	case string:
		return cast.ToString(input)
	default:
		return oj.JSON(input)
	}

	return cast.ToString(input)
}
