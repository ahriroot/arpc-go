package utils

import (
	"fmt"
	"strings"
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
	var params_list = make([]string, 0)
	var field_list = make([]string, 0)
	var st = fmt.Sprintf("type %s struct {", name)
	var tmp = make(map[string]string)
	var max_length = 0
	for _, param := range params {
		length := len(param.Name)
		if length > max_length {
			max_length = length
		}
		snake := Snake(param.Name)
		tmp[param.Name] = TypeStr2GoType(param.Type)
		params_list = append(params_list, fmt.Sprintf("%s %s", snake, tmp[param.Name]))
		field_list = append(field_list, fmt.Sprintf("b.%s = %s", param.Name, snake))
	}
	format := fmt.Sprintf("\n    %%-%ds %%s", max_length)
	for k, v := range tmp {
		st += fmt.Sprintf(format, k, v)
	}
	st += "\n}"

	var new_str = `

func (b *%s) New(%s) {
	%s
}`
	st += fmt.Sprintf(new_str, name, strings.Join(params_list, ", "), strings.Join(field_list, "\n    "))

	var serialize_str = `

func (b *%s) Serialize() ([]byte, error) {
	return json.Marshal(b)
}`
	st += fmt.Sprintf(serialize_str, name)

	var deserialize_str = `

func (b *%s) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}`
	st += fmt.Sprintf(deserialize_str, name)
	return st
}

// // New ApiRequestV1
// func (b *ApiRequestV1) New(user_id int) {
// 	b.UserId = user_id
// }

// // json Serialize ApiRequestV1
// func (b *ApiRequestV1) Serialize() ([]byte, error) {
// 	return json.Marshal(b)
// }

// // json Deserialize ApiRequestV1
// func (b *ApiRequestV1) Deserialize(data []byte) error {
// 	return json.Unmarshal(data, b)
// }
