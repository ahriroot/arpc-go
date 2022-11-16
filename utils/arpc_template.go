package utils

import (
	"fmt"
)

func TypeStr2GoType(type_str string) string {
	switch type_str {
	case "integer":
		return "int"
	case "int":
		return "int"
	case "int4":
		return "int4"
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "float":
		return "float32"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "double":
		return "float64"
	case "json":
		return "map[string]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return type_str
	}
}

func GenerateParamStruct(name string, params []Param) string {
	var st = fmt.Sprintf("type %s struct {", name)
	var tmp = make(map[string]string)
	var max_length = 0
	for _, param := range params {
		length := len(param.Name)
		if length > max_length {
			max_length = length
		}
		tmp[param.Name] = TypeStr2GoType(param.Type)
	}
	format := fmt.Sprintf("\n    %%-%ds %%s", max_length)
	for k, v := range tmp {
		st += fmt.Sprintf(format, k, v)
	}
	st += "\n}"
	return st
}
